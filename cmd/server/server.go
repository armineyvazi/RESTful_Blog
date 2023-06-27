package server

import (
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"

	"restful_blog/config"
	"restful_blog/router"
)

func StartServer() {
	config.InitEnvConfigs()

	e := echo.New()
	//read banner file
	bs, _ := os.ReadFile("cmd/server/banner.txt")
	fmt.Println(string(bs))

	router.Router(e)

	log.Fatal(e.Start(":" + config.EnvConfigs.LocalServerPort))

}
