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
        "/vote": {
            "post": {
                "description": "Cast a vote for a BBB participant",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "votes"
                ],
                "summary": "Cast a vote",
                "parameters": [
                    {
                        "description": "Vote details",
                        "name": "vote",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Vote"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully cast vote",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Failed to cast vote",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/votes/by-hour": {
            "get": {
                "description": "Get the number of votes cast per hour",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "votes"
                ],
                "summary": "Get votes by hour",
                "responses": {
                    "200": {
                        "description": "Votes per hour",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Failed to get votes by hour",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/votes/detailed": {
            "get": {
                "description": "Get detailed results of the voting, including votes per participant",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "votes"
                ],
                "summary": "Get detailed voting results",
                "responses": {
                    "200": {
                        "description": "Detailed voting results",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Failed to get detailed results",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/votes/total": {
            "get": {
                "description": "Get the total number of votes cast",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "votes"
                ],
                "summary": "Get total votes",
                "responses": {
                    "200": {
                        "description": "Total votes",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Failed to get total votes",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Vote": {
            "type": "object",
            "properties": {
                "device_type": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "ip_address": {
                    "type": "string"
                },
                "participant_id": {
                    "type": "string"
                },
                "region": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "user_agent": {
                    "type": "string"
                },
                "voter_id": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http", "https"},
	Title:            "BBB Voting API",
	Description:      "This is a voting API for BBB (Big Brother Brasil)",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
