#!/bin/bash

GOOS=js GOARCH=wasm go build -o ../server/test.wasm wasm.go