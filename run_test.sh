#!/bin/bash

go test ./... -cover

# go test -coverprofile=cover.out
# go tool cover -html=cover.out -o coverage.html

# go test -run=Dollar -v

# go test -short // skip time consuming test