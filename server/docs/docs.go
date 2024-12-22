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
        "/account/login/": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "id to login with",
                        "name": "details",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/queries.Account"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/deck/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deck"
                ],
                "summary": "Get deck by match id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search for deck by match id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.GameDeck"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {}
                    }
                }
            }
        },
        "/match/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "match"
                ],
                "summary": "Get match state by match id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "match id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.MatchDetailsResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {}
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "match"
                ],
                "summary": "Create new match",
                "parameters": [
                    {
                        "description": "MatchRequirements",
                        "name": "details",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vo.MatchRequirements"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/match/cards/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "match"
                ],
                "summary": "Get match cards by match id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "match id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/queries.GetMatchCardsRow"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {}
                    }
                }
            }
        },
        "/match/cut": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "match"
                ],
                "summary": "Cut deck by index of card selected",
                "parameters": [
                    {
                        "description": "Deck index that is to become the cut",
                        "name": "details",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vo.CutDeckReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/match/join": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "match"
                ],
                "summary": "Join match by id",
                "parameters": [
                    {
                        "description": "JoinMatchReq object",
                        "name": "details",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vo.JoinMatchReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.MatchDetailsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/match/player/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "players"
                ],
                "summary": "Get player by barcode",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search for match by id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/queries.Player"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "players"
                ],
                "summary": "Create new player",
                "parameters": [
                    {
                        "description": "player Object to save",
                        "name": "details",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/queries.Player"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
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
                    "players"
                ],
                "summary": "Get player by barcode",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search for match by barcode",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/queries.Player"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/match/player//kitty": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "match"
                ],
                "summary": "Update kitty with ids",
                "parameters": [
                    {
                        "description": "array of ids to add to kitty",
                        "name": "details",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vo.HandModifier"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/queries.Match"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/match/player/ready": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "players"
                ],
                "summary": "Update player by id to be ready. Returns true if all players are ready",
                "parameters": [
                    {
                        "description": "player id to update",
                        "name": "pReady",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/player.PReady"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/matches/open": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "match"
                ],
                "summary": "Get match by id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/vo.GameMatch"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {}
                    }
                }
            }
        },
        "/play": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "match"
                ],
                "summary": "Update play",
                "parameters": [
                    {
                        "description": "HandModifier object",
                        "name": "details",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vo.HandModifier"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/queries.Match"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "pgtype.InfinityModifier": {
            "type": "integer",
            "enum": [
                1,
                0,
                -1
            ],
            "x-enum-varnames": [
                "Infinity",
                "Finite",
                "NegativeInfinity"
            ]
        },
        "pgtype.Int4": {
            "type": "object",
            "properties": {
                "int32": {
                    "type": "integer"
                },
                "valid": {
                    "type": "boolean"
                }
            }
        },
        "pgtype.Timestamptz": {
            "type": "object",
            "properties": {
                "infinityModifier": {
                    "$ref": "#/definitions/pgtype.InfinityModifier"
                },
                "time": {
                    "type": "string"
                },
                "valid": {
                    "type": "boolean"
                }
            }
        },
        "player.PReady": {
            "type": "object",
            "properties": {
                "matchId": {
                    "description": "MatchId",
                    "type": "integer"
                },
                "playerId": {
                    "description": "PlayerId",
                    "type": "integer"
                }
            }
        },
        "queries.Account": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "queries.Card": {
            "type": "object",
            "properties": {
                "art": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "suit": {
                    "$ref": "#/definitions/queries.Cardsuit"
                },
                "value": {
                    "$ref": "#/definitions/queries.Cardvalue"
                }
            }
        },
        "queries.Cardstate": {
            "type": "string",
            "enum": [
                "Deck",
                "Hand",
                "Play",
                "Kitty"
            ],
            "x-enum-varnames": [
                "CardstateDeck",
                "CardstateHand",
                "CardstatePlay",
                "CardstateKitty"
            ]
        },
        "queries.Cardsuit": {
            "type": "string",
            "enum": [
                "Spades",
                "Clubs",
                "Hearts",
                "Diamonds"
            ],
            "x-enum-varnames": [
                "CardsuitSpades",
                "CardsuitClubs",
                "CardsuitHearts",
                "CardsuitDiamonds"
            ]
        },
        "queries.Cardvalue": {
            "type": "string",
            "enum": [
                "Ace",
                "Two",
                "Three",
                "Four",
                "Five",
                "Six",
                "Seven",
                "Eight",
                "Nine",
                "Ten",
                "Jack",
                "Queen",
                "King",
                "Joker"
            ],
            "x-enum-varnames": [
                "CardvalueAce",
                "CardvalueTwo",
                "CardvalueThree",
                "CardvalueFour",
                "CardvalueFive",
                "CardvalueSix",
                "CardvalueSeven",
                "CardvalueEight",
                "CardvalueNine",
                "CardvalueTen",
                "CardvalueJack",
                "CardvalueQueen",
                "CardvalueKing",
                "CardvalueJoker"
            ]
        },
        "queries.Deck": {
            "type": "object",
            "properties": {
                "cutmatchcardid": {
                    "$ref": "#/definitions/pgtype.Int4"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "queries.DeckMatchcard": {
            "type": "object",
            "properties": {
                "deckid": {
                    "type": "integer"
                },
                "matchcardid": {
                    "type": "integer"
                }
            }
        },
        "queries.Gamestate": {
            "type": "string",
            "enum": [
                "NewGameState",
                "JoinGameState",
                "WaitingState",
                "MatchReady",
                "DealState",
                "DiscardState",
                "CutState",
                "PlayState",
                "OpponentState",
                "KittyState",
                "GameWonState",
                "GameLostState",
                "MaxGameState"
            ],
            "x-enum-varnames": [
                "GamestateNewGameState",
                "GamestateJoinGameState",
                "GamestateWaitingState",
                "GamestateMatchReady",
                "GamestateDealState",
                "GamestateDiscardState",
                "GamestateCutState",
                "GamestatePlayState",
                "GamestateOpponentState",
                "GamestateKittyState",
                "GamestateGameWonState",
                "GamestateGameLostState",
                "GamestateMaxGameState"
            ]
        },
        "queries.GetMatchCardsRow": {
            "type": "object",
            "properties": {
                "card": {
                    "$ref": "#/definitions/queries.Card"
                },
                "deck": {
                    "$ref": "#/definitions/queries.Deck"
                },
                "deckMatchcard": {
                    "$ref": "#/definitions/queries.DeckMatchcard"
                },
                "matchcard": {
                    "$ref": "#/definitions/queries.Matchcard"
                }
            }
        },
        "queries.Match": {
            "type": "object",
            "properties": {
                "art": {
                    "type": "string"
                },
                "creationdate": {
                    "$ref": "#/definitions/pgtype.Timestamptz"
                },
                "currentplayerturn": {
                    "type": "integer"
                },
                "cutgamecardid": {
                    "type": "integer"
                },
                "deckid": {
                    "type": "integer"
                },
                "elorangemax": {
                    "type": "integer"
                },
                "elorangemin": {
                    "type": "integer"
                },
                "gamestate": {
                    "$ref": "#/definitions/queries.Gamestate"
                },
                "id": {
                    "type": "integer"
                },
                "privatematch": {
                    "type": "boolean"
                },
                "turnpasstimestamps": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/pgtype.Timestamptz"
                    }
                }
            }
        },
        "queries.Matchcard": {
            "type": "object",
            "properties": {
                "cardid": {
                    "type": "integer"
                },
                "currowner": {
                    "$ref": "#/definitions/pgtype.Int4"
                },
                "id": {
                    "type": "integer"
                },
                "origowner": {
                    "$ref": "#/definitions/pgtype.Int4"
                },
                "state": {
                    "$ref": "#/definitions/queries.Cardstate"
                }
            }
        },
        "queries.Player": {
            "type": "object",
            "properties": {
                "accountid": {
                    "type": "integer"
                },
                "art": {
                    "type": "string"
                },
                "hand": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "isready": {
                    "type": "boolean"
                },
                "kitty": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "play": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "score": {
                    "type": "integer"
                }
            }
        },
        "vo.CutDeckReq": {
            "type": "object",
            "properties": {
                "cutIndex": {
                    "type": "string"
                },
                "matchId": {
                    "type": "integer"
                }
            }
        },
        "vo.GameCard": {
            "type": "object",
            "properties": {
                "art": {
                    "type": "string"
                },
                "cardid": {
                    "type": "integer"
                },
                "currowner": {
                    "$ref": "#/definitions/pgtype.Int4"
                },
                "id": {
                    "type": "integer"
                },
                "origowner": {
                    "$ref": "#/definitions/pgtype.Int4"
                },
                "state": {
                    "$ref": "#/definitions/queries.Cardstate"
                },
                "suit": {
                    "$ref": "#/definitions/queries.Cardsuit"
                },
                "value": {
                    "$ref": "#/definitions/queries.Cardvalue"
                }
            }
        },
        "vo.GameDeck": {
            "type": "object",
            "properties": {
                "cards": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/vo.GameCard"
                    }
                },
                "cutmatchcardid": {
                    "$ref": "#/definitions/pgtype.Int4"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "vo.GameMatch": {
            "type": "object",
            "properties": {
                "art": {
                    "type": "string"
                },
                "creationdate": {
                    "$ref": "#/definitions/pgtype.Timestamptz"
                },
                "currentplayerturn": {
                    "type": "integer"
                },
                "cutgamecardid": {
                    "type": "integer"
                },
                "deckid": {
                    "type": "integer"
                },
                "elorangemax": {
                    "type": "integer"
                },
                "elorangemin": {
                    "type": "integer"
                },
                "gamestate": {
                    "$ref": "#/definitions/queries.Gamestate"
                },
                "id": {
                    "type": "integer"
                },
                "players": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/queries.Player"
                    }
                },
                "privatematch": {
                    "type": "boolean"
                },
                "turnpasstimestamps": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/pgtype.Timestamptz"
                    }
                }
            }
        },
        "vo.HandModifier": {
            "type": "object",
            "properties": {
                "cardIds": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "matchId": {
                    "type": "integer"
                },
                "playerId": {
                    "type": "integer"
                }
            }
        },
        "vo.JoinMatchReq": {
            "type": "object",
            "properties": {
                "accountId": {
                    "type": "integer"
                },
                "matchId": {
                    "type": "integer"
                }
            }
        },
        "vo.MatchDetailsResponse": {
            "type": "object",
            "properties": {
                "gameState": {
                    "$ref": "#/definitions/queries.Gamestate"
                },
                "matchId": {
                    "type": "integer"
                },
                "playerId": {
                    "type": "integer"
                }
            }
        },
        "vo.MatchRequirements": {
            "type": "object",
            "properties": {
                "accountId": {
                    "type": "integer"
                },
                "eloRangeMax": {
                    "type": "integer"
                },
                "eloRangeMin": {
                    "type": "integer"
                },
                "isPrivate": {
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.4",
	Host:             "localhost:1323",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "cribbage server",
	Description:      "cribbage rest server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
