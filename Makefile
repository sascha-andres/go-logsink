application: protobuf
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s'

protobuf:
	protoc -I logsink/ logsink/logsink.proto --go_out=plugins=grpc:logsink
