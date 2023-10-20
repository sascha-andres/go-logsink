linux:
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s' -o go-logsink
osx:
	CGO_ENABLED=0 GOOS=darwin go build -a -ldflags '-s' -o go-logsink.osx
windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -ldflags '-s' -o go-logsink.exe
protobuf:
	./check_protobuf.sh
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	protoc -I logsink/ logsink/logsink.proto --go-grpc_out=. --go_out=.
all: linux osx windows
