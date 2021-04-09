package database

import (
	"log"
	"sync"

	"go-backoffice-seller-api/src/entities"

	"go-backoffice-seller-api/src/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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
	url := "host=localhost user=postgres password=doismundos dbname=ownshop port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(config.Engine, url)
	if err != nil {
		log.Println("Database connection failed : ", err)
	} else {
		log.Println("Database connection established!")
	}
	log.Println("Database connection running")
	g.client = db
}

func InitTestDatabase(g *gormDatabase) {
	db, err := gorm.Open("sqlite3", "database.db")

	// db.LogMode(true)

	if err != nil {
		log.Println("Database connection failed : ", err)
	} else {
		log.Println("Database connection established!")
	}
	g.client = db
}

// Making sure gormClient only initialise once as singleton
func (g *gormDatabase) GetDatabase(config config.Database, env string) *gorm.DB {
	if g.client == nil {
		g.once.Do(func() {
			if env == "test" {
				InitTestDatabase(g)
			} else {
				InitDatabase(g, &config)
			}
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
