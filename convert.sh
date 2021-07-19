#!/usr/bin/env bash
npx typeconv -f ts -t oapi api.ts
sed -i '/$id/d' api.yaml
sed -i '/$comment/I,+3 d' api.yaml
npx @openapitools/openapi-generator-cli generate -i api.yaml -g go -o go
