#!/bin/bash

set -e

go generate ./...
rm -rf internal/api/grpc/pb
mkdir -p internal/api/grpc/pb
protoc \
      --proto_path=api/ \
      --go_out=internal/api/grpc/pb \
      --go-grpc_out=internal/api/grpc/pb \
      api/gatekeeper.proto
