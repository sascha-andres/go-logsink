## go-logsink completion powershell

Generate the autocompletion script for powershell

### Synopsis

Generate the autocompletion script for powershell.

To load completions in your current shell session:

	go-logsink completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.


```
go-logsink completion powershell [flags]
```

### Options

```
  -h, --help              help for powershell
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
      --config string     config file (default is $HOME/.go-logsink.yaml)
      --debug             Enable debug mode (statsviz)
      --lockfile string   Specify a lockfile
```

### SEE ALSO

* [go-logsink completion](go-logsink_completion.md)	 - Generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 8-Mar-2025
