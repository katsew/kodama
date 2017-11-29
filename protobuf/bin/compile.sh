#!/bin/bash

PROTOS=("ping")
PROJECT_NAME=kodama

cd `dirname $0`
cd ../../../

echo "exec dir: $PWD"
echo "Start compiling..."

for proto in ${PROTOS[@]}; do

echo "compile <$proto> start..."

protoc -I . ${PROJECT_NAME}/protobuf/${proto}/*.proto --go_out=plugins=grpc:../../

echo "compile <$proto> done."

done

echo "All compilation done."
