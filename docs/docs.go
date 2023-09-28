// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "My linkedin profile",
            "url": "https://www.linkedin.com/in/joaoschoen/",
            "email": "joaoschoen@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user": {
            "put": {
                "description": "Receives updated user object, returns updated object",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.PostUserResponse"
                        }
                    },
                    "400": {
                        "description": "Email already in use"
                    },
                    "404": {
                        "description": "User not found."
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            },
            "post": {
                "description": "Receives user email and password, returns UUID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create new user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User email",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.PostUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "422": {
                        "description": "Email already in use"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            },
            "delete": {
                "description": "Receives user ID, returns deleted ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.DeleteUserResponse"
                        }
                    },
                    "404": {
                        "description": "User not found."
                    }
                }
            }
        },
        "/user/list": {
            "get": {
                "description": "Can receive email as a query filter\nThis route is paged, it requrires a page number to operate if none is received, will return page 0",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get user list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GetUserListResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "description": "Receives ID by request param and retreives user data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get user data",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GetUserResponse"
                        }
                    },
                    "404": {
                        "description": "User not found."
                    }
                }
            }
        }
    },
    "definitions": {
        "model.DeleteUserResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "model.GetUserListResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SafeUser"
                    }
                },
                "paging": {
                    "$ref": "#/definitions/model.Paging"
                }
            }
        },
        "model.GetUserResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/model.SafeUser"
                }
            }
        },
        "model.Paging": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "model.PostUserResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "model.SafeUser": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "REST Echo Base",
	Description:      "This is an API base without database interaction",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
