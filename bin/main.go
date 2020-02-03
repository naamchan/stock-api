package main

import (
	"github.com/naamchan/stock-api/routers"
)

func main() {
	startServer()
}

func startServer() {
	routers.Init()
}
