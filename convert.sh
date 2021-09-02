#!/usr/bin/env bash
npx typeconv -f ts -t oapi -o tmp api.ts
sed -i '/$id/d' tmp/api.yaml
sed -i '/$comment/I,+3 d' tmp/api.yaml
sed -i '/paths: {}/d' tmp/api.yaml
sed -i 's*schema/*schemas/*g' tmp/api.yaml
cat ./paths.yaml >> tmp/api.yaml
rm -rf go
npx @openapitools/openapi-generator-cli generate -i tmp/api.yaml -g go -o go
