#!/bin/bash

cd ./cli
go build -o ../interstellar -ldflags '-w -s' -trimpath ./interstellar