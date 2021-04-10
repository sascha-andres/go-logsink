package console

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/sascha-andres/go-logsink/v2/logsink"
)

//coloredPrinter prints the line colored according to priority
func coloredPrinter(line *logsink.LineMessage) {
	errorColor := color.New(color.FgRed)
	warningColor := color.New(color.FgYellow)
	if line.Priority > 3 {
		_, _ = errorColor.Println(line.Line)
	} else if line.Priority == 2 || line.Priority == 3 {
		_, _ = warningColor.Println(line.Line)
	} else {
		fmt.Println(line.Line)
	}
}
