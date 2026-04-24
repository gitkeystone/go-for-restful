# 启动 Swagger UI
```bash
docker run --rm --network=host -e SWAGGER_JSON=/app/openapi.json -v .:/app swaggerapi/swagger-ui
```
