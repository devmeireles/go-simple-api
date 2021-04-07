package database

import (
	"log"
	"sync"

	"go-backoffice-seller-api/src/entities"

	"go-backoffice-seller-api/src/config"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type gormDatabase struct {
	client *gorm.DB
	once   sync.Once
}

func NewGormDatabase() IDatabaseEngine {
	return &gormDatabase{}
}

func InitDatabase(g *gormDatabase, config *config.Database) {
	// url := config.User + ":" + config.Password + "@tcp(" + config.Server + ":" +
	// config.Port + ")/" + config.Name + "?charset=utf8mb4&parseTime=True&loc=Local"

	url := "host=localhost user=postgres password=doismundos dbname=shop port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(config.Engine, url)
	if err != nil {
		log.Println("Database connection failed : ", err)
	} else {
		log.Println("Database connection established!")
	}
	log.Println("MySql connection running on port 3306")
	g.client = db
}

// Making sure gormClient only initialise once as singleton
func (g *gormDatabase) GetDatabase(config config.Database) *gorm.DB {
	if g.client == nil {
		g.once.Do(func() {
			InitDatabase(g, &config)
		})
	}
	return g.client
}

func (g *gormDatabase) RunMigration() {
	if g.client == nil {
		panic("Initialise gorm db before running migrations")
	}
	g.client.AutoMigrate(&entities.User{})
}
