package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Link struct {
	Title       string
	Url         string
	Description string
}

type Configuration struct {
	Theme       string
	Name        string
	Description string
	Blogroll    []Link
}

func main() {
	b, err := ioutil.ReadFile("config.json")

	if err != nil {
		log.Fatal(err)
		return
	}

	var m Configuration
	parse_err := json.Unmarshal(b, &m)
	if parse_err != nil {
		log.Fatal(parse_err)
		return
	}

	log.Print(m)
}
