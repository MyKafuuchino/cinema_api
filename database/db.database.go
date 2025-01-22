package database

import (
	"cinema_api/config"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
)

var (
	Db   *gorm.DB
	once sync.Once
)

func InitDb() {
	once.Do(func() {
		var err error
		dbConfig := config.GlobalDbConfig
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbConfig.DbUser,
			dbConfig.DbPassword,
			dbConfig.DbHost,
			dbConfig.DbPort,
			dbConfig.DbName,
		)
		Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

		if err != nil {
			log.Fatalf("Failed to connect to database : %v", err)
		}

		log.Info("Successfully connected to database")
	})
}
