// Copyright © 2021 Sascha Andres <sascha.andres@outlook.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"context"
	pb "github.com/sascha-andres/go-logsink/v2/logsink"
	"github.com/sirupsen/logrus"
)

//SendLine implements logsink.SendLine
// SendLine implements logsink.SendLine
func (s *server) SendLine(context context.Context, in *pb.LineMessage) (*pb.LineResult, error) {
	logrus.Println(in.Line)

	_ = context

	return &pb.LineResult{
		Result: true,
	}, nil
}
