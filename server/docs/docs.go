// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Alexander Soldatov",
            "email": "soldatovalex207z@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/Account/Me": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Просмотр информации о текущем авторизованном аккаунте",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AccountController"
                ],
                "summary": "Просмотр данных текущего аккаунта",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    }
                }
            }
        },
        "/api/Account/RefreshTokens": {
            "get": {
                "description": "Обновление токенов и получение access_token в JSON и refresh_token в cookie",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AccountController"
                ],
                "summary": "Обновление токенов",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/httphandlers.RefreshTokens.token"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    }
                }
            }
        },
        "/api/Account/SignIn": {
            "post": {
                "description": "Вход в аккаунт пользователя с использованием имени пользователя - username и паролем - password и получение jwt",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AccountController"
                ],
                "summary": "Вход в аккаунт",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httphandlers.UserSignIn.userCreadentials"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/httphandlers.UserSignIn.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    }
                }
            }
        },
        "/api/Account/SignOut": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Внесение текущего используемого токена доступа в черный список токенов",
                "tags": [
                    "AccountController"
                ],
                "summary": "Выход из аккаунта",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    }
                }
            }
        },
        "/api/Account/SignUp": {
            "post": {
                "description": "Регистрация пользовате и получение jwt",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AccountController"
                ],
                "summary": "Регистрация",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httphandlers.UserSignUp.userData"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/httphandlers.UserSignIn.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    }
                }
            }
        },
        "/api/Room": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Создание новой комнаты и получение ее uuid",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RoomController"
                ],
                "summary": "Создание новой комнаты",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httphandlers.CreateRoom.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    }
                }
            }
<<<<<<< HEAD
=======
        },
        "/api/Tracks": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Создание новой комнаты и получение ее uuid",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TrackController"
                ],
                "summary": "Информация о треках",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    }
                }
            }
>>>>>>> 7b9fc1984b09754619ee44ea4dd07a95ad8762c2
        }
    },
    "definitions": {
        "entities.User": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "httphandlers.CreateRoom.response": {
            "type": "object",
            "properties": {
                "roomId": {
                    "type": "string"
                }
            }
        },
        "httphandlers.RefreshTokens.token": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                }
            }
        },
        "httphandlers.ResponseUser": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "httphandlers.UserSignIn.response": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/httphandlers.ResponseUser"
                }
            }
        },
        "httphandlers.UserSignIn.userCreadentials": {
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
        "httphandlers.UserSignUp.userData": {
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
        "httputils.ResponseError": {
            "type": "object",
            "properties": {
                "err": {
                    "type": "string",
                    "example": "error occures"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:80",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "MTSAudio",
	Description:      "Server for mtsaudio",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
