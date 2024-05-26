# swagger2openapi3

🌍 *[English](README.md) ∙ [简体中文](README_zh-CN.md)*

Swagger2openapi3 提供了一个包，可以将Swagger 2.0规范的JSON和YAML转换为OpenAPI 3.0.1。它还提供了一个工具叫做swag2op，它集成了Swagger 2.0的生成，并支持转换为OpenAPI 3.0.3。

## 开始使用

1. 在你的API源代码中添加注释，参见Declarative Comments Format。
2. 通过以下命令安装swag：

```
go install github.com/zxmfke/swagger2openapi3/cmd/swag2op@latest
```

要从源代码构建，你需要安装Go（1.19或更高版本）。

3. 在包含main.go文件的项目根文件夹中运行 swag2op init。这将解析你的注释并生成所需的文件（docs文件夹和docs/docs.go）。

   ```bash
   swag2op init
   ```
   
确保导入生成的docs/docs.go，以便初始化你的特定配置。如果你的通用API注释不在main.go中，你可以使用-g标志告诉swag。

```go
import _ "example-module-name/docs"
```

```bash
swag2op init -g http/api.go
```

4. 下一节将介绍额外的 flag。

## swag2op CLI

swag2op 提供了三个额外的标志，以提供更多的灵活性：

1. `disableConvertOpenApiV3Flag`

   这个标志决定是否将生成的swagger.json转换为OpenAPI v3格式。默认情况下是启用的。

2. `disableOverwriteSwaggerV2Flag`

   这个标志决定是否在生成OpenAPI v3.0 json时不覆盖原始的swagger.json。默认情况下是禁用的。

3. `openapiOutputDirFlag`

   这个标志指定生成的OpenAPI v3规范的输出目录。默认值是./openapi。

## 路线图

1. 将多个Swagger JSON文件转换并合并为一个OpenAPI JSON文件。