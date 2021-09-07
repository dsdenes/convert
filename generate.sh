#!/usr/bin/env bash
GEN_FOLDER="./generated"
GEN_GO_FOLDER="./generated/go"
GEN_TS_FOLDER="./generated/ts"

echo "Removing old generated codes in ${GEN_FOLDER}"
rm -rf "${GEN_FOLDER}"
mkdir -p "${GEN_GO_FOLDER}"
mkdir -p "${GEN_TS_FOLDER}"

echo "Generating TS code to ${GEN_TS_FOLDER}"
npx @openapitools/openapi-generator-cli generate -i api.yaml -g typescript-fetch -o ${GEN_TS_FOLDER}

echo "Getting github.com/deepmap/oapi-codegen..."
go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen

echo "Generating Go code to ${GEN_GO_FOLDER}"
oapi-codegen -package openapi_gen -generate types openapi.yml > "${GEN_GO_FOLDER}/types.gen.go"
oapi-codegen -package openapi_gen -generate client openapi.yml > "${GEN_GO_FOLDER}/client.gen.go"
oapi-codegen -package openapi_gen -generate chi-server openapi.yml > "${GEN_GO_FOLDER}/server.gen.go"

