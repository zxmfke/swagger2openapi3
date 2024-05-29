# swagger2openapi3

🌍 *[English](README.md) ∙ [Chinese](README_zh-CN.md)*

Swagger2openapi3 provides a package to convert Swagger 2.0 specification JSON and YAML to OpenAPI 3.0.1. It also provides a tool called `swag2op`, which integrates [Swagger 2.0](https://github.com/swaggo/swag) generation and supports conversion to OpenAPI 3.0.3.

## Getting Started

1. Add comments to your API source code, See [Declarative Comments Format](https://github.com/swaggo/swag#declarative-comments-format).
2. Install swag by using:

```
go install github.com/zxmfke/swagger2openapi3/cmd/swag2op@latest
```

To build from source you need [Go](https://golang.org/dl/) (1.19 or newer).

3. Run `swag2op init` in the project's root folder which contains the `main.go` file. This will parse your comments and generate the required files (`docs` folder and `docs/docs.go`).

   ```bash
   swag2op init
   ```

Make sure to import the generated `docs/docs.go` so that your specific configuration gets `init`'ed. If your General API annotations do not live in `main.go`, you can let swag know with `-g` flag.

```go
import _ "example-module-name/docs"
```

```bash
swag2op init -g http/api.go
```

4. extra flags provided in next section

## swag2op cli

`swag2op` provides three additional flags for more flexibility:

1. `disableConvertOpenApiV3Flag`

   This flag determines whether to convert the generated swagger.json to OpenAPI v3 format. It is enabled by default.

2. `disableOverwriteSwaggerV2Flag`

   This flag determines whether to generate OpenAPI v3.0 json without overwriting the original swagger.json. It is disabled by default.

3. `openapiOutputDirFlag`

   This flag specifies the output directory for the generated OpenAPI v3 spec. The default value is `./openapi`.

## RoadMap

1. Convert and merge multiple Swagger JSON files into a single OpenAPI JSON file.
