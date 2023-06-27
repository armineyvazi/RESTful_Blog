package main

import (
	"restful_blog/cmd/server"
	"restful_blog/database/migrations"
)

func main() {
	// Migration start
	migrations.MigrateModels()
	//start server
	server.StartServer()
}
