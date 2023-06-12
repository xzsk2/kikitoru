package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"kikitoru/config"
	"kikitoru/internal/database"
	"kikitoru/internal/router"
	"kikitoru/logs"
)

func main() {
	logs.InitLogger()
	config.InitConfig()
	database.InitDataBase()

	r := router.InitRouter()
	var addr string
	if config.C.BlockRemoteConnection {
		addr = fmt.Sprintf("localhost:%d", config.C.ListenPort)
	} else {
		addr = fmt.Sprintf(":%d", config.C.ListenPort)
	}
	err := r.Run(addr)
	if err != nil {
		log.Fatal(err)
	}

}
