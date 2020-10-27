#!/bin/sh

# create access token
protoc --go_out=. / 
--go-grpc_out=. / 
--go-grpc_opt=paths=source_relative /
--go_opt=paths=source_relative databases/proto/kafkapb/kafkapb.proto 
