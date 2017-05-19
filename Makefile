linux:
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s' -o go-logsink
osx:
	CGO_ENABLED=0 GOOS=darwin go build -a -ldflags '-s' -o go-logsink.osx
windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -ldflags '-s' -o go-logsink.exe
protobuf:
	protoc -I logsink/ logsink/logsink.proto --go_out=plugins=grpc:logsink

all: linux osx windows
