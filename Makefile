gen-oapi:
    oapi-codegen -generate server -o gen/openapi_server.gen.go -package main openapi.yaml