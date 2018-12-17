package main

import (
	"log"
	"up"

	"github.com/spf13/viper"
)

func main() {
	var config up.Config
	var err error

	viper.SetConfigName("up")
	viper.AddConfigPath("/etc/up/")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal(err)
	}

	server, err := up.New(&config)
	if err != nil {
		log.Fatal(err)
	}

	defer server.Shutdown()
	go server.Run()
	<-server.Stop
}
