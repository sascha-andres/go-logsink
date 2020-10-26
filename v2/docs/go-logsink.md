## go-logsink

go-logsink is a simplistic log aggregator

### Synopsis


You can use go-logsink to combine multiple log streams
into one. For example you can combine multiple tails into one
output.

To do this start a server and connect any number of clients.

### Options

```
      --config string     config file (default is $HOME/.go-logsink.yaml)
      --lockfile string   Specify a lockfile
```

### SEE ALSO
* [go-logsink connect](go-logsink_connect.md)	 - Connect to a go-logsink server and forward stdin
* [go-logsink doc](go-logsink_doc.md)	 - Write out standard documentation
* [go-logsink listen](go-logsink_listen.md)	 - Start a server instance of go-logsink
* [go-logsink web](go-logsink_web.md)	 - Start a server instance with a web interface

###### Auto generated by spf13/cobra on 9-Oct-2019