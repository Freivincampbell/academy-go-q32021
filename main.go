package main

import (
	"academy-go-q32021/config"
	"academy-go-q32021/infrastructure/router"
	"academy-go-q32021/registry"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"log"
)

func main() {
	config.ReadConfig()

	r := registry.NewRegistry()

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
