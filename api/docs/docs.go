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
        "/api/v1": {
            "post": {
                "description": "Create short url",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create short url",
                "parameters": [
                    {
                        "description": "Query Parameters",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/docsobj.Request"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "docsobj.Request": {
            "type": "object",
            "properties": {
                "short": {
                    "description": "Custom Short Endpoint (Automatically generated if not provided)",
                    "type": "string"
                },
                "url": {
                    "description": "Input Long URL",
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
	BasePath:         "/",
	Schemes:          []string{"http", "https"},
	Title:            "ZipLink API",
	Description:      "## About\n\n`ZipLink` is an URL shortener created by [Sachin Kant](https://github.com/sachin-404)\n\n- Source Code: <https://github.com/sachin-404/ZipLink> ",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
