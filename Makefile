application: protobuf
	-mkdir -p build/linux_amd64
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s' -o build/linux_amd64/gologsink
	tar cvzf build/linux_amd64.tgz build/linux_amd64/go-logsink
osx:
	-mkdir -p build/darwin
	CGO_ENABLED=0 GOOS=darwin go build -a -ldflags '-s' -o build/darwin/go-logsink
	tar cvzf build/darwin.tgz build/darwin/go-logsink
windows:
	-mkdir -p build/windows
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -ldflags '-s' -o build/windows/go-logsink.exe
	zip build/windows.zip build/windows/go-logsink.exe
protobuf:
	protoc -I logsink/ logsink/logsink.proto --go_out=plugins=grpc:logsink
