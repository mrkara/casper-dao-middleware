// Package swagger GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package swagger

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "produces": [
        "application/json"
    ],
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
        "/accounts/{address}/aggregated-reputation-changes": {
            "get": {
                "tags": [
                    "Reputation"
                ],
                "summary": "user AggregatedReputationChange",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "10",
                        "description": "Number of items per page",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "ASC",
                            "DESC"
                        ],
                        "type": "string",
                        "default": "ASC",
                        "description": "Sorting direction",
                        "name": "order_direction",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "default": "date",
                        "description": "Comma-separated list of sorting fields (address)",
                        "name": "order_by",
                        "in": "query"
                    },
                    {
                        "maxLength": 66,
                        "type": "string",
                        "description": "Hash or PublicKey",
                        "name": "address",
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
                                    "$ref": "#/definitions/http_response.PaginatedResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/entities.AggregatedReputationChange"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http_response.ErrorResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/http_response.ErrorResult"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http_response.ErrorResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/http_response.ErrorResult"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http_response.ErrorResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/http_response.ErrorResult"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/accounts/{address}/total-reputation": {
            "get": {
                "tags": [
                    "Reputation"
                ],
                "summary": "Calculate address TotalReputation",
                "parameters": [
                    {
                        "maxLength": 66,
                        "type": "string",
                        "description": "Hash or PublicKey",
                        "name": "address",
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
                                    "$ref": "#/definitions/http_response.SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entities.TotalReputation"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http_response.ErrorResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/http_response.ErrorResult"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http_response.ErrorResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/http_response.ErrorResult"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http_response.ErrorResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/http_response.ErrorResult"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/accounts/{address}/votes": {
            "get": {
                "tags": [
                    "Vote"
                ],
                "summary": "Return paginated list of votes for address",
                "parameters": [
                    {
                        "maxLength": 66,
                        "type": "string",
                        "description": "Hash or PublicKey",
                        "name": "address",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Optional fields' schema (voting{})",
                        "name": "includes",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "10",
                        "description": "Number of items per page",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "ASC",
                            "DESC"
                        ],
                        "type": "string",
                        "default": "ASC",
                        "description": "Sorting direction",
                        "name": "order_direction",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "default": "voting_id",
                        "description": "Comma-separated list of sorting fields (voting_id,address)",
                        "name": "order_by",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http_response.PaginatedResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entities.Vote"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http_response.ErrorResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/http_response.ErrorResult"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http_response.ErrorResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/http_response.ErrorResult"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http_response.ErrorResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/http_response.ErrorResult"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/votings": {
            "get": {
                "tags": [
                    "Voting"
                ],
                "summary": "Return paginated list of votings",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "IsFormal flag (boolean)",
                        "name": "is_formal",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "IsActive flag (boolean)",
                        "name": "is_active",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Optional fields' schema (votes_number{}, account_vote(hash))",
                        "name": "includes",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "10",
                        "description": "Number of items per page",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "ASC",
                            "DESC"
                        ],
                        "type": "string",
                        "default": "ASC",
                        "description": "Sorting direction",
                        "name": "order_direction",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "default": "voting_id",
                        "description": "Comma-separated list of sorting fields (voting_id)",
                        "name": "order_by",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http_response.PaginatedResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entities.Voting"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http_response.ErrorResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/http_response.ErrorResult"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http_response.ErrorResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/http_response.ErrorResult"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http_response.ErrorResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/http_response.ErrorResult"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/votings/{voting_id}/votes": {
            "get": {
                "tags": [
                    "Vote"
                ],
                "summary": "Return paginated list of votes for votingID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Comma-separated list of VotingIDs (number)",
                        "name": "voting_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Optional fields' schema (voting{})",
                        "name": "includes",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "10",
                        "description": "Number of items per page",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "ASC",
                            "DESC"
                        ],
                        "type": "string",
                        "default": "ASC",
                        "description": "Sorting direction",
                        "name": "order_direction",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "default": "voting_id",
                        "description": "Comma-separated list of sorting fields (voting_id, address)",
                        "name": "order_by",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http_response.PaginatedResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entities.Vote"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http_response.ErrorResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/http_response.ErrorResult"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http_response.ErrorResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/http_response.ErrorResult"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http_response.ErrorResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/http_response.ErrorResult"
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
        "entities.AggregatedReputationChange": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "earned_amount": {
                    "type": "integer"
                },
                "lost_amount": {
                    "type": "integer"
                },
                "released_amount": {
                    "type": "integer"
                },
                "staked_amount": {
                    "type": "integer"
                },
                "timestamp": {
                    "type": "string"
                },
                "voting_id": {
                    "type": "integer"
                }
            }
        },
        "entities.TotalReputation": {
            "type": "object",
            "properties": {
                "available_amount": {
                    "type": "integer"
                },
                "staked_amount": {
                    "type": "integer"
                }
            }
        },
        "entities.Vote": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "amount": {
                    "type": "integer"
                },
                "deploy_hash": {
                    "type": "string"
                },
                "is_in_favour": {
                    "type": "boolean"
                },
                "timestamp": {
                    "type": "string"
                },
                "voting_id": {
                    "type": "integer"
                }
            }
        },
        "entities.Voting": {
            "type": "object",
            "properties": {
                "creator": {
                    "type": "string"
                },
                "deploy_hash": {
                    "type": "string"
                },
                "has_ended": {
                    "type": "boolean"
                },
                "informal_voting_id": {
                    "type": "integer"
                },
                "is_formal": {
                    "type": "boolean"
                },
                "timestamp": {
                    "type": "string"
                },
                "voting_id": {
                    "type": "integer"
                },
                "voting_quorum": {
                    "type": "integer"
                },
                "voting_time": {
                    "type": "integer"
                }
            }
        },
        "http_response.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/http_response.ErrorResult"
                }
            }
        },
        "http_response.ErrorResult": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "http_response.PaginatedResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "item_count": {
                    "type": "integer"
                },
                "page_count": {
                    "type": "integer"
                }
            }
        },
        "http_response.SuccessResponse": {
            "type": "object",
            "properties": {
                "data": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Casper-CRDao API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
