#!/bin/sh

# create access token
protoc --go_out=. / 
--go-grpc_out=. / 
--go-grpc_opt=paths=source_relative /
--go_opt=paths=source_relative databases/proto/kafkapb/kafkapb.proto 


# protoc --proto_path=IMPORT_PATH --cpp_out=DST_DIR --java_out=DST_DIR --python_out=DST_DIR --go_out=DST_DIR --ruby_out=DST_DIR --objc_out=DST_DIR --csharp_out=DST_DIR path/to/file.proto