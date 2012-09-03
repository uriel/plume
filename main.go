package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"plume/configuration"
	"plume/theme"
)

func main() {
	var conf_path string
	flag.StringVar(&conf_path, "conf", "config.json", "Location of configuration file. Defaults to config.json in the present working directory.")
	flag.Parse()

	conf, err := configuration.Load(conf_path)

	if err != nil {
		log.Fatal(err)
		return
	}

	theme_path := filepath.Join("themes", conf.Theme, "index.html")
	tmpl, err := theme.Load(conf.Theme, theme_path)

	if err != nil {
		log.Fatal(err)
		return
	}

	tmpl.Execute(os.Stdout, conf)
}
