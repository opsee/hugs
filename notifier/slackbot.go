package notifier

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"golang.org/x/net/context"

	"github.com/hoisie/mustache"
	"github.com/opsee/basic/schema"
	opsee "github.com/opsee/basic/service"
	"github.com/opsee/hugs/obj"
	"github.com/opsee/hugs/store"
	log "github.com/opsee/logrus"
	slacktmpl "github.com/opsee/notification-templates/dist/go/slack"
)

type SlackBotSender struct {
	templates  map[string]*mustache.Template
	catsClient opsee.CatsClient
}

// Send notification to customer.  At this point we have done basic validation on notification and event
func (this SlackBotSender) Send(n *obj.Notification, e *obj.Event) error {
	result := e.Result

	templateKey := "check-passing"
	if !result.Passing {
		templateKey = "check-failing"
	}

	// Bleh. This is copypasta from email.go
	// TODO(greg): When we move to a generic model, we can figure out a way
	// to centralize all of this logic so that senders can finally be dumb.
	failingResponses := result.FailingResponses()

	// It's a possible error state that if the CheckResult.Passing field is false,
	// i.e. this is a failing event, that there are somehow no constituent failing
	// CheckResponse objects contained within the CheckResult. We cannot know _why_
	// these CheckResponse objects aren't failing. Because we cannot ordain the reason
	// for this error state, let us first err on the side of not bugging a customer.
	if len(failingResponses) < 1 && !result.Passing {
		return errors.New("Received failing CheckResult with no failing responses.")
	}

	if slackTemplate, ok := this.templates[templateKey]; ok {
		token, err := this.getSlackToken(n)
		if err != nil {
			return err
		}

		templateContent := map[string]interface{}{
			"check_id":       result.CheckId,
			"check_name":     result.CheckName,
			"group_name":     result.Target.Id,
			"instance_count": len(result.Responses),
			"fail_count":     len(failingResponses),
			"type":           "target",
			"token":          token,
			"channel":        n.Value,
		}

		if e.Nocap != nil && e.Nocap.JSONUrl != "" {
			templateContent["json_url"] = fmt.Sprintf("/event?json=%s&", url.QueryEscape(e.Nocap.JSONUrl))
		} else {
			templateContent["json_url"] = "?"
		}

		if result.Target.Type == "external_host" {
			catsResponse, err := this.catsClient.GetCheckResults(context.Background(), &opsee.GetCheckResultsRequest{
				CheckId:    result.CheckId,
				CustomerId: result.CustomerId,
			})
			if err != nil {
				return err
			}
			results := catsResponse.Results

			var (
				instanceCount = len(results)
				failCount     int
			)

			for _, r := range results {
				failCount += r.FailingCount()
			}

			// we have inconsistent results, so don't do anything
			if !result.Passing && failCount == 0 {
				return fmt.Errorf("Failing result, but fail count == 0")
			}

			templateContent["instance_count"] = instanceCount
			templateContent["fail_count"] = failCount
			templateContent["type"] = "points-of-presence (PoPs)"
		}

		postMessageRequest := &obj.SlackPostChatMessageRequest{}
		log.Debug(string(slackTemplate.Render(templateContent)))
		err = json.Unmarshal([]byte(slackTemplate.Render(templateContent)), postMessageRequest)
		if err != nil {
			return err
		}

		slackPostMessageResponse, err := postMessageRequest.Do("https://slack.com/api/chat.postMessage")
		if err != nil {
			log.WithFields(log.Fields{"slackbot": "Send", "error": err}).Error("Error sending notification to slack.")
			return err
		}
		if slackPostMessageResponse.OK != true {
			return fmt.Errorf(slackPostMessageResponse.Error)
		}
	}

	return nil
}

func (this SlackBotSender) getSlackToken(n *obj.Notification) (string, error) {
	s, err := store.NewPostgres()
	if err != nil {
		return "", err
	}

	oaResponse, err := s.GetSlackOAuthResponse(&schema.User{CustomerId: n.CustomerId})
	if err != nil {
		return "", err
	}

	// if for whatever reason we don't have a bot
	if oaResponse.Bot == nil {
		log.WithFields(log.Fields{"slackbot": "getSlackToken"}).Error("User does not have a bot token associated with this slack integration.")
		return "", fmt.Errorf("integration_inactive")

	}

	return oaResponse.Bot.BotAccessToken, nil
}

func NewSlackBotSender() (*SlackBotSender, error) {

	// initialize check failing template
	failTemplate, err := mustache.ParseString(slacktmpl.CheckFailing)
	if err != nil {
		return nil, err
	}

	// initialize check passing template
	passTemplate, err := mustache.ParseString(slacktmpl.CheckPassing)
	if err != nil {
		return nil, err
	}

	templateMap := map[string]*mustache.Template{
		"check-failing": failTemplate,
		"check-passing": passTemplate,
	}

	catsConn, err := grpc.Dial(
		"cats.in.opsee.com:443",
		grpc.WithTransportCredentials(
			credentials.NewTLS(&tls.Config{
				InsecureSkipVerify: true,
			}),
		),
	)
	if err != nil {
		return nil, err
	}

	return &SlackBotSender{
		templates:  templateMap,
		catsClient: opsee.NewCatsClient(catsConn),
	}, nil
}
