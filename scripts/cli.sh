#!/bin/bash

cd ./cli
go mod tidy
go build -o ../interstellar -ldflags '-w -s' -trimpath ./interstellar