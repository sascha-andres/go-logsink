#! /bin/bash

which protoc
if [ $? != 0 ]; then
	echo "Protocol buffer installation not found"
	exit 1
fi