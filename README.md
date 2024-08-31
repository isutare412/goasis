Simple go client/server code generation example using [OAS(OpenAPI
Specification)](https://spec.openapis.org/oas/latest.html).

Generated Codes
- [Client Code](/pkg/oapi/client.go)
- [Server Code](/pkg/oapi/server.go)

# Development

## Generate code from OAS

```bash
# Generate client code
make gen-oapi-client

# Generate server code
make gen-oapi-server
```

## Tests

```bash
make test
```
