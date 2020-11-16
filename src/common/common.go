package common

import (
	"log"

	"github.com/jinzhu/gorm"
)

type GormDB struct {
	*gorm.DB
}

func Connect(c Config) GormDB {
	db, err := gorm.Open("postgres", c.connectionString())
	if err != nil {
		log.Fatalf("failed to connect database %v", err)
	}

	db.LogMode(c.Debug)

	return GormDB{DB: db}
}

func (g GormDB) Stop() {
	log.Println("Closing DB...")

	err := g.Close()
	if err != nil {
		log.Printf("Close DB: %v", err)
	}
}
