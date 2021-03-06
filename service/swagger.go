package service

type j map[string]interface{}
type k []string

var swaggerMap = j{
	"basePath": "/",
	"swagger":  "2.0",
	"info": j{
		"title":       "Notifications API",
		"version":     "0.0.1",
		"description": "API for bastion management",
	},

	"paths": j{
		"/services/pagerduty": j{
			"post": j{
				"parameters": []j{
					j{
						"description": "Pagerduty account",
						"in":          "query",
						"name":        "account",
						"required":    true,
						"type":        "string",
					},
					j{
						"description": "Pagerduty service_key.",
						"in":          "query",
						"name":        "service_key",
						"required":    true,
						"type":        "string",
					},
					j{
						"description": "Pagerduty service_name",
						"in":          "service_name",
						"name":        "state",
						"required":    true,
						"type":        "string",
					},
					j{
						"description": "Pagerduty error",
						"in":          "error",
						"name":        "state",
						"required":    false,
						"type":        "string",
					},
				},
				"responses": j{
					"200": j{
						"description": "",
						"schema": j{
							"$ref": "#/definitions/PagerDutyOAuthResponse",
						},
					},
				},
				"summary": "Saves a pagerduty token.",
				"tags":    k{"token"},
			},
			"get": j{
				"parameters": []j{},
				"responses": j{
					"200": j{
						"description": "Retrieves user's pagerduty token.",
						"schema": j{
							"$ref": "#/definitions/PagerDutyOAuthResponse",
						},
					},
				},
				"summary": "Create a new notification.",
				"tags":    k{"getpagerdutytoken"},
			},
		},
		"/services/pagerduty/test": j{
			"post": j{
				"parameters": []j{
					j{
						"description": "",
						"in":          "body",
						"name":        "Notifications",
						"required":    true,
						"schema": j{
							"$ref": "#/definitions/Notifications",
						},
					},
				},
				"responses": j{
					"200": j{
						"description": "",
						"schema":      j{},
					},
				},
				"summary": "Test alert on a notification",
				"tags":    k{"notifications"},
			},
		},

		"/services/email/test": j{
			"post": j{
				"parameters": []j{
					j{
						"description": "",
						"in":          "body",
						"name":        "Notifications",
						"required":    true,
						"schema": j{
							"$ref": "#/definitions/Notifications",
						},
					},
				},
				"responses": j{
					"200": j{
						"description": "",
						"schema":      j{},
					},
				},
				"summary": "Test alert on a notification",
				"tags":    k{"notifications"},
			},
		},

		"/services/webhook/test": j{
			"post": j{
				"parameters": []j{
					j{
						"description": "",
						"in":          "body",
						"name":        "Notifications",
						"required":    true,
						"schema": j{
							"$ref": "#/definitions/Notifications",
						},
					},
				},
				"responses": j{
					"200": j{
						"description": "",
						"schema":      j{},
					},
				},
				"summary": "Test alert on a notification",
				"tags":    k{"notifications"},
			},
		},

		"/services/slack/channels": j{
			"get": j{
				"parameters": []j{},
				"responses": j{
					"200": j{
						"description": "List customer's slack channels.",
						"schema": j{
							"$ref": "#/definitions/SlackChannelsResponse",
						},
					},
				},
				"summary": "Create a new notification.",
				"tags":    k{"getslackchannels"},
			},
		},
		"/services/slack": j{
			"post": j{
				"parameters": []j{
					j{
						"description": "",
						"in":          "body",
						"name":        "SlackOAuthRequest",
						"required":    true,
						"schema": j{
							"$ref": "#/definitions/SlackOAuthRequest",
						},
					},
				},
				"responses": j{
					"200": j{
						"description": "",
						"schema": j{
							"$ref": "#/definitions/SlackOAuthResponse",
						},
					},
				},
				"summary": "Create a new notification.",
				"tags":    k{"postslackauthcode"},
			},
			"get": j{
				"parameters": []j{},
				"responses": j{
					"200": j{
						"description": "Get a customer's slack token",
						"schema": j{
							"$ref": "#/definitions/SlackOAuthResponse",
						},
					},
				},
				"summary": "Get a customer's slack token.",
				"tags":    k{"getslacktoken"},
			},
		},
		"/services/slack/code": j{
			"get": j{
				"parameters": []j{
					j{
						"description": "Slack authorization code.",
						"in":          "query",
						"name":        "code",
						"required":    true,
						"type":        "string",
					},
					j{
						"description": "Slack redirect_uri.",
						"in":          "query",
						"name":        "redirect_uri",
						"required":    false,
						"type":        "string",
					},
					j{
						"description": "Slack authorization state.",
						"in":          "query",
						"name":        "state",
						"required":    false,
						"type":        "string",
					},
				},
				"responses": j{
					"200": j{
						"description": "",
						"schema": j{
							"$ref": "#/definitions/SlackOAuthResponse",
						},
					},
				},
				"summary": "Retrieves a test code.",
				"tags":    k{"token"},
			},
		},
		"/services/slack/test": j{
			"post": j{
				"parameters": []j{
					j{
						"description": "",
						"in":          "body",
						"name":        "Notifications",
						"required":    true,
						"schema": j{
							"$ref": "#/definitions/Notifications",
						},
					},
				},
				"responses": j{
					"200": j{
						"description": "",
						"schema":      j{},
					},
				},
				"summary": "Test alert on a notification",
				"tags":    k{"notifications"},
			},
		},

		"/notifications": j{
			"get": j{
				"responses": j{
					"200": j{
						"description": "",
						"schema": j{
							"items": j{
								"$ref": "#/definitions/Notifications",
							},
							"type": "array",
						},
					},
				},
				"summary": "Retrieve all of a customer's notifications",
				"tags":    k{"notifications"},
			},
			"delete": j{
				"parameters": []j{
					j{
						"description": "",
						"in":          "body",
						"name":        "Notifications",
						"required":    true,
						"schema": j{
							"$ref": "#/definitions/Notifications",
						},
					},
				},
				"responses": j{
					"200": j{
						"description": "",
						"schema": j{
							"$ref": "#/definitions/Notifications",
						},
					},
				},
				"summary": "Delete these notifications",
				"tags":    k{"notifications"},
			},
			"post": j{
				"parameters": []j{
					j{
						"description": "",
						"in":          "body",
						"name":        "Notifications",
						"required":    true,
						"schema": j{
							"$ref": "#/definitions/Notifications",
						},
					},
				},
				"responses": j{
					"200": j{
						"description": "",
						"schema": j{
							"$ref": "#/definitions/Notifications",
						},
					},
				},
				"summary": "Create a new notification.",
				"tags":    k{"notifications"},
			},
		},
		"/notifications-multicheck": j{
			"post": j{
				"parameters": []j{
					j{
						"description": "",
						"in":          "body",
						"name":        "NotificationsArray",
						"required":    true,
						"schema": j{
							"items": j{
								"$ref": "#/definitions/Notifications",
							},
							"type": "array",
						},
					},
				},
				"responses": j{
					"200": j{
						"description": "",
						"schema": j{
							"items": j{
								"$ref": "#/definitions/Notifications",
							},
							"type": "array",
						},
					},
				},
				"summary": "Create a new notification.",
				"tags":    k{"notifications"},
			},
		},
		"/notifications/{check_id}": j{
			"delete": j{
				"parameters": []j{
					j{
						"description": "",
						"in":          "path",
						"name":        "check_id",
						"required":    true,
						"type":        "string",
					},
				},
				"responses": j{
					"default": j{
						"description": "",
					},
				},
				"summary": "Deletes a notification.",
				"tags":    k{"notifications"},
			},
			"get": j{
				"parameters": []j{
					j{
						"description": "",
						"in":          "path",
						"name":        "check_id",
						"required":    true,
						"type":        "string",
					},
				},
				"responses": j{
					"200": j{
						"description": "",
						"schema": j{
							"$ref": "#/definitions/Notifications",
						},
					},
				},
				"summary": "Retrieves a notification.",
				"tags":    k{"notifications"},
			},
			"put": j{
				"parameters": []j{
					j{
						"description": "",
						"in":          "body",
						"name":        "Notifications",
						"required":    true,
						"schema": j{
							"$ref": "#/definitions/Notifications",
						},
					},
					j{
						"description": "",
						"in":          "path",
						"name":        "check_id",
						"required":    true,
						"type":        "string",
					},
				},
				"responses": j{
					"200": j{
						"description": "",
						"schema": j{
							"$ref": "#/definitions/Notifications",
						},
					},
				},
				"summary": "Replaces a notification.",
				"tags":    k{"notifications"},
			},
		},
	},

	"definitions": j{
		"Notifications": j{
			"properties": j{
				"check_id": j{
					"type": "string",
				},
				"notifications": j{
					"items": j{
						"$ref": "#/definitions/notification",
					},
					"type": "array",
				},
			},
			"required": k{
				"check_id",
				"notifications",
			},
			"type": "object",
		},
		"Notification": j{
			"properties": j{
				"type": j{
					"type": "string",
				},
				"value": j{
					"type": "string",
				},
			},
			"required": k{
				"type",
				"value",
			},
			"type": "object",
		},
		"PagerDutyOAuthResponse": j{
			"properties": j{
				"account": j{
					"type": "string",
				},
				"service_key": j{
					"type": "string",
				},
				"service_name": j{
					"type": "string",
				},
				"enabled": j{
					"type": "boolean",
				},
				"error": j{
					"type": "string",
				},
			},
		},

		"SlackOAuthRequest": j{
			"properties": j{
				"client_id": j{
					"type": "string",
				},
				"client_secret": j{
					"type": "string",
				},
				"code": j{
					"type": "string",
				},
				"redirect_uri": j{
					"type": "string",
				},
			},
			"required": k{
				"code",
			},
			"type": "object",
		},
		"SlackOAuthResponse": j{
			"properties": j{
				"access_token": j{
					"type": "string",
				},
				"scope": j{
					"type": "string",
				},
				"team_name": j{
					"type": "string",
				},
				"team_domain": j{
					"type": "string",
				},
				"team_id": j{
					"type": "string",
				},
				"ok": j{
					"type": "bool",
				},
				"error": j{
					"type": "string",
				},
			},
			"required": k{
				"access_token", "scope", "team_name", "team_id", "ok",
			},
			"type": "object",
		},
		"SlackChannels": j{
			"properties": j{
				"channels": j{
					"items": j{
						"$ref": "#/definitions/SlackChannel",
					},
					"type": "array",
				},
			},
			"required": k{
				"channels",
			},
			"type": "object",
		},
		"SlackChannel": j{
			"properties": j{
				"id": j{
					"type": "string",
				},
				"name": j{
					"type": "string",
				},
			},
			"required": k{
				"id",
				"name",
			},
			"type": "object",
		},
	},

	"consumes": j{},
	"produces": j{},
}
