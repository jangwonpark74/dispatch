///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "VMware Dispatch Identity Manager\n",
    "title": "Identity Manager",
    "contact": {
      "email": "dispatch@vmware.com"
    },
    "version": "1.0.0"
  },
  "basePath": "/",
  "paths": {
    "/": {
      "get": {
        "security": [],
        "summary": "a placeholder root page, no authentication is required for this",
        "operationId": "root",
        "responses": {
          "200": {
            "description": "home page",
            "schema": {
              "$ref": "./models.json#/definitions/Message"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          }
        }
      }
    },
    "/v1/iam/auth": {
      "get": {
        "summary": "handles authorization",
        "operationId": "auth",
        "responses": {
          "202": {
            "description": "default response if authorized",
            "schema": {
              "$ref": "./models.json#/definitions/Message"
            },
            "headers": {
              "X-Dispatch-Org": {
                "type": "string"
              }
            }
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          }
        }
      }
    },
    "/v1/iam/home": {
      "get": {
        "summary": "a placeholder home page",
        "operationId": "home",
        "responses": {
          "200": {
            "description": "home page",
            "schema": {
              "$ref": "./models.json#/definitions/Message"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          }
        }
      }
    },
    "/v1/iam/policy": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "policy"
        ],
        "summary": "List all existing policies",
        "operationId": "getPolicies",
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "./models.json#/definitions/Policy"
              }
            }
          },
          "500": {
            "description": "Internal Error",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "default": {
            "description": "Unexpected Error",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
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
          "policy"
        ],
        "summary": "Add a new policy",
        "operationId": "addPolicy",
        "parameters": [
          {
            "description": "Policy Object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "./models.json#/definitions/Policy"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "created",
            "schema": {
              "$ref": "./models.json#/definitions/Policy"
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "409": {
            "description": "Already Exists",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Error",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          }
        }
      }
    },
    "/v1/iam/policy/{policyName}": {
      "get": {
        "description": "get an Policy by name",
        "produces": [
          "application/json"
        ],
        "tags": [
          "policy"
        ],
        "summary": "Find Policy by name",
        "operationId": "getPolicy",
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "./models.json#/definitions/Policy"
            }
          },
          "400": {
            "description": "Invalid Name supplied",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "404": {
            "description": "Policy not found",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
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
          "policy"
        ],
        "summary": "Update a Policy",
        "operationId": "updatePolicy",
        "parameters": [
          {
            "description": "Policy object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "./models.json#/definitions/Policy"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful update",
            "schema": {
              "$ref": "./models.json#/definitions/Policy"
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "404": {
            "description": "Policy not found",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "policy"
        ],
        "summary": "Deletes an Policy",
        "operationId": "deletePolicy",
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "./models.json#/definitions/Policy"
            }
          },
          "400": {
            "description": "Invalid Name supplied",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "404": {
            "description": "Policy not found",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          }
        }
      },
      "parameters": [
        {
          "pattern": "^[\\w\\d\\-]+$",
          "type": "string",
          "description": "Name of Policy to work on",
          "name": "policyName",
          "in": "path",
          "required": true
        }
      ]
    },
    "/v1/iam/redirect": {
      "get": {
        "summary": "redirect to localhost for vs-cli login (testing)",
        "operationId": "redirect",
        "parameters": [
          {
            "type": "string",
            "description": "the local server url redirecting to",
            "name": "redirect",
            "in": "query"
          }
        ],
        "responses": {
          "302": {
            "description": "redirect",
            "headers": {
              "Location": {
                "type": "string",
                "description": "redirect location"
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          }
        }
      }
    },
    "/v1/iam/serviceaccount": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "serviceaccount"
        ],
        "summary": "List all existing service accounts",
        "operationId": "getServiceAccounts",
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "./models.json#/definitions/ServiceAccount"
              }
            }
          },
          "500": {
            "description": "Internal Error",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "default": {
            "description": "Unexpected Error",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
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
          "serviceaccount"
        ],
        "summary": "Add a new service account",
        "operationId": "addServiceAccount",
        "parameters": [
          {
            "description": "Service Account Object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "./models.json#/definitions/ServiceAccount"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "created",
            "schema": {
              "$ref": "./models.json#/definitions/ServiceAccount"
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "409": {
            "description": "Already Exists",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Error",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          }
        }
      }
    },
    "/v1/iam/serviceaccount/{serviceAccountName}": {
      "get": {
        "description": "get a Service Account by name",
        "produces": [
          "application/json"
        ],
        "tags": [
          "serviceaccount"
        ],
        "summary": "Find Service Account by name",
        "operationId": "getServiceAccount",
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "./models.json#/definitions/ServiceAccount"
            }
          },
          "400": {
            "description": "Invalid Name supplied",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "404": {
            "description": "Service Account not found",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
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
          "serviceaccount"
        ],
        "summary": "Update a Service Account",
        "operationId": "updateServiceAccount",
        "parameters": [
          {
            "description": "Service Account object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "./models.json#/definitions/ServiceAccount"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful update",
            "schema": {
              "$ref": "./models.json#/definitions/ServiceAccount"
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "404": {
            "description": "Service Account not found",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "serviceaccount"
        ],
        "summary": "Deletes an Service Account",
        "operationId": "deleteServiceAccount",
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "./models.json#/definitions/ServiceAccount"
            }
          },
          "400": {
            "description": "Invalid Name supplied",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "404": {
            "description": "Service Account not found",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "./models.json#/definitions/Error"
            }
          }
        }
      },
      "parameters": [
        {
          "pattern": "^[\\w\\d\\-]+$",
          "type": "string",
          "description": "Name of ServiceAccount to work on",
          "name": "serviceAccountName",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "securityDefinitions": {
    "bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    },
    "cookie": {
      "description": "use cookies for authentication, when the user already logged in",
      "type": "apiKey",
      "name": "Cookie",
      "in": "header"
    }
  },
  "security": [
    {
      "cookie": []
    },
    {
      "bearer": []
    }
  ],
  "tags": [
    {
      "name": "authentication"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "VMware Dispatch Identity Manager\n",
    "title": "Identity Manager",
    "contact": {
      "email": "dispatch@vmware.com"
    },
    "version": "1.0.0"
  },
  "basePath": "/",
  "paths": {
    "/": {
      "get": {
        "security": [],
        "summary": "a placeholder root page, no authentication is required for this",
        "operationId": "root",
        "responses": {
          "200": {
            "description": "home page",
            "schema": {
              "$ref": "#/definitions/message"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/iam/auth": {
      "get": {
        "summary": "handles authorization",
        "operationId": "auth",
        "responses": {
          "202": {
            "description": "default response if authorized",
            "schema": {
              "$ref": "#/definitions/message"
            },
            "headers": {
              "X-Dispatch-Org": {
                "type": "string"
              }
            }
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/iam/home": {
      "get": {
        "summary": "a placeholder home page",
        "operationId": "home",
        "responses": {
          "200": {
            "description": "home page",
            "schema": {
              "$ref": "#/definitions/message"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/iam/policy": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "policy"
        ],
        "summary": "List all existing policies",
        "operationId": "getPolicies",
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/policy"
              }
            }
          },
          "500": {
            "description": "Internal Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Unexpected Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
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
          "policy"
        ],
        "summary": "Add a new policy",
        "operationId": "addPolicy",
        "parameters": [
          {
            "description": "Policy Object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/policy"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "created",
            "schema": {
              "$ref": "#/definitions/policy"
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "409": {
            "description": "Already Exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/iam/policy/{policyName}": {
      "get": {
        "description": "get an Policy by name",
        "produces": [
          "application/json"
        ],
        "tags": [
          "policy"
        ],
        "summary": "Find Policy by name",
        "operationId": "getPolicy",
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/policy"
            }
          },
          "400": {
            "description": "Invalid Name supplied",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Policy not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
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
          "policy"
        ],
        "summary": "Update a Policy",
        "operationId": "updatePolicy",
        "parameters": [
          {
            "description": "Policy object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/policy"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful update",
            "schema": {
              "$ref": "#/definitions/policy"
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Policy not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "policy"
        ],
        "summary": "Deletes an Policy",
        "operationId": "deletePolicy",
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/policy"
            }
          },
          "400": {
            "description": "Invalid Name supplied",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Policy not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "pattern": "^[\\w\\d\\-]+$",
          "type": "string",
          "description": "Name of Policy to work on",
          "name": "policyName",
          "in": "path",
          "required": true
        }
      ]
    },
    "/v1/iam/redirect": {
      "get": {
        "summary": "redirect to localhost for vs-cli login (testing)",
        "operationId": "redirect",
        "parameters": [
          {
            "type": "string",
            "description": "the local server url redirecting to",
            "name": "redirect",
            "in": "query"
          }
        ],
        "responses": {
          "302": {
            "description": "redirect",
            "headers": {
              "Location": {
                "type": "string",
                "description": "redirect location"
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/iam/serviceaccount": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "serviceaccount"
        ],
        "summary": "List all existing service accounts",
        "operationId": "getServiceAccounts",
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/serviceAccount"
              }
            }
          },
          "500": {
            "description": "Internal Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Unexpected Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
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
          "serviceaccount"
        ],
        "summary": "Add a new service account",
        "operationId": "addServiceAccount",
        "parameters": [
          {
            "description": "Service Account Object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/serviceAccount"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "created",
            "schema": {
              "$ref": "#/definitions/serviceAccount"
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "409": {
            "description": "Already Exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/iam/serviceaccount/{serviceAccountName}": {
      "get": {
        "description": "get a Service Account by name",
        "produces": [
          "application/json"
        ],
        "tags": [
          "serviceaccount"
        ],
        "summary": "Find Service Account by name",
        "operationId": "getServiceAccount",
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/serviceAccount"
            }
          },
          "400": {
            "description": "Invalid Name supplied",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Service Account not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
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
          "serviceaccount"
        ],
        "summary": "Update a Service Account",
        "operationId": "updateServiceAccount",
        "parameters": [
          {
            "description": "Service Account object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/serviceAccount"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful update",
            "schema": {
              "$ref": "#/definitions/serviceAccount"
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Service Account not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "serviceaccount"
        ],
        "summary": "Deletes an Service Account",
        "operationId": "deleteServiceAccount",
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/serviceAccount"
            }
          },
          "400": {
            "description": "Invalid Name supplied",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Service Account not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "pattern": "^[\\w\\d\\-]+$",
          "type": "string",
          "description": "Name of ServiceAccount to work on",
          "name": "serviceAccountName",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "error": {
      "description": "Error error",
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "description": "code",
          "type": "integer",
          "format": "int64",
          "x-go-name": "Code"
        },
        "message": {
          "description": "message",
          "type": "string",
          "x-go-name": "Message"
        }
      },
      "x-go-package": "github.com/vmware/dispatch/pkg/api/v1"
    },
    "message": {
      "description": "Message message",
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "description": "message",
          "type": "string",
          "x-go-name": "Message"
        }
      },
      "x-go-package": "github.com/vmware/dispatch/pkg/api/v1"
    },
    "policy": {
      "description": "Policy policy",
      "type": "object",
      "required": [
        "name",
        "rules"
      ],
      "properties": {
        "createdTime": {
          "description": "created time",
          "type": "integer",
          "format": "int64",
          "x-go-name": "CreatedTime",
          "readOnly": true
        },
        "id": {
          "description": "id",
          "type": "string",
          "format": "uuid",
          "x-go-name": "ID"
        },
        "kind": {
          "description": "kind",
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$",
          "x-go-name": "Kind",
          "readOnly": true
        },
        "modifiedTime": {
          "description": "modified time",
          "type": "integer",
          "format": "int64",
          "x-go-name": "ModifiedTime",
          "readOnly": true
        },
        "name": {
          "description": "name",
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$",
          "x-go-name": "Name"
        },
        "rules": {
          "description": "rules",
          "type": "array",
          "items": {
            "$ref": "#/definitions/policyRulesItems"
          },
          "x-go-name": "Rules"
        },
        "status": {
          "description": "Status status",
          "type": "string",
          "x-go-package": "github.com/vmware/dispatch/pkg/api/v1"
        }
      },
      "x-go-package": "github.com/vmware/dispatch/pkg/api/v1"
    },
    "policyRulesItems": {
      "description": "Rule rule",
      "type": "object",
      "required": [
        "actions",
        "resources",
        "subjects"
      ],
      "properties": {
        "actions": {
          "description": "actions",
          "type": "array",
          "items": {
            "type": "string"
          },
          "x-go-name": "Actions"
        },
        "resources": {
          "description": "resources",
          "type": "array",
          "items": {
            "type": "string"
          },
          "x-go-name": "Resources"
        },
        "subjects": {
          "description": "subjects",
          "type": "array",
          "items": {
            "type": "string"
          },
          "x-go-name": "Subjects"
        }
      },
      "x-go-gen-location": "models",
      "x-go-package": "github.com/vmware/dispatch/pkg/api/v1"
    },
    "serviceAccount": {
      "description": "ServiceAccount service account",
      "type": "object",
      "required": [
        "name",
        "publicKey"
      ],
      "properties": {
        "createdTime": {
          "description": "created time",
          "type": "integer",
          "format": "int64",
          "x-go-name": "CreatedTime",
          "readOnly": true
        },
        "domain": {
          "description": "domain",
          "type": "string",
          "pattern": "^[\\w\\d\\-\\.]+$",
          "x-go-name": "Domain"
        },
        "id": {
          "description": "id",
          "type": "string",
          "format": "uuid",
          "x-go-name": "ID"
        },
        "kind": {
          "description": "kind",
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$",
          "x-go-name": "Kind",
          "readOnly": true
        },
        "modifiedTime": {
          "description": "modified time",
          "type": "integer",
          "format": "int64",
          "x-go-name": "ModifiedTime",
          "readOnly": true
        },
        "name": {
          "description": "name",
          "type": "string",
          "pattern": "^[\\w\\d\\-\\.]+$",
          "x-go-name": "Name"
        },
        "publicKey": {
          "description": "public key",
          "type": "string",
          "x-go-name": "PublicKey"
        },
        "status": {
          "description": "Status status",
          "type": "string",
          "x-go-package": "github.com/vmware/dispatch/pkg/api/v1"
        }
      },
      "x-go-package": "github.com/vmware/dispatch/pkg/api/v1"
    }
  },
  "securityDefinitions": {
    "bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    },
    "cookie": {
      "description": "use cookies for authentication, when the user already logged in",
      "type": "apiKey",
      "name": "Cookie",
      "in": "header"
    }
  },
  "security": [
    {
      "cookie": []
    },
    {
      "bearer": []
    }
  ],
  "tags": [
    {
      "name": "authentication"
    }
  ]
}`))
}
