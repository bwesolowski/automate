package api

func init() {
	Swagger.Add("infra_proxy", `{
  "swagger": "2.0",
  "info": {
    "title": "external/infra_proxy/infra_proxy.proto",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/api/v0/infra/servers": {
      "get": {
        "operationId": "InfraProxy_GetServers",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.GetServers"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "tags": [
          "InfraProxy"
        ]
      },
      "post": {
        "operationId": "InfraProxy_CreateServer",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CreateServer"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.request.CreateServer"
            }
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{id}": {
      "get": {
        "operationId": "InfraProxy_GetServer",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.GetServer"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      },
      "delete": {
        "operationId": "InfraProxy_DeleteServer",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.DeleteServer"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      },
      "put": {
        "operationId": "InfraProxy_UpdateServer",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.UpdateServer"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.request.UpdateServer"
            }
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/orgs": {
      "get": {
        "operationId": "InfraProxy_GetOrgs",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.GetOrgs"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      },
      "post": {
        "operationId": "InfraProxy_CreateOrg",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CreateOrg"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.request.CreateOrg"
            }
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/orgs/{id}": {
      "get": {
        "operationId": "InfraProxy_GetOrg",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.GetOrg"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      },
      "delete": {
        "operationId": "InfraProxy_DeleteOrg",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.DeleteOrg"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      },
      "put": {
        "operationId": "InfraProxy_UpdateOrg",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.UpdateOrg"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.request.UpdateOrg"
            }
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/orgs/{id}/reset-key": {
      "put": {
        "operationId": "InfraProxy_ResetOrgAdminKey",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.ResetOrgAdminKey"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.request.ResetOrgAdminKey"
            }
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/orgs/{org_id}/clients": {
      "get": {
        "operationId": "InfraProxy_GetClients",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Clients"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "search_query.q",
            "description": "The search query used to identify a list of items.",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "search_query.page",
            "description": "Starting page for the results.",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "search_query.per_page",
            "description": "The number of results on each page.",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      },
      "post": {
        "operationId": "InfraProxy_CreateClient",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CreateClient"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.request.CreateClient"
            }
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/orgs/{org_id}/clients/{name}": {
      "get": {
        "operationId": "InfraProxy_GetClient",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Client"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Client name.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      },
      "delete": {
        "operationId": "InfraProxy_DeleteClient",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Client"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Client name.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/orgs/{org_id}/clients/{name}/reset": {
      "put": {
        "operationId": "InfraProxy_ResetClientKey",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.ResetClient"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Client name.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.request.ClientKey"
            }
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/orgs/{org_id}/cookbooks": {
      "get": {
        "operationId": "InfraProxy_GetCookbooks",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Cookbooks"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "ID of the Server",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "ID of the Org.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/orgs/{org_id}/cookbooks/{name}": {
      "get": {
        "operationId": "InfraProxy_GetCookbookVersions",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookVersions"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "ID of the Server.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "ID of the Org.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Name of the cookbook.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/orgs/{org_id}/cookbooks/{name}/{version}": {
      "get": {
        "operationId": "InfraProxy_GetCookbook",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Cookbook"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "ID of the Server.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "ID of the Org.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Name of the cookbook.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "version",
            "description": "Version of the cookbook.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/orgs/{org_id}/cookbooks/{name}/{version}/file-content": {
      "get": {
        "operationId": "InfraProxy_GetCookbookFileContent",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookFileContent"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "ID of the server.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "ID of the org.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Name of the cookbook.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "version",
            "description": "Version of the cookbook.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "url",
            "description": "Cookbook data file URL.",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/orgs/{org_id}/data_bags": {
      "get": {
        "operationId": "InfraProxy_GetDataBags",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.DataBags"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      },
      "post": {
        "operationId": "InfraProxy_CreateDataBag",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CreateDataBag"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.request.CreateDataBag"
            }
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/orgs/{org_id}/data_bags/{name}": {
      "get": {
        "operationId": "InfraProxy_GetDataBagItems",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.DataBagItems"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Data bag name.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "search_query.q",
            "description": "The search query used to identify a list of items.",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "search_query.page",
            "description": "Starting page for the results.",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "search_query.per_page",
            "description": "The number of results on each page.",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      },
      "delete": {
        "operationId": "InfraProxy_DeleteDataBag",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.DataBag"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Data bag name.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      },
      "post": {
        "operationId": "InfraProxy_CreateDataBagItem",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CreateDataBagItem"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Data bag name.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.request.CreateDataBagItem"
            }
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/orgs/{org_id}/data_bags/{name}/{item_id}": {
      "put": {
        "operationId": "InfraProxy_UpdateDataBagItem",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.UpdateDataBagItem"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Data bag name.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "item_id",
            "description": "Data bag item ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.request.UpdateDataBagItem"
            }
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/orgs/{org_id}/data_bags/{name}/{item}": {
      "get": {
        "operationId": "InfraProxy_GetDataBagItem",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.DataBagItem"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Data bag name.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "item",
            "description": "Data bag item name.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      },
      "delete": {
        "operationId": "InfraProxy_DeleteDataBagItem",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.DataBagItem"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Data bag name.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "item",
            "description": "Data bag item name.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/orgs/{org_id}/environments": {
      "get": {
        "operationId": "InfraProxy_GetEnvironments",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Environments"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "search_query.q",
            "description": "The search query used to identify a list of items.",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "search_query.page",
            "description": "Starting page for the results.",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "search_query.per_page",
            "description": "The number of results on each page.",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      },
      "post": {
        "operationId": "InfraProxy_CreateEnvironment",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Environment"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.request.CreateEnvironment"
            }
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/orgs/{org_id}/environments/{name}": {
      "get": {
        "operationId": "InfraProxy_GetEnvironment",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Environment"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Environment name.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      },
      "delete": {
        "operationId": "InfraProxy_DeleteEnvironment",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Environment"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Environment name.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      },
      "put": {
        "operationId": "InfraProxy_UpdateEnvironment",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Environment"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Environment name.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.request.UpdateEnvironment"
            }
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/orgs/{org_id}/environments/{name}/recipes": {
      "get": {
        "operationId": "InfraProxy_GetEnvironmentRecipes",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.EnvironmentRecipesList"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Environment name.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/orgs/{org_id}/node/{name}/runlist/{environment}": {
      "get": {
        "operationId": "InfraProxy_GetNodeExpandedRunList",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.NodeExpandedRunList"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Node name.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "environment",
            "description": "Node environment.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/orgs/{org_id}/nodes": {
      "get": {
        "operationId": "InfraProxy_GetNodes",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Nodes"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "search_query.q",
            "description": "The search query used to identify a list of items.",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "search_query.page",
            "description": "Starting page for the results.",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "search_query.per_page",
            "description": "The number of results on each page.",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/orgs/{org_id}/nodes/{name}": {
      "get": {
        "operationId": "InfraProxy_GetNode",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Node"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Node name.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      },
      "delete": {
        "operationId": "InfraProxy_DeleteNode",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.DeleteNode"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Node name.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      },
      "put": {
        "operationId": "InfraProxy_UpdateNode",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Node"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Node name.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.request.NodeDetails"
            }
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/orgs/{org_id}/nodes/{name}/attributes": {
      "put": {
        "operationId": "InfraProxy_UpdateNodeAttributes",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.UpdateNodeAttributes"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Node name.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.request.UpdateNodeAttributes"
            }
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/orgs/{org_id}/nodes/{name}/environment": {
      "put": {
        "operationId": "InfraProxy_UpdateNodeEnvironment",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.UpdateNodeEnvironment"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Node name.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.request.UpdateNodeEnvironment"
            }
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/orgs/{org_id}/nodes/{name}/tags": {
      "put": {
        "operationId": "InfraProxy_UpdateNodeTags",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.UpdateNodeTags"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Node name.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.request.UpdateNodeTags"
            }
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/orgs/{org_id}/policyfiles": {
      "get": {
        "operationId": "InfraProxy_GetPolicyfiles",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Policyfiles"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef Organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/orgs/{org_id}/policyfiles/{name}": {
      "get": {
        "operationId": "InfraProxy_GetPolicyfile",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Policyfile"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef Organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Policyfile name.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "revision_id",
            "description": "Policyfile revision ID.",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      },
      "delete": {
        "operationId": "InfraProxy_DeletePolicyfile",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.DeletePolicyfile"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef Organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Policyfile name.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/orgs/{org_id}/roles": {
      "get": {
        "operationId": "InfraProxy_GetRoles",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Roles"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "search_query.q",
            "description": "The search query used to identify a list of items.",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "search_query.page",
            "description": "Starting page for the results.",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "search_query.per_page",
            "description": "The number of results on each page.",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      },
      "post": {
        "operationId": "InfraProxy_CreateRole",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Role"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.request.CreateRole"
            }
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/orgs/{org_id}/roles/{name}": {
      "get": {
        "operationId": "InfraProxy_GetRole",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Role"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Role name.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      },
      "delete": {
        "operationId": "InfraProxy_DeleteRole",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Role"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Role name.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      },
      "put": {
        "operationId": "InfraProxy_UpdateRole",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Role"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Role name.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.request.UpdateRole"
            }
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/orgs/{org_id}/roles/{name}/environments": {
      "get": {
        "operationId": "InfraProxy_GetRoleEnvironments",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.RoleEnvironments"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Role name.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/api/v0/infra/servers/{server_id}/orgs/{org_id}/roles/{name}/runlist/{environment}": {
      "get": {
        "operationId": "InfraProxy_GetRoleExpandedRunList",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.ExpandedRunList"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "server_id",
            "description": "Chef Infra Server ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "org_id",
            "description": "Chef organization ID.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "description": "Role name.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "environment",
            "description": "Role environment.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "InfraProxy"
        ]
      }
    }
  },
  "definitions": {
    "chef.automate.api.infra_proxy.request.ClientKey": {
      "type": "object",
      "properties": {
        "org_id": {
          "type": "string",
          "description": "Chef organization ID."
        },
        "server_id": {
          "type": "string",
          "description": "Chef Infra Server ID."
        },
        "name": {
          "type": "string",
          "description": "Client name."
        },
        "key": {
          "type": "string",
          "description": "Client key name."
        }
      }
    },
    "chef.automate.api.infra_proxy.request.CreateClient": {
      "type": "object",
      "properties": {
        "org_id": {
          "type": "string",
          "description": "Chef organization ID."
        },
        "server_id": {
          "type": "string",
          "description": "Chef Infra Server ID."
        },
        "name": {
          "type": "string",
          "description": "Client name."
        },
        "validator": {
          "type": "boolean",
          "format": "boolean",
          "description": "Boolean indicates client type is validator or not."
        },
        "create_key": {
          "type": "boolean",
          "format": "boolean",
          "description": "Boolean indicates whether it's required to create a key or not."
        }
      }
    },
    "chef.automate.api.infra_proxy.request.CreateDataBag": {
      "type": "object",
      "properties": {
        "org_id": {
          "type": "string",
          "description": "Chef organization ID."
        },
        "server_id": {
          "type": "string",
          "description": "Chef Infra Server ID."
        },
        "name": {
          "type": "string",
          "description": "Data bag name."
        }
      }
    },
    "chef.automate.api.infra_proxy.request.CreateDataBagItem": {
      "type": "object",
      "properties": {
        "org_id": {
          "type": "string",
          "description": "Chef organization ID."
        },
        "server_id": {
          "type": "string",
          "description": "Chef Infra Server ID."
        },
        "name": {
          "type": "string",
          "description": "Data bag name."
        },
        "data": {
          "type": "object",
          "description": "Data bag item JSON data."
        }
      }
    },
    "chef.automate.api.infra_proxy.request.CreateEnvironment": {
      "type": "object",
      "properties": {
        "org_id": {
          "type": "string",
          "description": "Chef organization ID."
        },
        "server_id": {
          "type": "string",
          "description": "Chef Infra Server ID."
        },
        "name": {
          "type": "string",
          "description": "Environment name."
        },
        "description": {
          "type": "string",
          "description": "Environment description."
        },
        "json_class": {
          "type": "string",
          "description": "Class name."
        },
        "cookbook_versions": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "description": "Environment versioned cookbooks constraints."
        },
        "default_attributes": {
          "type": "object",
          "description": "Environment default attributes JSON."
        },
        "override_attributes": {
          "type": "object",
          "description": "Environment override attributes JSON."
        }
      }
    },
    "chef.automate.api.infra_proxy.request.CreateOrg": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Chef organization ID."
        },
        "name": {
          "type": "string",
          "description": "Chef organization name."
        },
        "admin_user": {
          "type": "string",
          "description": "Chef organization admin user."
        },
        "admin_key": {
          "type": "string",
          "description": "Chef organization admin key."
        },
        "server_id": {
          "type": "string",
          "description": "Chef Infra Server ID."
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of projects this chef organization belongs to. May be empty."
        }
      }
    },
    "chef.automate.api.infra_proxy.request.CreateRole": {
      "type": "object",
      "properties": {
        "org_id": {
          "type": "string",
          "description": "Chef organization ID."
        },
        "server_id": {
          "type": "string",
          "description": "Chef Infra Server ID."
        },
        "name": {
          "type": "string",
          "description": "Role name."
        },
        "description": {
          "type": "string",
          "description": "Role description."
        },
        "default_attributes": {
          "type": "object",
          "description": "Role default attributes JSON."
        },
        "override_attributes": {
          "type": "object",
          "description": "Role override attributes JSON."
        },
        "run_list": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Role run list."
        },
        "env_run_lists": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.request.EnvRunList"
          },
          "description": "Environment based run list."
        }
      }
    },
    "chef.automate.api.infra_proxy.request.CreateServer": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Chef Infra Server ID."
        },
        "name": {
          "type": "string",
          "description": "Chef Infra Server name."
        },
        "fqdn": {
          "type": "string",
          "description": "Chef Infra Server FQDN."
        },
        "ip_address": {
          "type": "string",
          "description": "Chef Infra Server IP address."
        }
      }
    },
    "chef.automate.api.infra_proxy.request.EnvRunList": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Environment name."
        },
        "run_list": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Role run list."
        }
      }
    },
    "chef.automate.api.infra_proxy.request.NodeDetails": {
      "type": "object",
      "properties": {
        "org_id": {
          "type": "string",
          "description": "Chef organization ID."
        },
        "server_id": {
          "type": "string",
          "description": "Chef Infra Server ID."
        },
        "name": {
          "type": "string",
          "description": "Node name."
        },
        "environment": {
          "type": "string",
          "description": "Node environment."
        },
        "policy_name": {
          "type": "string",
          "description": "Node policy name."
        },
        "policy_group": {
          "type": "string",
          "description": "Node policy group."
        },
        "run_list": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Node run-list."
        },
        "automatic_attributes": {
          "type": "object",
          "description": "Node automatic attributes JSON."
        },
        "default_attributes": {
          "type": "object",
          "description": "Node default attributes JSON."
        },
        "normal_attributes": {
          "type": "object",
          "description": "Node normal attributes JSON."
        },
        "override_attributes": {
          "type": "object",
          "description": "Node override attributes JSON."
        }
      }
    },
    "chef.automate.api.infra_proxy.request.ResetOrgAdminKey": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Chef organization ID."
        },
        "server_id": {
          "type": "string",
          "description": "Chef Infra Server ID."
        },
        "admin_key": {
          "type": "string",
          "description": "Chef organization admin key."
        }
      }
    },
    "chef.automate.api.infra_proxy.request.SearchQuery": {
      "type": "object",
      "properties": {
        "q": {
          "type": "string",
          "description": "The search query used to identify a list of items."
        },
        "page": {
          "type": "integer",
          "format": "int32",
          "description": "Starting page for the results."
        },
        "per_page": {
          "type": "integer",
          "format": "int32",
          "description": "The number of results on each page."
        }
      }
    },
    "chef.automate.api.infra_proxy.request.UpdateDataBagItem": {
      "type": "object",
      "properties": {
        "org_id": {
          "type": "string",
          "description": "Chef organization ID."
        },
        "server_id": {
          "type": "string",
          "description": "Chef Infra Server ID."
        },
        "name": {
          "type": "string",
          "description": "Data bag name."
        },
        "item_id": {
          "type": "string",
          "description": "Data bag item ID."
        },
        "data": {
          "type": "object",
          "description": "Data bag item JSON data."
        }
      }
    },
    "chef.automate.api.infra_proxy.request.UpdateEnvironment": {
      "type": "object",
      "properties": {
        "org_id": {
          "type": "string",
          "description": "Chef organization ID."
        },
        "server_id": {
          "type": "string",
          "description": "Chef Infra Server ID."
        },
        "name": {
          "type": "string",
          "description": "Environment name."
        },
        "description": {
          "type": "string",
          "description": "Environment description."
        },
        "json_class": {
          "type": "string",
          "description": "Class name."
        },
        "cookbook_versions": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "description": "Environment versioned cookbooks constraints."
        },
        "default_attributes": {
          "type": "object",
          "description": "Environment default attributes JSON."
        },
        "override_attributes": {
          "type": "object",
          "description": "Environment override attributes JSON."
        }
      }
    },
    "chef.automate.api.infra_proxy.request.UpdateNodeAttributes": {
      "type": "object",
      "properties": {
        "org_id": {
          "type": "string",
          "description": "Chef organization ID."
        },
        "server_id": {
          "type": "string",
          "description": "Chef Infra Server ID."
        },
        "name": {
          "type": "string",
          "description": "Node name."
        },
        "attributes": {
          "type": "object",
          "description": "Node attributes JSON."
        }
      }
    },
    "chef.automate.api.infra_proxy.request.UpdateNodeEnvironment": {
      "type": "object",
      "properties": {
        "org_id": {
          "type": "string",
          "description": "Chef organization ID."
        },
        "server_id": {
          "type": "string",
          "description": "Chef Infra Server ID."
        },
        "name": {
          "type": "string",
          "description": "Node name."
        },
        "environment": {
          "type": "string",
          "description": "Node environment name."
        }
      }
    },
    "chef.automate.api.infra_proxy.request.UpdateNodeTags": {
      "type": "object",
      "properties": {
        "org_id": {
          "type": "string",
          "description": "Chef organization ID."
        },
        "server_id": {
          "type": "string",
          "description": "Chef Infra Server ID."
        },
        "name": {
          "type": "string",
          "description": "Node name."
        },
        "action": {
          "type": "string",
          "description": "Node tags action (e.g. 'add', 'delete', 'set')."
        },
        "tags": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Node tags."
        }
      }
    },
    "chef.automate.api.infra_proxy.request.UpdateOrg": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Chef organization ID."
        },
        "name": {
          "type": "string",
          "description": "Chef organization name."
        },
        "admin_user": {
          "type": "string",
          "description": "Chef organization admin user."
        },
        "server_id": {
          "type": "string",
          "description": "Chef Infra Server ID."
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of projects this chef organization belongs to. May be empty."
        }
      }
    },
    "chef.automate.api.infra_proxy.request.UpdateRole": {
      "type": "object",
      "properties": {
        "org_id": {
          "type": "string",
          "description": "Chef organization ID."
        },
        "server_id": {
          "type": "string",
          "description": "Chef Infra Server ID."
        },
        "name": {
          "type": "string",
          "description": "Role name."
        },
        "description": {
          "type": "string",
          "description": "Role description."
        },
        "default_attributes": {
          "type": "object",
          "description": "Role default attributes JSON."
        },
        "override_attributes": {
          "type": "object",
          "description": "Role override attributes JSON."
        },
        "run_list": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Role run list."
        },
        "env_run_lists": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.request.EnvRunList"
          },
          "description": "Environment based run list."
        }
      }
    },
    "chef.automate.api.infra_proxy.request.UpdateServer": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Chef Infra Server ID."
        },
        "name": {
          "type": "string",
          "description": "Chef Infra Server name."
        },
        "fqdn": {
          "type": "string",
          "description": "Chef Infra Server FQDN."
        },
        "ip_address": {
          "type": "string",
          "description": "Chef Infra Server IP address."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.Client": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Client name."
        },
        "client_name": {
          "type": "string",
          "description": "Client name return by Chef Infra Server API."
        },
        "org_name": {
          "type": "string",
          "description": "Chef organization name."
        },
        "validator": {
          "type": "boolean",
          "format": "boolean",
          "description": "Boolean indicates client type is validator or not."
        },
        "json_class": {
          "type": "string",
          "description": "Client JSON class."
        },
        "chef_type": {
          "type": "string",
          "description": "Chef object type."
        },
        "client_key": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.ClientAccessKey",
          "description": "Client key detail."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.ClientAccessKey": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Client key name."
        },
        "public_key": {
          "type": "string",
          "description": "Client public key."
        },
        "expiration_date": {
          "type": "string",
          "description": "Client key expiration date string."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.ClientKey": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Client key name."
        },
        "public_key": {
          "type": "string",
          "description": "Client public key."
        },
        "expiration_date": {
          "type": "string",
          "description": "Client key expiration date string."
        },
        "private_key": {
          "type": "string",
          "description": "Client private key."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.ClientListItem": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Client name."
        },
        "validator": {
          "type": "boolean",
          "format": "boolean",
          "description": "Boolean indicates client type is validator or not."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.Clients": {
      "type": "object",
      "properties": {
        "clients": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.ClientListItem"
          },
          "description": "Client list."
        },
        "page": {
          "type": "integer",
          "format": "int32",
          "description": "Starting page for the results."
        },
        "total": {
          "type": "integer",
          "format": "int32",
          "description": "Total number of records."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.Cookbook": {
      "type": "object",
      "properties": {
        "cookbook_name": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "version": {
          "type": "string"
        },
        "chef_type": {
          "type": "string"
        },
        "frozen": {
          "type": "boolean",
          "format": "boolean"
        },
        "json_class": {
          "type": "string"
        },
        "files": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookItem"
          }
        },
        "templates": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookItem"
          }
        },
        "attributes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookItem"
          }
        },
        "recipes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookItem"
          }
        },
        "definitions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookItem"
          }
        },
        "libraries": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookItem"
          }
        },
        "providers": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookItem"
          }
        },
        "resources": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookItem"
          }
        },
        "root_files": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookItem"
          }
        },
        "metadata": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookMeta"
        },
        "access": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookAccess"
        }
      }
    },
    "chef.automate.api.infra_proxy.response.CookbookAccess": {
      "type": "object",
      "properties": {
        "read": {
          "type": "boolean",
          "format": "boolean"
        },
        "create": {
          "type": "boolean",
          "format": "boolean"
        },
        "grant": {
          "type": "boolean",
          "format": "boolean"
        },
        "update": {
          "type": "boolean",
          "format": "boolean"
        },
        "delete": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "chef.automate.api.infra_proxy.response.CookbookFileContent": {
      "type": "object",
      "properties": {
        "content": {
          "type": "string",
          "description": "Cookbook data file content."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.CookbookItem": {
      "type": "object",
      "properties": {
        "url": {
          "type": "string"
        },
        "path": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "checksum": {
          "type": "string"
        },
        "specificity": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.infra_proxy.response.CookbookLock": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Cookbook name."
        },
        "version": {
          "type": "string",
          "description": "Cookbook version."
        },
        "identifier": {
          "type": "string",
          "description": "Cookbook identifier."
        },
        "dotted_identifier": {
          "type": "string",
          "description": "Cookbook decimal number identifier."
        },
        "source": {
          "type": "string",
          "description": "Cookbook source."
        },
        "cache_key": {
          "type": "string",
          "description": "Cookbook cache key."
        },
        "SCMDetail": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.SCMDetail",
          "description": "SCM detail."
        },
        "source_options": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.SourceOptions",
          "description": "Cookbook source path."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.CookbookMeta": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "version": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "long_description": {
          "type": "string"
        },
        "maintainer": {
          "type": "string"
        },
        "maintainer_email": {
          "type": "string"
        },
        "license": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.infra_proxy.response.CookbookVersion": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Name of the cookbook."
        },
        "version": {
          "type": "string",
          "description": "Version of the cookbook."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.CookbookVersions": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Name of the cookbook."
        },
        "versions": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of all versions available for cookbook."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.Cookbooks": {
      "type": "object",
      "properties": {
        "cookbooks": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookVersion"
          },
          "description": "List of cookbooks with name and version."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.CreateClient": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Client name."
        },
        "client_key": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.ClientKey",
          "description": "Client key detail."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.CreateDataBag": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Data bag name."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.CreateDataBagItem": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Data bag name."
        },
        "id": {
          "type": "string",
          "description": "Data bag item ID."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.CreateOrg": {
      "type": "object",
      "properties": {
        "org": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Org",
          "description": "Chef organization."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.CreateServer": {
      "type": "object",
      "properties": {
        "server": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Server",
          "description": "Chef Infra Server."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.DataBag": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Data bag name."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.DataBagItem": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Data bag name."
        },
        "id": {
          "type": "string",
          "description": "Data bag item ID."
        },
        "data": {
          "type": "string",
          "description": "Stringified json of data bag item."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.DataBagItems": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "title": "Data bag name"
        },
        "items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.DataBagListItem"
          },
          "description": "Data bags item list."
        },
        "page": {
          "type": "integer",
          "format": "int32",
          "description": "Starting page for the results."
        },
        "total": {
          "type": "integer",
          "format": "int32",
          "description": "Total number of records."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.DataBagListItem": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Data bag item name."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.DataBags": {
      "type": "object",
      "properties": {
        "data_bags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.DataBagListItem"
          },
          "description": "Data bags item list."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.DeleteNode": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Node name."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.DeleteOrg": {
      "type": "object",
      "properties": {
        "org": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Org",
          "description": "Chef organization."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.DeletePolicyfile": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Policyfile name."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.DeleteServer": {
      "type": "object",
      "properties": {
        "server": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Server",
          "description": "Chef Infra Server."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.Environment": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Environment name."
        },
        "chef_type": {
          "type": "string",
          "description": "Chef object type."
        },
        "description": {
          "type": "string",
          "description": "Environment description."
        },
        "json_class": {
          "type": "string",
          "description": "Environment JSON class."
        },
        "cookbook_versions": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "description": "Environment versined cookbooks constraints."
        },
        "default_attributes": {
          "type": "string",
          "description": "Environment default attributes JSON."
        },
        "override_attributes": {
          "type": "string",
          "description": "Environment override attributes JSON."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.EnvironmentListItem": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Environment name."
        },
        "description": {
          "type": "string",
          "description": "Environment description."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.EnvironmentRecipesList": {
      "type": "object",
      "properties": {
        "recipes": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Environment recipes list."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.Environments": {
      "type": "object",
      "properties": {
        "environments": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.EnvironmentListItem"
          },
          "description": "Environments list."
        },
        "page": {
          "type": "integer",
          "format": "int32",
          "description": "Starting page for the results."
        },
        "total": {
          "type": "integer",
          "format": "int32",
          "description": "Total number of records."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.ExpandedRunList": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "ID of the run list collection."
        },
        "run_list": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.RunList"
          },
          "description": "List of the run list."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.GetOrg": {
      "type": "object",
      "properties": {
        "org": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Org",
          "description": "Chef organization."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.GetOrgs": {
      "type": "object",
      "properties": {
        "orgs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Org"
          },
          "description": "Chef organization list."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.GetServer": {
      "type": "object",
      "properties": {
        "server": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Server",
          "description": "Chef Infra Server."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.GetServers": {
      "type": "object",
      "properties": {
        "servers": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Server"
          },
          "description": "List of Chef Infra Servers."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.IncludedPolicyLock": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Included Policyfile name."
        },
        "revision_id": {
          "type": "string",
          "description": "Policyfile revision ID."
        },
        "source_options": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.SourceOptions",
          "description": "Included policyfile source options."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.NamedRunList": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Run list name."
        },
        "run_list": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Run list associated with the policy."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.Node": {
      "type": "object",
      "properties": {
        "node_id": {
          "type": "string",
          "description": "Node ID."
        },
        "name": {
          "type": "string",
          "description": "Node name."
        },
        "environment": {
          "type": "string",
          "description": "Node environment."
        },
        "policy_name": {
          "type": "string",
          "description": "Node policy name."
        },
        "policy_group": {
          "type": "string",
          "description": "Node policy group."
        },
        "run_list": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Node run-list."
        },
        "tags": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Node tags."
        },
        "automatic_attributes": {
          "type": "string",
          "description": "Node automatic attributes JSON."
        },
        "default_attributes": {
          "type": "string",
          "description": "Node default attributes JSON."
        },
        "normal_attributes": {
          "type": "string",
          "description": "Node normal attributes JSON."
        },
        "override_attributes": {
          "type": "string",
          "description": "Node override attributes JSON."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.NodeAttribute": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Node ID."
        },
        "name": {
          "type": "string",
          "description": "Node name."
        },
        "check_in": {
          "type": "string",
          "description": "Node last checkin."
        },
        "uptime": {
          "type": "string",
          "description": "Node uptime."
        },
        "platform": {
          "type": "string",
          "description": "Node platform."
        },
        "environment": {
          "type": "string",
          "description": "Node environment name."
        },
        "policy_group": {
          "type": "string",
          "description": "Node policy group."
        },
        "fqdn": {
          "type": "string",
          "description": "Node FQDN."
        },
        "ip_address": {
          "type": "string",
          "description": "Node IP address."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.NodeExpandedRunList": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "ID of the run list collection."
        },
        "run_list": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.RunList"
          },
          "description": "List of the run list."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.Nodes": {
      "type": "object",
      "properties": {
        "nodes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.NodeAttribute"
          },
          "description": "Node list."
        },
        "page": {
          "type": "integer",
          "format": "int32",
          "description": "Starting page for the results."
        },
        "total": {
          "type": "integer",
          "format": "int32",
          "description": "Total number of records."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.Org": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Chef organization ID."
        },
        "name": {
          "type": "string",
          "description": "Chef organization name."
        },
        "admin_user": {
          "type": "string",
          "description": "Chef organization admin user."
        },
        "credential_id": {
          "type": "string",
          "description": "Chef organization credential ID."
        },
        "server_id": {
          "type": "string",
          "description": "Chef Infra Server ID."
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of projects this chef organization belongs to. May be empty."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.Policyfile": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Policyfile name."
        },
        "policy_group": {
          "type": "string",
          "description": "Policy group name."
        },
        "revision_id": {
          "type": "string",
          "description": "Policy revision ID."
        },
        "run_list": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Run-list associated with the policy."
        },
        "named_run_list": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.NamedRunList"
          },
          "description": "Named run-list associated with the policy."
        },
        "included_policy_locks": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.IncludedPolicyLock"
          },
          "description": "Included policy locks files."
        },
        "cookbook_locks": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.CookbookLock"
          },
          "description": "List of cookbook locks under this policy."
        },
        "default_attributes": {
          "type": "string",
          "description": "Policyfile default attributes JSON."
        },
        "override_attributes": {
          "type": "string",
          "description": "Policyfile override attributes JSON."
        },
        "expanded_run_list": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.ExpandedRunList"
          },
          "description": "Expanded run-list associated with the policy."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.PolicyfileListItem": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Policyfile name."
        },
        "revision_id": {
          "type": "string",
          "description": "Policyfile Revision ID."
        },
        "policy_group": {
          "type": "string",
          "description": "Policyfile policy group."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.Policyfiles": {
      "type": "object",
      "properties": {
        "policies": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.PolicyfileListItem"
          },
          "description": "Policyfiles list."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.ResetClient": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Client name."
        },
        "client_key": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.ClientKey",
          "description": "Client key detail."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.ResetOrgAdminKey": {
      "type": "object",
      "properties": {
        "org": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Org",
          "description": "Chef organization."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.Role": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Name of the role."
        },
        "chef_type": {
          "type": "string",
          "description": "Type of the chef object."
        },
        "description": {
          "type": "string",
          "description": "Descrption of the role."
        },
        "default_attributes": {
          "type": "string",
          "description": "Role default attributes JSON."
        },
        "override_attributes": {
          "type": "string",
          "description": "Role override attributes JSON."
        },
        "json_class": {
          "type": "string",
          "description": "Json class name."
        },
        "run_list": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Run list for the role."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.RoleEnvironments": {
      "type": "object",
      "properties": {
        "environments": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Role environment list."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.RoleListItem": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Name of the role."
        },
        "description": {
          "type": "string",
          "description": "Desscription of the role."
        },
        "environments": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Environment for the role."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.Roles": {
      "type": "object",
      "properties": {
        "roles": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.RoleListItem"
          },
          "description": "List of the roles item."
        },
        "page": {
          "type": "integer",
          "format": "int32",
          "description": "Starting page for the results."
        },
        "total": {
          "type": "integer",
          "format": "int32",
          "description": "Total number of records."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.RunList": {
      "type": "object",
      "properties": {
        "type": {
          "type": "string",
          "description": "Run list item type (e.g. 'recipe')."
        },
        "name": {
          "type": "string",
          "description": "Run list item name."
        },
        "version": {
          "type": "string",
          "description": "Run list item version."
        },
        "skipped": {
          "type": "boolean",
          "format": "boolean",
          "description": "Boolean denoting whether or not the run list item was skipped."
        },
        "position": {
          "type": "integer",
          "format": "int32",
          "title": "Run list item position"
        },
        "error": {
          "type": "string",
          "title": "Run list error"
        },
        "children": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.RunList"
          },
          "description": "List of the run list."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.SCMDetail": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "SCM name."
        },
        "remote": {
          "type": "string",
          "description": "SCM remote location."
        },
        "revision": {
          "type": "string",
          "description": "SCM revision detail."
        },
        "working_tree_clean": {
          "type": "boolean",
          "format": "boolean",
          "description": "Boolean that denotes if the working tree is clean or not."
        },
        "published": {
          "type": "boolean",
          "format": "boolean",
          "description": "Source's published information."
        },
        "synchronized_remote_branches": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Synchronized remote branches list."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.Server": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Chef Infra Server ID."
        },
        "name": {
          "type": "string",
          "description": "Chef Infra Server name."
        },
        "fqdn": {
          "type": "string",
          "description": "Chef Infra Server FQDN."
        },
        "ip_address": {
          "type": "string",
          "description": "Chef Infra Server IP address."
        },
        "orgs_count": {
          "type": "integer",
          "format": "int32",
          "description": "Chef organizations count associated with Chef Infra Server."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.SourceOptions": {
      "type": "object",
      "properties": {
        "path": {
          "type": "string",
          "description": "Source options path."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.UpdateDataBagItem": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Data bag name."
        },
        "item_id": {
          "type": "string",
          "description": "Data bag item ID."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.UpdateNodeAttributes": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Node name."
        },
        "attributes": {
          "type": "string",
          "description": "Node attributes JSON."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.UpdateNodeEnvironment": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Node name."
        },
        "environment": {
          "type": "string",
          "description": "Node environment name."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.UpdateNodeTags": {
      "type": "object",
      "properties": {
        "tags": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Node tags."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.UpdateOrg": {
      "type": "object",
      "properties": {
        "org": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Org",
          "description": "Chef organization."
        }
      }
    },
    "chef.automate.api.infra_proxy.response.UpdateServer": {
      "type": "object",
      "properties": {
        "server": {
          "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Server",
          "description": "Chef Infra Server."
        }
      }
    },
    "google.protobuf.Any": {
      "type": "object",
      "properties": {
        "type_url": {
          "type": "string"
        },
        "value": {
          "type": "string",
          "format": "byte"
        }
      }
    },
    "google.protobuf.NullValue": {
      "type": "string",
      "enum": [
        "NULL_VALUE"
      ],
      "default": "NULL_VALUE",
      "description": "` + "`" + `NullValue` + "`" + ` is a singleton enumeration to represent the null value for the\n` + "`" + `Value` + "`" + ` type union.\n\n The JSON representation for ` + "`" + `NullValue` + "`" + ` is JSON ` + "`" + `null` + "`" + `.\n\n - NULL_VALUE: Null value."
    },
    "grpc.gateway.runtime.Error": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        },
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "details": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/google.protobuf.Any"
          }
        }
      }
    }
  }
}
`)
}
