// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/confirm/{token}": {
            "get": {
                "description": "Confirms a subscription using the token sent in the confirmation email.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscription"
                ],
                "summary": "Confirm email subscription",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Confirmation token",
                        "name": "token",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Subscription confirmed successfully"
                    },
                    "400": {
                        "description": "Invalid token"
                    },
                    "404": {
                        "description": "Token not found"
                    }
                }
            }
        },
        "/subscribe": {
            "post": {
                "description": "Subscribe an email to receive weather updates for a specific city with chosen frequency.",
                "consumes": [
                    "application/json",
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscription"
                ],
                "summary": "Subscribe to weather updates",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email address to subscribe",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "City for weather updates",
                        "name": "city",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "enum": [
                            "hourly",
                            "daily"
                        ],
                        "type": "string",
                        "description": "Frequency of updates (hourly or daily)",
                        "name": "frequency",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Subscription successful. Confirmation email sent."
                    },
                    "400": {
                        "description": "Invalid input"
                    },
                    "409": {
                        "description": "Email already subscribed"
                    }
                }
            }
        },
        "/unsubscribe/{token}": {
            "get": {
                "description": "Unsubscribes an email from weather updates using the token sent in emails.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscription"
                ],
                "summary": "Unsubscribe from weather updates",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Unsubscribe token",
                        "name": "token",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Unsubscribed successfully"
                    },
                    "400": {
                        "description": "Invalid token"
                    },
                    "404": {
                        "description": "Token not found"
                    }
                }
            }
        },
        "/weather": {
            "get": {
                "description": "Returns the current weather forecast for the specified city using WeatherAPI.com.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "weather"
                ],
                "summary": "Get current weather for a city",
                "parameters": [
                    {
                        "type": "string",
                        "description": "City name for weather forecast",
                        "name": "city",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.weatherResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request"
                    },
                    "404": {
                        "description": "City not found"
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.weatherResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "humidity": {
                    "type": "number"
                },
                "temperature": {
                    "type": "number"
                }
            }
        }
    },
    "tags": [
        {
            "description": "Weather forecast operations",
            "name": "weather"
        },
        {
            "description": "Subscription management operations",
            "name": "subscription"
        }
    ]
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "weather-forecast-sub-app.onrender.com",
	BasePath:         "/api",
	Schemes:          []string{"http", "https"},
	Title:            "Weather Forecast API",
	Description:      "Weather API application that allows users to subscribe to weather updates for their city.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
