// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "117503445",
            "url": "http://www.117503445.top",
            "email": "t117503445@gmail.com"
        },
        "license": {
            "name": "GNU GPL 3.0",
            "url": "https://github.com/TGclub/smartlab_backend_go/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/Bulletin": {
            "get": {
                "description": "读取所有公告",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "BulletinReadAll",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.BulletinOut"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "创建公告，需要管理员权限。",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "BulletinCreate",
                "parameters": [
                    {
                        "description": "dto.BulletinIn",
                        "name": "BulletinIn",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.BulletinIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.BulletinOut"
                            }
                        }
                    }
                }
            }
        },
        "/Bulletin/{id}": {
            "get": {
                "description": "读取指定公告",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "BulletinRead",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Bulletin.ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.BulletinOut"
                            }
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "修改指定公告",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "BulletinUpdate",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Bulletin.ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "dto.BulletinIn",
                        "name": "BulletinIn",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.BulletinIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.BulletinOut"
                            }
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "删除公告，需要管理员权限。",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "BulletinDelete",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "DeleteBulletin.ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.BulletinOut"
                            }
                        }
                    }
                }
            }
        },
        "/ping": {
            "post": {
                "description": "返回 pong",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "状态检查",
                "operationId": "ping",
                "responses": {
                    "200": {
                        "description": "pong",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "登陆账户，返回 JWT",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "dto.userLoginIn",
                        "name": "userLoginIn",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserLoginIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.UserOut"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.BulletinIn": {
            "type": "object",
            "properties": {
                "imageUrl": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "dto.BulletinOut": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "imageUrl": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "dto.UserLoginIn": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.UserOut": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "role": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "JWT": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    },
    "x-extension-openapi": {
        "example": "value on a json format"
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost",
	BasePath:    "/api",
	Schemes:     []string{"http"},
	Title:       "smartlab",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
