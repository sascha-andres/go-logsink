linux: docs    ## build for linux
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s' -o go-logsink

osx: docs      ## build for osx
	CGO_ENABLED=0 GOOS=darwin go build -a -ldflags '-s' -o go-logsink.osx

windows: docs  ## build for windows
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -ldflags '-s' -o go-logsink.exe

protobuf: ## run protobuf stuff
	./check_protobuf.sh
	protoc -I logsink/ logsink/logsink.proto --go_out=plugins=grpc:logsink

ui:       ## create ui
	cd ui && npm install
	cd ui && npm run build

docs: ## generate docs
	@go run main.go doc

statik: ui  ## create statik src
	-rm internal/web/statik/statik.go
	cd internal/web && statik -src=../../ui/dist

all: linux osx windows  ## build all

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
.PHONY: ui statik windows linux osx protobuf docs
