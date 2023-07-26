// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://tmpf.me/terms/",
        "contact": {
            "name": "API Support",
            "url": "https://tmpf.me/credit",
            "email": "dev@tmpf.me"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "get the status of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "root"
                ],
                "summary": "Show the status of server.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/checkpw/{id}": {
            "get": {
                "description": "Check password of file item.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "summary": "Check password of file item.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "file id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "pw",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/dl/{id}": {
            "get": {
                "description": "Download file item.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "summary": "Download file item.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "file id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/file/new": {
            "post": {
                "description": "Upload file item.",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "summary": "Upload file item.",
                "parameters": [
                    {
                        "type": "file",
                        "description": "file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "pw",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "download limit",
                        "name": "X-Download-Limit",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "time limit",
                        "name": "X-Time-Limit",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/file/{id}": {
            "get": {
                "description": "Get file item.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "summary": "Get file item.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "file id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete file item.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "summary": "Delete file item.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "file id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/files": {
            "get": {
                "description": "List file items.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "summary": "List file items.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/database.FileTracking"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/info": {
            "get": {
                "description": "get the information of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "root"
                ],
                "summary": "Show the information of server.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "api name",
                        "name": "api",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/text/new": {
            "post": {
                "description": "Upload text item.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "text"
                ],
                "summary": "Upload text item.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "download limit",
                        "name": "X-Download-Limit",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "time limit",
                        "name": "X-Time-Limit",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/text/{id}": {
            "get": {
                "description": "Download text item.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "text"
                ],
                "summary": "Download text item.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "text id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete text item.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "text"
                ],
                "summary": "Delete text item.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "text id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/texts": {
            "get": {
                "description": "List text items.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "text"
                ],
                "summary": "List text items.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "database.FileTracking": {
            "type": "object",
            "properties": {
                "downloadCount": {
                    "type": "integer"
                },
                "downloadLimit": {
                    "type": "integer"
                },
                "expireTime": {
                    "type": "string"
                },
                "fileId": {
                    "type": "string"
                },
                "filename": {
                    "type": "string"
                },
                "isEncrypted": {
                    "type": "boolean"
                },
                "size": {
                    "type": "integer"
                },
                "uploadDate": {
                    "type": "string"
                }
            }
        },
        "utils.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "description": "\"요청이 정상적으로 처리되었습니다\" | \"서버에서 일시적인 오류가 발생했어요\"",
                    "type": "string"
                },
                "status": {
                    "description": "SUCCESS | FAIL",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Tempfiles API",
	Description:      "This is a Tempfiles server for file and text sharing.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
