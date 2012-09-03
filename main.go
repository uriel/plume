package main

import (
	"log"
	"plume/configuration"
)

func main() {
	conf, err := configuration.Load("config.json")

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Print(conf)
}
