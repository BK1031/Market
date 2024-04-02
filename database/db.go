package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"market-mono/config"
	"market-mono/utils"
	"time"
)

var DB *gorm.DB

var dbRetries = 0

func InitializeDB(sslmode string, timezone string) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", config.PostgresHost, config.PostgresUser, config.PostgresPassword, config.PostgresDB, config.PostgresPort, sslmode, timezone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		if dbRetries < 10 {
			dbRetries++
			utils.SugarLogger.Errorln("failed to connect database, retrying in 5s... ")
			time.Sleep(time.Second * 5)
			InitializeDB(sslmode, timezone)
		} else {
			utils.SugarLogger.Fatalln("failed to connect database after 15 attempts, terminating program...")
		}
	} else {
		utils.SugarLogger.Infoln("Connected to postgres database")
		if err := db.AutoMigrate(); err != nil {
			utils.SugarLogger.Fatalln("AutoMigration failed: ", err)
		}
		utils.SugarLogger.Infoln("AutoMigration complete")
		DB = db
	}
}
