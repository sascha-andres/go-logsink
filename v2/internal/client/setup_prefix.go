package client

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//setupPrefix returns prefix to use
func setupPrefix() string {
	if !viper.GetBool("connect.pass-through") {
		logrus.Printf("Connecting to %s\n", viper.GetString("connect.address"))
	}
	return viper.GetString("connect.prefix")
}
