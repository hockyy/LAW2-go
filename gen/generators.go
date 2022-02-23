package gen

//go:generate oapi-codegen -generate server -o ./openapi/server.gen.go -package gen openapi.yaml
//go:generate oapi-codegen -generate types -o ./openapi/types.gen.go -package gen openapi.yaml
//go:generate oapi-codegen -generate client -o ./openapi/client.gen.go -package gen openapi.yaml
//go:generate oapi-codegen -generate spec -o ./openapi/spec.gen.go -package gen openapi.yaml
