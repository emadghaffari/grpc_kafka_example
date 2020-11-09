#!/bin/sh

# read document for golang
# https://github.com/grpc/grpc-go

# elasticsearch
protoc --go_out=plugins=grpc:. elasticsearchpb/elasticsearch.proto

# ping
protoc --go_out=plugins=grpc:. pingpb/ping.proto

# twitter
protoc --go_out=plugins=grpc:. twitterpb/twitter.proto

# protoc -I . \
#    --go_out . --go_opt paths=source_relative \
#    --go-grpc_out . --go-grpc_opt paths=source_relative \
#    -.proto