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
            "name": "Andy Ahmedov",
            "url": "https://github.com/andy-ahmedov",
            "email": "andy.ahmedov@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/sign-in": {
            "post": {
                "description": "User authentication by email and password.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "SignIn",
                "operationId": "sign-in",
                "parameters": [
                    {
                        "description": "User info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.SignInInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "The JWT token was successfully generated.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.errResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.errResponse"
                        }
                    }
                }
            }
        },
        "/auth/sign-up": {
            "post": {
                "description": "Registration of a new user in the system.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "SignUp",
                "operationId": "sign-up",
                "parameters": [
                    {
                        "description": "User info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.SignUpInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "The user has been successfully registered.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.errResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.errResponse"
                        }
                    }
                }
            }
        },
        "/books": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Getting all books.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "getAllBooks",
                "operationId": "get-all-books",
                "responses": {
                    "200": {
                        "description": "Books have been successfully received.",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Book"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.errResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Adding a book to the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "CreateBook",
                "operationId": "add-book",
                "parameters": [
                    {
                        "description": "Book information",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Book"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "The data has been successfully written.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.errResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.errResponse"
                        }
                    }
                }
            }
        },
        "/books/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieves a book by ID. If the book is not found, returns an error.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "id"
                ],
                "summary": "GetBookByID",
                "operationId": "get-book-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Book"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.errResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.errResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Updating book data by ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "id"
                ],
                "summary": "updateBook",
                "operationId": "update-book",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Book update information",
                        "name": "updateBook",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UpdateBookInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.errResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.errResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Deleting a book by ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "id"
                ],
                "summary": "deleteBook",
                "operationId": "delete-book",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "The data has been successfully written.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.errResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.errResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Book": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "publish_date": {
                    "type": "string"
                },
                "rating": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "domain.SignInInput": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                }
            }
        },
        "domain.SignUpInput": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "minLength": 2
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                }
            }
        },
        "domain.UpdateBookInput": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "publish_date": {
                    "type": "string"
                },
                "rating": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "rest.errResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
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
	Version:          "1.2",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "CRUD API Service",
	Description:      "Service implementing crud operations",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
