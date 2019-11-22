package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/philmacfly/gopherrator/feedtypes"
)

type config struct {
	Gopherpath     string
	Outletfile     string
	ItemsperOutlet int
}

func loadconfig(input io.Reader) (config, error) {
	var result config
	decoder := json.NewDecoder(input)
	err := decoder.Decode(&result)
	return result, err
}

func main() {
	var configpath string
	flag.StringVar(&configpath, "config", "config.json", "Complete path to a config file")
	flag.Parse()
	c, err := os.Open(configpath)
	if err != nil {
		log.Fatalln("Error opening config file:", err.Error())
	}
	defer c.Close()
	cfg, err := loadconfig(c)
	if err != nil {
		log.Fatalln("Error loading config file:", err.Error())
	}
	of, err := os.Open(cfg.Outletfile)
	if err != nil {
		log.Fatalln("Error opening Outletfile:", err.Error())
	}
	defer of.Close()
	outlets, err := feedtypes.LoadOutlets(of)
	if err != nil {
		log.Fatalln("Error loading Outletfile:", err.Error())
	}
	newsfeeds, err := feedtypes.ParseOutlets(outlets)
	if err != nil {
		log.Fatalln("Error parsing Outlets", err.Error())
	}
	fmt.Println(newsfeeds)
	err = feedtypes.WriteEntriestoGopher(newsfeeds, cfg.ItemsperOutlet, cfg.Gopherpath)
	if err != nil {
		log.Fatalln("Error writing Gopher", err.Error())
	}
}
