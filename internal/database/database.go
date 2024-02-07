package database

import (
	"bwa-startup/config"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgre(cfg *config.Database) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	var i = 0
	for {
		if i >= 5 {
			break
		}
		i++

		db, err = gorm.Open(postgres.Open(cfg.DNS()), &gorm.Config{})

		if err != nil {
			log.Printf("failed to create db connection [retry %d time(s)]", i)
			time.Sleep(5 * time.Second)
			continue
		}

		// test ping database
		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("failed to create db connection [retry %d time(s)]", i)
			time.Sleep(5 * time.Second)
			continue
		}

		if err = sqlDB.Ping(); err != nil {
			log.Printf("failed to create db connection [rety %d time(s)", i)
		} else {
			sqlDB.SetMaxIdleConns(cfg.MaxIdleConnectionsInSecond)
			sqlDB.SetMaxOpenConns(cfg.MaxOpenConnectionsInSecond)
			sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnectionMaxLifetimeInSecond) * time.Second)

			log.Println("Successfully connected to PostgreSQL database")

			return db, nil
		}

	}

	return nil, err
}
