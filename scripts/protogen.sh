#!/bin/bash

[ -z "$(which protoc)" ] && { echo "install protoc" ; exit 1 ; }

for dir in ./proto/protos/*; do 
    protoc -I ../cosmos-sdk/proto  -I ../gogoproto -I ../cosmos-proto/proto -I ../ibc-go/proto \
        --go_out=. --go-grpc_out=. --proto_path . ${dir}/*
done
cp -r github.com/doggystylez/interstellar/proto/* ./proto/
rm -rf github.com

cd proto
go get github.com/cosmos/cosmos-proto@v1.0.0-beta.2
go mod tidy