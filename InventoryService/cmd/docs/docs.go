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
        "/ingridient": {
            "put": {
                "description": "Change ingridient in the database by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Inventory manipulation"
                ],
                "summary": "Changes ingridient information",
                "parameters": [
                    {
                        "description": "Item id and it's new data",
                        "name": "item",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ItemModel"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted"
                    }
                }
            },
            "post": {
                "description": "Add new ingridient to database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Inventory manipulation"
                ],
                "summary": "Adds new ingridient",
                "parameters": [
                    {
                        "description": "Item data",
                        "name": "item",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ItemModelCreate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            },
            "delete": {
                "description": "Delete ingridient from database by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Inventory manipulation"
                ],
                "summary": "Deletes ingridient",
                "parameters": [
                    {
                        "type": "string",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted"
                    }
                }
            }
        },
        "/ingridient/search": {
            "get": {
                "description": "Get list of ingridients filtering by it's name and/or parameters",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Inventory manipulation"
                ],
                "summary": "Gets list ingridients by filter",
                "parameters": [
                    {
                        "type": "number",
                        "name": "alcoholmax",
                        "in": "query"
                    },
                    {
                        "type": "number",
                        "name": "alcoholmin",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "name": "isDry",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/inventory": {
            "get": {
                "description": "Get complete list of all ingridients availible for user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Inventory manipulation"
                ],
                "summary": "Gets list of all ingridients",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "do ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Common"
                ],
                "summary": "ping",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ItemModel": {
            "type": "object",
            "properties": {
                "alcohol": {
                    "type": "number"
                },
                "id": {
                    "type": "string"
                },
                "isDry": {
                    "type": "boolean"
                },
                "measurmentUnit": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.ItemModelCreate": {
            "type": "object",
            "properties": {
                "alcohol": {
                    "type": "number"
                },
                "isDry": {
                    "type": "boolean"
                },
                "measurmentUnit": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
