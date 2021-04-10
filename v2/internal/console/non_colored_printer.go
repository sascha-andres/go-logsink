package console

import (
	"fmt"
	"github.com/sascha-andres/go-logsink/v2/logsink"
)

//nonColoredPrinter prints the line colored according to priority
func nonColoredPrinter(line *logsink.LineMessage) {
	fmt.Println(line.Line)
}
