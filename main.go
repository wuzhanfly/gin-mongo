package main

import (
	"flag"
	"gin-mongo-backend/app"
)

func main() {
	configFile := flag.String("c", "config.yaml", "Config file")
	flag.Parse()
	app.Run(*configFile)
}
