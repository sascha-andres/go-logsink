linux: protobuf
	-mkdir -p build/linux_amd64
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s' -o build/linux_amd64/go-logsink
	cp -R www build/linux_amd64/
	rm build/linux_amd64.tgz
	cd build/linux_amd64/ && tar cvzf ../linux_amd64.tgz *
osx:
	-mkdir -p build/darwin
	CGO_ENABLED=0 GOOS=darwin go build -a -ldflags '-s' -o build/darwin/go-logsink
	cp -R www build/darwin/
	rm build/darwin.tgz
	cd build/darwin && tar cvzf ../darwin.tgz *
windows:
	-mkdir -p build/windows
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -ldflags '-s' -o build/windows/go-logsink.exe
	cp -R www build/windows/
	rm build/windows.zip
	cd build/windows && zip -r ../windows.zip *
protobuf:
	protoc -I logsink/ logsink/logsink.proto --go_out=plugins=grpc:logsink

all: linux osx windows
