package main

import (
	"github.com/oliverperboni/GoApi/config"
	"github.com/oliverperboni/GoApi/router"
)

func main() {
	router.Initialize()
	config.ConnectDatabase()

}
