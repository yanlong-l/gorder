package main

import (
	"github.com/spf13/viper"
	"github.com/yanlong-l/gorder/common/config"
	"log"
)

func init() {
	err := config.NewViperConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	log.Printf("%v", viper.Get("order"))
}
