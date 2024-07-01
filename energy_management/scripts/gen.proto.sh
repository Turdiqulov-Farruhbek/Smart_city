#!/bin/sh

CURRENT_DIR=$1
rm -rf ${CURRENT_DIR}/genproto
mkdir -p ${CURRENT_DIR}/genproto

# Exclude hidden directories (those starting with .) from the find results
for x in $(find ${CURRENT_DIR}/restaurant_proto* -type d ! -path '*/\.*'); do
  # Compile .proto files if they exist in the directory
  if ls ${x}/*.proto 1> /dev/null 2>&1; then
    protoc -I=${x} -I/usr/local/include \
      --go_out=${CURRENT_DIR}/genproto --go_opt=paths=source_relative \
      --go-grpc_out=${CURRENT_DIR}/genproto --go-grpc_opt=paths=source_relative \
      ${x}/*.proto
  fi
done
