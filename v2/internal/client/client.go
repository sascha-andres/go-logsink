// Copyright © 2017 Sascha Andres <sascha.andres@outlook.com>
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

package client

import (
	"bufio"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"

	pb "github.com/sascha-andres/go-logsink/v2/logsink"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"go.starlark.net/starlark"
)

var (
	linePrefix             string
	thread                 *starlark.Thread
	filterFunction         starlark.Value
	filterFunctionProvided bool
)

// setupFilter creates the Starlark filter function if provided
func setupFilter() error {
	ff := viper.GetString("connect.filter-function")
	if ff == "" {
		return nil
	}

	// Execute Starlark program in a file.
	thread = &starlark.Thread{Name: "filter thread"}
	globals, err := starlark.ExecFile(thread, ff, nil, nil)
	if err != nil {
		return err
	}

	if _, ok := globals["filter"]; !ok {
		return errors.New("function filter not found")
	}

	// Retrieve a module global.
	filterFunction = globals["filter"]

	filterFunctionProvided = true

	return nil
}

// filtered returns true if a filter function is provided and it evaluates to true
func filtered(line string) bool {
	if !filterFunctionProvided {
		return false
	}
	v, err := starlark.Call(thread, filterFunction, starlark.Tuple{starlark.String(line)}, nil)
	if err != nil {
		return false
	}
	return v == starlark.True
}

func setup() string {
	if !viper.GetBool("connect.pass-through") {
		log.Printf("Connecting to %s\n", viper.GetString("connect.address"))
	}
	return viper.GetString("connect.prefix")
}

// Connect is used to connect to a go-logsink server
func Connect() {
	linePrefix = setup()
	err := setupFilter()
	if err != nil {
		log.Fatalf("could not setup filter: %s", err)
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(viper.GetString("connect.address"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	c := pb.NewLogTransferClient(conn)

	scanner := bufio.NewScanner(os.Stdin)
	client, err  := c.SendLine(context.Background())
	if err != nil {
		log.Panicf("error creating client to send log entries: %s", err)
	}
	for scanner.Scan() {
		var (
			err error
		)
		content := scanner.Text()
		if filtered(content) {
			continue
		}
		if viper.GetBool("connect.pass-through") {
			fmt.Println(content)
		}

		if "" == linePrefix {
			err = client.Send(&pb.LineMessage{Line: content, Priority: int32(viper.GetInt("connect.priority"))})
		} else {
			err = client.Send(&pb.LineMessage{Line: fmt.Sprintf("[%s] %s", linePrefix, content), Priority: int32(viper.GetInt("connect.priority"))})
		}
		if nil != err {
			log.Fatal(err)
		}
	}
	res, err := client.CloseAndRecv()
	if !(nil != res && res.Result) || nil != err {
		log.Fatal(err)
	}
	if err := scanner.Err(); err != nil {
		log.Warnf("reading standard input:", err)
	}
}
