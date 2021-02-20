package web

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/sascha-andres/go-logsink/v2/logsink"
)

// server is used to implement logsink.LogTransferServer.
type (
	server struct {
		pb.UnimplementedLogTransferServer
		hub           *hub
		numberOfLines int64
	}
	lineType struct {
		Line     string
		Priority int32
		Key      string
	}
)

// SendLine implements logsink.SendLine
func (s *server) SendLine(stream pb.LogTransfer_SendLineServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		numberOfLines.Inc()
		breakAt := viper.GetInt("web.break")
		prio := int32(math.Max(0, math.Min(9, float64(in.Priority))))
		if viper.GetBool("debug") {
			fmt.Println(in.Line)
		}
		if breakAt == 0 {
			s.broadcastLine(in.Line, prio)
		} else {
			iterations := int(len(in.Line) / breakAt)
			for start := 0; start <= iterations; start++ {
				s.broadcastLine(in.Line[start*breakAt:int32(math.Min(float64((start+1)*breakAt), float64(len(in.Line))))], prio)
			}
		}
		stream.Send(&pb.LineResult{
			Result:               true,
			Sequence:             in.Sequence,
		})
	}
}

func (s *server) broadcastLine(line string, priority int32) {
	s.numberOfLines++
	obj := lineType{
		Line:     line,
		Priority: priority,
		Key:      fmt.Sprintf("%d", s.numberOfLines),
	}
	if line == "" {
		obj.Line = " "
	}
	data, err := json.Marshal(obj)
	if err != nil {
		log.Errorf("error creating message for websocket: %s", err)
	} else {
		s.hub.broadcast <- data
	}
}

func (s *server) run() {
	s.hub = newHub()
	go s.hub.run()
	lis, err := net.Listen("tcp", viper.GetString("web.bind"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	pb.RegisterLogTransferServer(srv, s)
	// Register reflection service on gRPC server.
	reflection.Register(srv)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
