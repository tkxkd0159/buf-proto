#!/usr/bin/env bash

cd proto
find . -maxdepth 1 -name '*.proto' -print0 |
while IFS= read -r -d '' file; do
  if grep "option go_package" "$file" &> /dev/null ; then
    protoc --go_out=".." "$file"
    fi
  done

cd ..
cp -r github.com/tkxkd0159/buf-proto/grpc/* ./
rm -rf github.com

go mod tidy

