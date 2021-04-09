package tests

import (
	"encoding/json"
	configs "go-backoffice-seller-api/src/config"
	"go-backoffice-seller-api/src/database"
	"go-backoffice-seller-api/src/entities"
	"go-backoffice-seller-api/src/routes"
	router "go-backoffice-seller-api/src/utils"
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
)

var (
	config     configs.Config
	httpRouter router.IRouter
	gormDb     database.IDatabaseEngine
	gDb        *gorm.DB
)

func TestMain(m *testing.M) {
	initConfig()

	httpRouter = router.NewMuxRouter()
	httpRouter.ADDVERSION("/api/v1")

	gormDb = database.NewGormDatabase()
	gDb = gormDb.GetDatabase(config.Database, "test")
	gormDb.RunMigration()

	routes.UserRoute(gDb, httpRouter)
	routes.AuthRoute(gDb, httpRouter)

	os.Exit(m.Run())
	gDb.DropTable(&entities.User{})
}

func initConfig() {
	file, err := os.Open("../config.json")
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
