package configuration

import (
	"encoding/json"
	"io/ioutil"
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

func Load(config_path string) (Configuration, error) {
	var config Configuration
	b, read_err := ioutil.ReadFile(config_path)

	if read_err != nil {
		return config, read_err
	}

	parse_err := json.Unmarshal(b, &config)
	if parse_err != nil {
		return config, parse_err
	}

	return config, nil
}
