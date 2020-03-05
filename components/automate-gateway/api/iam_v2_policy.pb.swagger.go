package api

func init() {
	Swagger.Add("iam_v2_policy", `{
  "swagger": "2.0",
  "info": {
    "title": "components/automate-gateway/api/iam/v2/policy.proto",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/iam/v2/introspect_projects": {
      "get": {
        "operationId": "IntrospectAllProjects",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.ListProjectsResp"
            }
          }
        },
        "tags": [
          "hidden"
        ]
      }
    },
    "/iam/v2/policies": {
      "get": {
        "summary": "List all policies",
        "description": "List all policies.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\niam:policies:list\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "ListPolicies",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.ListPoliciesResp"
            }
          }
        },
        "tags": [
          "policies"
        ]
      },
      "post": {
        "summary": "Create a custom policy",
        "description": "Creates a custom IAM policy used to control permissions in Automate.\nA policy is composed of one or more statements that grant permissions to a set of members.\nEach statement contains a role as well as a list of projects.\n\nThe role defines a set of actions that the statement is scoped to.\nThe project list defines the set of resources that the statement is scoped to.\nPass ` + "`" + `\"projects\": [\"*\"]` + "`" + ` to scope a statement to every project.\n\nA policy's *top-level* projects list defines which projects the policy belongs to (for filtering policies by their projects),\nwhereas the *statement-level* projects list defines which projects the statement applies to.\n\nThis example creates a new policy not associated with any project (because the top-level ` + "`" + `projects` + "`" + ` property is empty) that grants the ` + "`" + `viewer` + "`" + ` role\non a few projects for all LDAP teams and a custom role ` + "`" + `qa` + "`" + ` on a specific project.\n\nExample:\n` + "`" + `` + "`" + `` + "`" + `\n{\n\"name\": \"My Viewer Policy\",\n\"id\": \"viewer-policy\",\n\"members\": [\"team:ldap:*\"],\n\"statements\" : [\n{\n\"role\": \"viewer\",\n\"projects\": [\"project1\", \"project2\"]\n},\n{\n\"role\": \"qa\",\n\"projects\": [\"acceptanceProject\"]\n}\n],\n\"projects\": []\n}\n` + "`" + `` + "`" + `` + "`" + `\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\niam:policies:create\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "CreatePolicy",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.CreatePolicyResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.CreatePolicyReq"
            }
          }
        ],
        "tags": [
          "policies"
        ]
      }
    },
    "/iam/v2/policies/{id}": {
      "get": {
        "summary": "Get a policy",
        "description": "Returns the details for a policy.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\niam:policies:get\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "GetPolicy",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.GetPolicyResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "ID of the policy.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "policies"
        ]
      },
      "delete": {
        "summary": "Delete a custom policy",
        "description": "Delete a specified custom policy. You cannot delete Chef-managed policies.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\niam:policies:delete\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "DeletePolicy",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.DeletePolicyResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "ID of the policy.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "policies"
        ]
      },
      "put": {
        "summary": "Update a custom policy",
        "description": "This operation overwrites all fields excepting ID,\nincluding those omitted from the request, so be sure to specify all properties.\nProperties that you do not include are reset to empty values.\nThe only exception is the policy ID, which is immutable; it can only be set at creation time.\n\nWhile you can use this endpoint to update members on a policy, if that is the only\nproperty you wish to modify you might find it more convenient to use one of these endpoints instead:\nAdd policy members, Remove policy members, or Replace policy members.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\niam:policies:update\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "UpdatePolicy",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.UpdatePolicyResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "Unique ID. Cannot be changed.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.UpdatePolicyReq"
            }
          }
        ],
        "tags": [
          "policies"
        ]
      }
    },
    "/iam/v2/policies/{id}/members": {
      "get": {
        "summary": "List policy members",
        "description": "List all members of a specific policy.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\niam:policyMembers:get\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "ListPolicyMembers",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.ListPolicyMembersResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "ID of the policy.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "policies"
        ]
      },
      "put": {
        "summary": "Replace policy members",
        "description": "Replace the entire member list of a specific policy with a new list.\nYou may use this endpoint to update members of either Custom or Chef-managed policies.\n\nEnsure each element of the members array is in the correct\n[Member Expression](https://automate.chef.io/docs/iam-v2-guide/#member-expressions) format.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\niam:policyMembers:update\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "ReplacePolicyMembers",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.ReplacePolicyMembersResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "ID of the policy.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.ReplacePolicyMembersReq"
            }
          }
        ],
        "tags": [
          "policies"
        ]
      }
    },
    "/iam/v2/policies/{id}/members:add": {
      "post": {
        "summary": "Add policy members",
        "description": "Add members to the member list of a specific policy.\nYou may use this endpoint to update members of either Custom or Chef-managed policies.\n\nEnsure each element of the members array is in the correct\n[Member Expression](https://automate.chef.io/docs/iam-v2-guide/#member-expressions) format.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\niam:policyMembers:create\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "AddPolicyMembers",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.AddPolicyMembersResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "ID of the policy.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.AddPolicyMembersReq"
            }
          }
        ],
        "tags": [
          "policies"
        ]
      }
    },
    "/iam/v2/policies/{id}/members:remove": {
      "post": {
        "summary": "Remove policy members",
        "description": "Remove members from the member list of a specific policy. Silently ignores\nmembers that are not already part of the member list.\nYou may use this endpoint to update members of either Custom or Chef-managed policies.\n\nEnsure each element of the members array is in the correct\n[Member Expression](https://automate.chef.io/docs/iam-v2-guide/#member-expressions) format.\n\nThe removed members will still exist within Chef Automate, but are no longer associated with this policy.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\niam:policyMembers:delete\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "RemovePolicyMembers",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.RemovePolicyMembersResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "ID of the policy.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.RemovePolicyMembersReq"
            }
          }
        ],
        "tags": [
          "policies"
        ]
      }
    },
    "/iam/v2/policy_version": {
      "get": {
        "summary": "Get IAM version",
        "description": "Returns the major and minor version of IAM that your automate installation is running.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\niam:policies:get\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "GetPolicyVersion",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.GetPolicyVersionResp"
            }
          }
        },
        "tags": [
          "policies"
        ]
      }
    },
    "/iam/v2/projects": {
      "get": {
        "summary": "List all projects",
        "description": "List all projects.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\niam:projects:list\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "ListProjects",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.ListProjectsResp"
            }
          }
        },
        "tags": [
          "projects"
        ]
      },
      "post": {
        "summary": "Create a project",
        "description": "Creates a new project to be used in the policies that control permissions in Automate.\n\nA project defines the scope of resources in a policy statement. Resources can be in more than one project.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\niam:projects:create\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "CreateProject",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.CreateProjectResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.CreateProjectReq"
            }
          }
        ],
        "tags": [
          "projects"
        ]
      }
    },
    "/iam/v2/projects/{id}": {
      "get": {
        "summary": "Get a project",
        "description": "Returns the details for a project.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\niam:projects:get\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "GetProject",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.GetProjectResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "ID of the project.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "projects"
        ]
      },
      "delete": {
        "summary": "Delete a project",
        "description": "Delete a project from any resources tagged with it.\n\nAlso deletes this project from any project list in any policy statements.\nIf the resulting project list for a given statement is empty, it is deleted.\nIf the resulting policy has no statements, it is also deleted.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\niam:projects:delete\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "DeleteProject",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.DeleteProjectResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "ID of the project.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "projects"
        ]
      },
      "put": {
        "summary": "Update a project",
        "description": "Updates the name of an existing project.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\niam:projects:update\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "UpdateProject",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.UpdateProjectResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "Unique ID. Cannot be changed.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.UpdateProjectReq"
            }
          }
        ],
        "tags": [
          "projects"
        ]
      }
    },
    "/iam/v2/roles": {
      "get": {
        "summary": "List all roles",
        "description": "List all *Chef-managed* and *Custom* roles.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\niam:roles:list\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "ListRoles",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.ListRolesResp"
            }
          }
        },
        "tags": [
          "roles"
        ]
      },
      "post": {
        "summary": "Create a custom role",
        "description": "Creates a new role to be used in the policies that control permissions in Automate.\n\nA role defines the scope of actions in a policy statement.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\niam:roles:create\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "CreateRole",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.CreateRoleResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.CreateRoleReq"
            }
          }
        ],
        "tags": [
          "roles"
        ]
      }
    },
    "/iam/v2/roles/{id}": {
      "get": {
        "summary": "Get a role",
        "description": "Returns the details for a role.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\niam:roles:get\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "GetRole",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.GetRoleResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "ID of the role.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "roles"
        ]
      },
      "delete": {
        "summary": "Delete a custom role",
        "description": "Delete a specified custom role (you cannot delete Chef-managed roles) and remove it from any statements that may have been using it.\nIf such a statement has no other associated actions, the statement is deleted as well.\nSimilarly, if that statement removal results in a policy with no other statements,\nthat policy is removed as well.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\niam:roles:delete\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "DeleteRole",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.DeleteRoleResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "ID of the role.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "roles"
        ]
      },
      "put": {
        "summary": "Update a custom role",
        "description": "This operation overwrites all fields excepting ID,\nincluding those omitted from the request, so be sure to specify all properties.\nProperties that you do not include are reset to empty values.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\niam:roles:update\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "UpdateRole",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.UpdateRoleResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "Unique ID. Cannot be changed.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.UpdateRoleReq"
            }
          }
        ],
        "tags": [
          "roles"
        ]
      }
    }
  },
  "definitions": {
    "chef.automate.api.iam.v2.AddPolicyMembersReq": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "ID of the policy."
        },
        "members": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of members to add to the policy."
        }
      },
      "required": [
        "id",
        "members"
      ]
    },
    "chef.automate.api.iam.v2.AddPolicyMembersResp": {
      "type": "object",
      "properties": {
        "members": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Resulting list of policy members."
        }
      }
    },
    "chef.automate.api.iam.v2.CreatePolicyReq": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Unique ID. Cannot be changed."
        },
        "name": {
          "type": "string",
          "description": "Name for the policy."
        },
        "members": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Members affected by this policy."
        },
        "statements": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.iam.v2.Statement"
          },
          "description": "Statements for the policy."
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of projects this policy belongs to."
        }
      },
      "description": "Does not contain type as the enduser can only create 'custom' policies.",
      "required": [
        "id",
        "name",
        "statements"
      ]
    },
    "chef.automate.api.iam.v2.CreatePolicyResp": {
      "type": "object",
      "properties": {
        "policy": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Policy"
        }
      }
    },
    "chef.automate.api.iam.v2.CreateProjectReq": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Unique ID. Cannot be changed."
        },
        "name": {
          "type": "string",
          "description": "Name for the new project."
        }
      },
      "required": [
        "id",
        "name"
      ]
    },
    "chef.automate.api.iam.v2.CreateProjectResp": {
      "type": "object",
      "properties": {
        "project": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Project"
        }
      }
    },
    "chef.automate.api.iam.v2.CreateRoleReq": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Unique ID. Cannot be changed."
        },
        "name": {
          "type": "string",
          "description": "Name for the role."
        },
        "actions": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of actions that this role scopes to."
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of projects this role belongs to."
        }
      },
      "description": "Does not contain type as the enduser can only create 'custom' roles.",
      "required": [
        "id",
        "name",
        "actions"
      ]
    },
    "chef.automate.api.iam.v2.CreateRoleResp": {
      "type": "object",
      "properties": {
        "role": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Role"
        }
      }
    },
    "chef.automate.api.iam.v2.DeletePolicyResp": {
      "type": "object"
    },
    "chef.automate.api.iam.v2.DeleteProjectResp": {
      "type": "object"
    },
    "chef.automate.api.iam.v2.DeleteRoleResp": {
      "type": "object"
    },
    "chef.automate.api.iam.v2.GetPolicyResp": {
      "type": "object",
      "properties": {
        "policy": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Policy"
        }
      }
    },
    "chef.automate.api.iam.v2.GetPolicyVersionResp": {
      "type": "object",
      "properties": {
        "version": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Version"
        }
      }
    },
    "chef.automate.api.iam.v2.GetProjectResp": {
      "type": "object",
      "properties": {
        "project": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Project"
        }
      }
    },
    "chef.automate.api.iam.v2.GetRoleResp": {
      "type": "object",
      "properties": {
        "role": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Role"
        }
      }
    },
    "chef.automate.api.iam.v2.ListPoliciesResp": {
      "type": "object",
      "properties": {
        "policies": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.iam.v2.Policy"
          }
        }
      }
    },
    "chef.automate.api.iam.v2.ListPolicyMembersResp": {
      "type": "object",
      "properties": {
        "members": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of policy members."
        }
      }
    },
    "chef.automate.api.iam.v2.ListProjectsResp": {
      "type": "object",
      "properties": {
        "projects": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.iam.v2.Project"
          }
        }
      }
    },
    "chef.automate.api.iam.v2.ListRolesResp": {
      "type": "object",
      "properties": {
        "roles": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.iam.v2.Role"
          }
        }
      }
    },
    "chef.automate.api.iam.v2.Policy": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Name for the policy."
        },
        "id": {
          "type": "string",
          "description": "Unique ID. Cannot be changed."
        },
        "type": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Type",
          "description": "This doc-comment is ignored for an enum."
        },
        "members": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Members affected by this policy. May be empty."
        },
        "statements": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.iam.v2.Statement"
          },
          "description": "Statements for the policy. Will contain one or more."
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of projects this policy belongs to. May be empty."
        }
      }
    },
    "chef.automate.api.iam.v2.Project": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Name for the project."
        },
        "id": {
          "type": "string",
          "description": "Unique ID. Cannot be changed."
        },
        "type": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Type",
          "description": "Whether this policy is user created (` + "`" + `CUSTOM` + "`" + `) or chef managed (` + "`" + `CHEF_MANAGED` + "`" + `)."
        },
        "status": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.ProjectRulesStatus",
          "description": "The current status of the rules for this project."
        }
      }
    },
    "chef.automate.api.iam.v2.ProjectRulesStatus": {
      "type": "string",
      "enum": [
        "PROJECT_RULES_STATUS_UNSET",
        "RULES_APPLIED",
        "EDITS_PENDING",
        "NO_RULES"
      ],
      "default": "PROJECT_RULES_STATUS_UNSET"
    },
    "chef.automate.api.iam.v2.RemovePolicyMembersReq": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "ID of the policy."
        },
        "members": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of members to remove from the policy."
        }
      },
      "required": [
        "id",
        "members"
      ]
    },
    "chef.automate.api.iam.v2.RemovePolicyMembersResp": {
      "type": "object",
      "properties": {
        "members": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Resulting list of policy members."
        }
      }
    },
    "chef.automate.api.iam.v2.ReplacePolicyMembersReq": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "ID of the policy."
        },
        "members": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of members that replaces previous policy member list."
        }
      },
      "required": [
        "id"
      ]
    },
    "chef.automate.api.iam.v2.ReplacePolicyMembersResp": {
      "type": "object",
      "properties": {
        "members": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Resulting list of policy members."
        }
      }
    },
    "chef.automate.api.iam.v2.Role": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "Name for the role."
        },
        "id": {
          "type": "string",
          "description": "Unique ID. Cannot be changed."
        },
        "type": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Type",
          "description": "Whether this policy is user created (` + "`" + `CUSTOM` + "`" + `) or chef managed (` + "`" + `CHEF_MANAGED` + "`" + `)."
        },
        "actions": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of actions this role scopes to. Will contain one or more."
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of projects this role belongs to. May be empty."
        }
      }
    },
    "chef.automate.api.iam.v2.Statement": {
      "type": "object",
      "properties": {
        "effect": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Statement.Effect",
          "description": "This doc-comment is ignored for an enum."
        },
        "actions": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Actions defined inline. May be empty.\nBest practices recommend that you use custom roles rather than inline actions where practical."
        },
        "role": {
          "type": "string",
          "description": "The role defines a set of actions that the statement is scoped to."
        },
        "resources": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "DEPRECATED: Resources defined inline. Use projects instead."
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "The project list defines the set of resources that the statement is scoped to. May be empty."
        }
      }
    },
    "chef.automate.api.iam.v2.Statement.Effect": {
      "type": "string",
      "enum": [
        "ALLOW",
        "DENY"
      ],
      "default": "ALLOW"
    },
    "chef.automate.api.iam.v2.Type": {
      "type": "string",
      "enum": [
        "CHEF_MANAGED",
        "CUSTOM"
      ],
      "default": "CHEF_MANAGED"
    },
    "chef.automate.api.iam.v2.UpdatePolicyReq": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Unique ID. Cannot be changed."
        },
        "members": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Members affected by this policy."
        },
        "statements": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.iam.v2.Statement"
          },
          "description": "Statements for the policy."
        },
        "name": {
          "type": "string",
          "description": "Name for this policy."
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of projects this policy belongs to."
        }
      },
      "description": "Does not contain type as the enduser can only create 'custom' policies.",
      "required": [
        "id",
        "name",
        "statements"
      ]
    },
    "chef.automate.api.iam.v2.UpdatePolicyResp": {
      "type": "object",
      "properties": {
        "policy": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Policy"
        }
      }
    },
    "chef.automate.api.iam.v2.UpdateProjectReq": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Unique ID. Cannot be changed."
        },
        "name": {
          "type": "string",
          "description": "Name for the project."
        }
      },
      "required": [
        "id",
        "name"
      ]
    },
    "chef.automate.api.iam.v2.UpdateProjectResp": {
      "type": "object",
      "properties": {
        "project": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Project"
        }
      }
    },
    "chef.automate.api.iam.v2.UpdateRoleReq": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Unique ID. Cannot be changed."
        },
        "name": {
          "type": "string",
          "description": "Name for the role."
        },
        "actions": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of actions that this role scopes to."
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "List of projects this role belongs to."
        }
      },
      "required": [
        "id",
        "name",
        "actions"
      ]
    },
    "chef.automate.api.iam.v2.UpdateRoleResp": {
      "type": "object",
      "properties": {
        "role": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Role"
        }
      }
    },
    "chef.automate.api.iam.v2.Version": {
      "type": "object",
      "properties": {
        "major": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Version.VersionNumber"
        },
        "minor": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Version.VersionNumber"
        }
      },
      "title": "the only values that may be returned by GetPolicyVersion"
    },
    "chef.automate.api.iam.v2.Version.VersionNumber": {
      "type": "string",
      "enum": [
        "V0",
        "V1",
        "V2"
      ],
      "default": "V0"
    }
  }
}
`)
}
