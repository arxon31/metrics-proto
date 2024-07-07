# Write makefile to compile protofiles in golang grpc server and client

.PHONY: proto
proto:
	protoc ./proto/*.proto --go_out=./pkg && \
        protoc ./proto/*.proto --go-grpc_out=./pkg
