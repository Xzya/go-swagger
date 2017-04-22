package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a sample server Petstore server.\n\n[Learn about Swagger](http://swagger.wordnik.com) or join the IRC channel '#swagger' on irc.freenode.net.\n\nFor this sample, you can use the api key 'special-key' to test the authorization filters\n",
    "title": "Swagger Petstore",
    "termsOfService": "http://helloreverb.com/terms/",
    "contact": {
      "name": "apiteam@wordnik.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "petstore.swagger.wordnik.com",
  "basePath": "/v2",
  "paths": {
    "/pets": {
      "put": {
        "consumes": [
          "application/json",
          "application/xml"
        ],
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Update an existing pet",
        "operationId": "updatePet",
        "security": [
          {
            "petstore_auth": [
              "write_pets",
              "read_pets"
            ]
          }
        ],
        "parameters": [
          {
            "description": "Pet object that needs to be added to the store",
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Pet"
            }
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Pet not found"
          },
          "405": {
            "description": "Validation exception"
          }
        }
      },
      "post": {
        "consumes": [
          "application/json",
          "application/xml"
        ],
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Add a new pet to the store",
        "operationId": "addPet",
        "security": [
          {
            "petstore_auth": [
              "write_pets",
              "read_pets"
            ]
          }
        ],
        "parameters": [
          {
            "description": "Pet object that needs to be added to the store",
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Pet"
            }
          }
        ],
        "responses": {
          "405": {
            "description": "Invalid input"
          }
        }
      }
    },
    "/pets/findByStatus": {
      "get": {
        "description": "Multiple status values can be provided with comma seperated strings",
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Finds Pets by status",
        "operationId": "findPetsByStatus",
        "security": [
          {
            "petstore_auth": [
              "write_pets",
              "read_pets"
            ]
          }
        ],
        "parameters": [
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi",
            "description": "Status values that need to be considered for filter",
            "name": "status",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Pet"
              }
            }
          },
          "400": {
            "description": "Invalid status value"
          }
        }
      }
    },
    "/pets/findByTags": {
      "get": {
        "description": "Muliple tags can be provided with comma seperated strings. Use tag1, tag2, tag3 for testing.",
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Finds Pets by tags",
        "operationId": "findPetsByTags",
        "security": [
          {
            "petstore_auth": [
              "write_pets",
              "read_pets"
            ]
          }
        ],
        "parameters": [
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi",
            "description": "Tags to filter by",
            "name": "tags",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Pet"
              }
            }
          },
          "400": {
            "description": "Invalid tag value"
          }
        }
      }
    },
    "/pets/{petId}": {
      "get": {
        "description": "Returns a pet when ID \u003c 10.  ID \u003e 10 or nonintegers will simulate API error conditions",
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Find pet by ID",
        "operationId": "getPetById",
        "security": [
          {
            "api_key": []
          },
          {
            "petstore_auth": [
              "write_pets",
              "read_pets"
            ]
          }
        ],
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of pet that needs to be fetched",
            "name": "petId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Pet"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Pet not found"
          }
        }
      },
      "post": {
        "consumes": [
          "application/x-www-form-urlencoded"
        ],
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Updates a pet in the store with form data",
        "operationId": "updatePetWithForm",
        "security": [
          {
            "petstore_auth": [
              "write_pets",
              "read_pets"
            ]
          }
        ],
        "parameters": [
          {
            "type": "string",
            "description": "ID of pet that needs to be updated",
            "name": "petId",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "Updated name of the pet",
            "name": "name",
            "in": "formData",
            "required": true
          },
          {
            "type": "string",
            "description": "Updated status of the pet",
            "name": "status",
            "in": "formData",
            "required": true
          }
        ],
        "responses": {
          "405": {
            "description": "Invalid input"
          }
        }
      },
      "delete": {
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Deletes a pet",
        "operationId": "deletePet",
        "security": [
          {
            "petstore_auth": [
              "write_pets",
              "read_pets"
            ]
          }
        ],
        "parameters": [
          {
            "type": "string",
            "name": "api_key",
            "in": "header",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "Pet id to delete",
            "name": "petId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid pet value"
          }
        }
      }
    },
    "/stores/order": {
      "post": {
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "store"
        ],
        "summary": "Place an order for a pet",
        "operationId": "placeOrder",
        "parameters": [
          {
            "description": "order placed for purchasing the pet",
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Order"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Order"
            }
          },
          "400": {
            "description": "Invalid Order"
          }
        }
      }
    },
    "/stores/order/{orderId}": {
      "get": {
        "description": "For valid response try integer IDs with value \u003c= 5 or \u003e 10. Other values will generated exceptions",
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "store"
        ],
        "summary": "Find purchase order by ID",
        "operationId": "getOrderById",
        "parameters": [
          {
            "type": "string",
            "description": "ID of pet that needs to be fetched",
            "name": "orderId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Order"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Order not found"
          }
        }
      },
      "delete": {
        "description": "For valid response try integer IDs with value \u003c 1000. Anything above 1000 or nonintegers will generate API errors",
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "store"
        ],
        "summary": "Delete purchase order by ID",
        "operationId": "deleteOrder",
        "parameters": [
          {
            "type": "string",
            "description": "ID of the order that needs to be deleted",
            "name": "orderId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Order not found"
          }
        }
      }
    },
    "/users": {
      "post": {
        "description": "This can only be done by the logged in user.",
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "user"
        ],
        "summary": "Create user",
        "operationId": "createUser",
        "parameters": [
          {
            "description": "Created user object",
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "default": {
            "description": "successful operation"
          }
        }
      }
    },
    "/users/createWithArray": {
      "post": {
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "user"
        ],
        "summary": "Creates list of users with given input array",
        "operationId": "createUsersWithArrayInput",
        "parameters": [
          {
            "description": "List of user object",
            "name": "body",
            "in": "body",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/User"
              }
            }
          }
        ],
        "responses": {
          "default": {
            "description": "successful operation"
          }
        }
      }
    },
    "/users/createWithList": {
      "post": {
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "user"
        ],
        "summary": "Creates list of users with given input array",
        "operationId": "createUsersWithListInput",
        "parameters": [
          {
            "description": "List of user object",
            "name": "body",
            "in": "body",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/User"
              }
            }
          }
        ],
        "responses": {
          "default": {
            "description": "successful operation"
          }
        }
      }
    },
    "/users/login": {
      "get": {
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "user"
        ],
        "summary": "Logs user into the system",
        "operationId": "loginUser",
        "parameters": [
          {
            "type": "string",
            "description": "The user name for login",
            "name": "username",
            "in": "query"
          },
          {
            "type": "string",
            "description": "The password for login in clear text",
            "name": "password",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "string"
            }
          },
          "400": {
            "description": "Invalid username/password supplied"
          }
        }
      }
    },
    "/users/logout": {
      "get": {
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "user"
        ],
        "summary": "Logs out current logged in user session",
        "operationId": "logoutUser",
        "responses": {
          "default": {
            "description": "successful operation"
          }
        }
      }
    },
    "/users/{username}": {
      "get": {
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "user"
        ],
        "summary": "Get user by user name",
        "operationId": "getUserByName",
        "parameters": [
          {
            "type": "string",
            "description": "The name that needs to be fetched. Use user1 for testing.",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Invalid username supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      },
      "put": {
        "description": "This can only be done by the logged in user.",
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "user"
        ],
        "summary": "Updated user",
        "operationId": "updateUser",
        "parameters": [
          {
            "type": "string",
            "description": "name that need to be deleted",
            "name": "username",
            "in": "path",
            "required": true
          },
          {
            "description": "Updated user object",
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid user supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      },
      "delete": {
        "description": "This can only be done by the logged in user.",
        "produces": [
          "application/json",
          "application/xml"
        ],
        "tags": [
          "user"
        ],
        "summary": "Delete user",
        "operationId": "deleteUser",
        "parameters": [
          {
            "type": "string",
            "description": "The name that needs to be deleted",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid username supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      }
    }
  },
  "definitions": {
    "Category": {
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "Order": {
      "properties": {
        "complete": {
          "type": "boolean"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "petId": {
          "type": "integer",
          "format": "int64"
        },
        "quantity": {
          "type": "integer",
          "format": "int32"
        },
        "shipDate": {
          "type": "string",
          "format": "date-time"
        },
        "status": {
          "description": "Order Status",
          "type": "string"
        }
      }
    },
    "Pet": {
      "required": [
        "name",
        "photoUrls"
      ],
      "properties": {
        "category": {
          "$ref": "#/definitions/Category"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string",
          "example": "doggie"
        },
        "photoUrls": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "status": {
          "description": "pet status in the store",
          "type": "string"
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Tag"
          }
        }
      }
    },
    "Tag": {
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "User": {
      "properties": {
        "email": {
          "type": "string"
        },
        "firstName": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "lastName": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "phone": {
          "type": "string"
        },
        "userStatus": {
          "description": "User Status",
          "type": "integer",
          "format": "int32"
        },
        "username": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "api_key": {
      "type": "apiKey",
      "name": "api_key",
      "in": "header"
    },
    "petstore_auth": {
      "type": "oauth2",
      "flow": "implicit",
      "authorizationUrl": "http://petstore.swagger.wordnik.com/api/oauth/dialog",
      "scopes": {
        "read_pets": "read your pets",
        "write_pets": "modify pets in your account"
      }
    }
  }
}`))
}
