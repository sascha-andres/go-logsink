#! /bin/bash

which protoc
if [ $? != 0 ]; then
	echo "Protocol buffer installation not found"
	exit 1
fi

which protoc-gen-go
if [ $? != 0 ]; then
	echo "Installing go extension for protocol buffers"
	go get -u github.com/golang/protobuf/protoc-gen-go
fi
