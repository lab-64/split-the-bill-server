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
        "/api/bill": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bill"
                ],
                "summary": "Create Bill",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.BillInputDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.GeneralResponseDTO"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.BillOutputDTO"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/bill/item": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bill"
                ],
                "summary": "Add Item to Bill",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ItemInputDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.GeneralResponseDTO"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.ItemOutputDTO"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/bill/item/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bill"
                ],
                "summary": "Get Item by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.GeneralResponseDTO"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.ItemOutputDTO"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bill"
                ],
                "summary": "Change Item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Request Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ItemInputDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.GeneralResponseDTO"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.ItemOutputDTO"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/bill/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bill"
                ],
                "summary": "Get Bill by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bill ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.GeneralResponseDTO"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.BillOutputDTO"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/group": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Group"
                ],
                "summary": "Create Group",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.GroupInputDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.GeneralResponseDTO"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.GroupOutputDTO"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/group/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Group"
                ],
                "summary": "Get Group",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Group Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.GeneralResponseDTO"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.GroupOutputDTO"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/invitation": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Invitation"
                ],
                "summary": "Create Group Invitation",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.GroupInvitationInputDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GeneralResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/invitation/accept": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Invitation"
                ],
                "summary": "Accept Group Invitation",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.HandleInvitationInputDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GeneralResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/invitation/decline": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Invitation"
                ],
                "summary": "Decline Group Invitation",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.HandleInvitationInputDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GeneralResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/invitation/user/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Invitation"
                ],
                "summary": "Get All Group Invitations From User",
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
                            "$ref": "#/definitions/dto.GeneralResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/invitation/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Invitation"
                ],
                "summary": "Get Group Invitation By ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Invitation ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GeneralResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/user": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get all Users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.GeneralResponseDTO"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/dto.UserOutputDTO"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/user/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Login User",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CredentialsInputDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GeneralResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/user/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Register User",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserInputDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GeneralResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/user/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get User by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.GeneralResponseDTO"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.UserOutputDTO"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Username",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GeneralResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/user/{username}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get User by username",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Username",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.GeneralResponseDTO"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.UserOutputDTO"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.BillInputDTO": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "groupID": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.ItemInputDTO"
                    }
                },
                "name": {
                    "type": "string"
                },
                "ownerID": {
                    "type": "string"
                }
            }
        },
        "dto.BillOutputDTO": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.ItemOutputDTO"
                    }
                },
                "name": {
                    "type": "string"
                },
                "ownerID": {
                    "type": "string"
                }
            }
        },
        "dto.CredentialsInputDTO": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.GeneralResponseDTO": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "dto.GroupInputDTO": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "ownerID": {
                    "type": "string"
                }
            }
        },
        "dto.GroupInvitationInputDTO": {
            "type": "object",
            "properties": {
                "groupID": {
                    "type": "string"
                },
                "inviteeIDs": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "issuerID": {
                    "type": "string"
                }
            }
        },
        "dto.GroupOutputDTO": {
            "type": "object",
            "properties": {
                "billIDs": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "id": {
                    "type": "string"
                },
                "memberIDs": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "ownerID": {
                    "type": "string"
                }
            }
        },
        "dto.HandleInvitationInputDTO": {
            "type": "object",
            "properties": {
                "invitationID": {
                    "type": "string"
                },
                "issuerID": {
                    "type": "string"
                }
            }
        },
        "dto.ItemInputDTO": {
            "type": "object",
            "properties": {
                "billId": {
                    "type": "string"
                },
                "contributorIDs": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "dto.ItemOutputDTO": {
            "type": "object",
            "properties": {
                "billId": {
                    "type": "string"
                },
                "contributorIDs": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "dto.UserInputDTO": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.UserOutputDTO": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "groups": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.GroupOutputDTO"
                    }
                },
                "id": {
                    "type": "string"
                },
                "invitationIDs": {
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
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Split The Bill API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
