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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/cluster": {
            "get": {
                "description": "list cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "k8s cluster"
                ],
                "summary": "list cluster",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page number, optional",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page size, optional",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "install cluster master",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "k8s cluster"
                ],
                "summary": "install cluster",
                "parameters": [
                    {
                        "description": "install cluster",
                        "name": "cluster",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Cluster"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/v1/cluster/nodes": {
            "post": {
                "description": "install k8s nodes",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "k8s cluster"
                ],
                "summary": "install cluster nodes",
                "parameters": [
                    {
                        "description": "install cluster slave",
                        "name": "cluster",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/kubernetes.JoinNodes"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/v1/cluster/upgrade": {
            "post": {
                "description": "install cluster slave",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "k8s cluster"
                ],
                "summary": "upgrade k8s",
                "parameters": [
                    {
                        "description": "k8s all nodes",
                        "name": "upgrade",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/kubernetes.Upgrade"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/v1/images/info": {
            "get": {
                "description": "image pull info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "image"
                ],
                "summary": "image",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/v1/images/pull": {
            "post": {
                "description": "image push to oss",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "image"
                ],
                "summary": "image push to oss",
                "parameters": [
                    {
                        "description": "pull Image, then push to OSS",
                        "name": "cluster",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/images.Pull"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/v1/proxy": {
            "post": {
                "description": "install proxy",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "proxy"
                ],
                "summary": "install proxy",
                "parameters": [
                    {
                        "description": "install proxy",
                        "name": "cluster",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/proxy.Install"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "object"
                }
            }
        },
        "images.Image": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "default": "nginx"
                },
                "version": {
                    "type": "string",
                    "default": "latest"
                }
            }
        },
        "images.Local": {
            "type": "object",
            "properties": {
                "path": {
                    "type": "string",
                    "default": "/data/images"
                }
            }
        },
        "images.Pull": {
            "type": "object",
            "properties": {
                "local": {
                    "type": "object",
                    "$ref": "#/definitions/images.Local"
                },
                "registry": {
                    "type": "object",
                    "$ref": "#/definitions/images.Registry"
                },
                "remoteHost": {
                    "type": "object",
                    "$ref": "#/definitions/images.RemoteHost"
                },
                "source": {
                    "type": "object",
                    "$ref": "#/definitions/images.Repo"
                },
                "type": {
                    "type": "string",
                    "default": "local"
                }
            }
        },
        "images.Registry": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                }
            }
        },
        "images.RemoteHost": {
            "type": "object",
            "properties": {
                "ip": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "port": {
                    "type": "integer",
                    "default": 22
                },
                "user": {
                    "type": "string",
                    "default": "root"
                }
            }
        },
        "images.Repo": {
            "type": "object",
            "properties": {
                "addr": {
                    "type": "string",
                    "default": "docker.io"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/images.Image"
                    }
                },
                "org": {
                    "type": "string",
                    "default": "library"
                },
                "password": {
                    "type": "string"
                },
                "user": {
                    "type": "string"
                }
            }
        },
        "kubernetes.JoinNodes": {
            "type": "object",
            "properties": {
                "controllerNodes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Host"
                    }
                },
                "master": {
                    "type": "object",
                    "$ref": "#/definitions/models.Host"
                },
                "workNodes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Host"
                    }
                }
            }
        },
        "kubernetes.Upgrade": {
            "type": "object",
            "properties": {
                "dryRun": {
                    "type": "boolean"
                },
                "nodes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Host"
                    }
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "models.Cluster": {
            "type": "object",
            "required": [
                "controlPlaneEndpoint",
                "master",
                "name"
            ],
            "properties": {
                "controlPlaneEndpoint": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "master": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Host"
                    }
                },
                "name": {
                    "type": "string",
                    "default": "test"
                },
                "networkPlug": {
                    "type": "string",
                    "default": "cilium"
                },
                "podCidr": {
                    "type": "string",
                    "default": "10.244.0.0/16"
                },
                "registry": {
                    "type": "string",
                    "default": "registry.aliyuncs.com/google_containers"
                },
                "serviceCidr": {
                    "type": "string",
                    "default": "10.96.0.0/16"
                },
                "version": {
                    "type": "string",
                    "default": "1.22.15"
                },
                "workNodes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Host"
                    }
                }
            }
        },
        "models.Host": {
            "type": "object",
            "properties": {
                "ip": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "port": {
                    "type": "integer",
                    "default": 22
                },
                "user": {
                    "type": "string",
                    "default": "root"
                }
            }
        },
        "proxy.Install": {
            "type": "object",
            "properties": {
                "commonName": {
                    "type": "string",
                    "default": "test.hello.com"
                },
                "externalPort": {
                    "description": "used for nginx tls， listen to proxy",
                    "type": "integer",
                    "default": 20001
                },
                "ip": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "port": {
                    "type": "integer",
                    "default": 22
                },
                "proxyPort": {
                    "type": "integer",
                    "default": 20000
                },
                "user": {
                    "type": "string",
                    "default": "root"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
        }
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
	Version:     "2.0",
	Host:        "",
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "Swagger Example API",
	Description: "This is a sample server celler server.",
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
