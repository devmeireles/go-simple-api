package main

import (
	"encoding/json"
	configs "go-backoffice-seller-api/src/config"
	"go-backoffice-seller-api/src/database"
	"go-backoffice-seller-api/src/routes"
	router "go-backoffice-seller-api/src/utils"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

var (
	config     configs.Config
	httpRouter router.IRouter
	gormDb     database.IDatabaseEngine
	gDb        *gorm.DB
)

func main() {
	startServer()
}

func startServer() {
	initConfig()
	httpRouter = router.NewMuxRouter()
	httpRouter.ADDVERSION("/api/v1")

	gormDb = database.NewGormDatabase()
	gDb = gormDb.GetDatabase(config.Database, "dev")
	gormDb.RunMigration()

	initRoutes()

	httpRouter.SERVE(config.App.Port)
}

func initConfig() {
	file, err := os.Open("./config.json")
	if err != nil {
		log.Printf("No ./config.json file found!! Terminating the server, error: %s\n", err.Error())
		panic("No config file found! Error : " + err.Error())
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Printf("Error occurred while decoding json to config model, error: %s\n", err.Error())
		panic(err.Error())
	}
}

func initRoutes() {
	routes.UserRoute(gDb, httpRouter)
	routes.AuthRoute(gDb, httpRouter)
}
