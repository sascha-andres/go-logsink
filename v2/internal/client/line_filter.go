package client

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//lineFilter reduces lines to those which are not filtered out
func lineFilter(in chan<- string) chan<- string {
	out := make(chan string)
	go func() {
		err := setupFilter()
		if err != nil {
			logrus.Fatalf("could not setupPrefix filter: %s", err)
			close(out)
			return
		}
		passThrough := viper.GetBool("connect.pass-through")

		for line := range in {
			if filtered(line) {
				continue
			}
			out <- line
			if passThrough {
				fmt.Println(line)
			}
		}
		close(out)
	}()

	return out
}

