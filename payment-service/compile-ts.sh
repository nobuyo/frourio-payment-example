PLUGIN_TS=../frontend/node_modules/.bin/protoc-gen-ts
PLUGIN_GRPC=../frontend/node_modules/.bin/grpc_tools_node_protoc_plugin
DIST_DIR=./client

protoc \
--js_out=import_style=commonjs,binary:"${DIST_DIR}"/ \
--ts_out=import_style=commonjs,binary:"${DIST_DIR}"/ \
--grpc_out="${DIST_DIR}"/ \
--plugin=protoc-gen-grpc="${PLUGIN_GRPC}" \
--plugin=protoc-gen-ts="${PLUGIN_TS}" \
--proto_path=./proto \
-I $DIST_DIR \
./proto/*.proto
