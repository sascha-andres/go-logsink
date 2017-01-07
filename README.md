# go-logsink

[![Go Report Card](https://goreportcard.com/badge/github.com/sascha-andres/go-logsink)](https://goreportcard.com/report/github.com/sascha-andres/go-logsink) [![codebeat badge](https://codebeat.co/badges/6e2d5bf5-5ca2-41a3-842d-631ba32d196c)](https://codebeat.co/projects/github-com-sascha-andres-go-logsink)

## What is go-logsink

go-logsink is a client and server application to aggregate multiple steams into one.

It does so by sending all data piped into the client to the server.

## Usage

To start a server just type `go-logsink listen`. By default the server listens on port 50051. You can change the default binding definition using the `bind` flag:

    go-logsink listen --bind ":55555"

In this sample, the server would listen on port 55555.

To send data to the server you have to start at least one client. For example to send the syslog to the server:

    sudo tail -f /var/log/syslog | go-logsink connect --address "localhost:55555"

Using the `address` flag it is possible to send data to the non default destination (`localhost:50051`)

An advanced usage would be to forward all logs from running docker containers:

    docker logs -f $(docker ps -q) | go-logsink connect &

This assumes a runnint go-logsink server at localhost:50051