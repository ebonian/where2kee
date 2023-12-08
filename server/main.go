package main

import (
	"github.com/ebonian/where2kee/server/cmd/server"
	"github.com/ebonian/where2kee/server/pkg/config"
)

func main() {
	config.LoadAllConfigs(".env")
	server.Serve()
}
