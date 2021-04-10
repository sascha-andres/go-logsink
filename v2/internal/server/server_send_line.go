package web

//// SendLine implements logsink.SendLine
//func (s *server) SendLine(stream logsink.LogTransfer_SendLineServer) error {
//	for {
//		in, err := stream.Recv()
//		if err == io.EOF {
//			return stream.SendAndClose(&logsink.Empty{})
//		}
//		if err != nil {
//			logrus.Warnf("error reading request: %#v", err)
//			return stream.SendAndClose(&logsink.Empty{})
//		}
//
//		s.numberOfLines.Inc()
//		breakAt := viper.GetInt("web.break")
//		priority := int32(math.Max(0, math.Min(9, float64(in.Priority))))
//
//		if viper.GetBool("debug") {
//			fmt.Println(in.Line)
//		}
//		if breakAt == 0 {
//			s.broadcastLine(in.Line, priority)
//		} else {
//			iterations := int(len(in.Line) / breakAt)
//			for start := 0; start <= iterations; start++ {
//				s.broadcastLine(in.Line[start*breakAt:int32(math.Min(float64((start+1)*breakAt), float64(len(in.Line))))], priority)
//			}
//		}
//	}
//	return stream.SendAndClose(&logsink.Empty{})
//}
