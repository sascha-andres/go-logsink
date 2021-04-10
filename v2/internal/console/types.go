package console

import pb "github.com/sascha-andres/go-logsink/v2/logsink"

type (
	LinePrinter = func(line *pb.LineMessage)
)
