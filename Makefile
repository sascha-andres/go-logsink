protobuf:
	protoc -I logsink/ logsink/logsink.proto --go_out=plugins=grpc:logsink
