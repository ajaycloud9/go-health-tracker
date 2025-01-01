
# Command to generate the Go server/client code for API
```bash
cd go-health-tracker/src/go-health-app/infra

# Server code
swagger-codegen generate -i openapi.yaml -l go-server -o rest/generated/server

# API documentation
swagger-codegen generate -i openapi.yaml -l html2 -o rest/generated/html

# Client code
swagger-codegen generate -i openapi.yaml -l go -o /Users/ajaysingh/ws/go-health-tracker/common/gohealthtrackerclient/swagger

```