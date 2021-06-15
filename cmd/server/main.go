package main

import "log"

func main() {
	flag.Init()

	config, err := config.ReadNewConfig()
	if err != nil {
		log.Fatal("Can't read config file")
	}
	
}
