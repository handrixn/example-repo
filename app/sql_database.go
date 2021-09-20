package app

import (
	"fmt"
	"github.com/handrixn/example-repo/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strconv"
	"time"
)

func NewSQLDatabase() *gorm.DB {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbTimeZone := os.Getenv("DB_TIME_ZONE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s", dbUsername, dbPassword, dbHost, dbPort, dbName, dbTimeZone)

	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		SkipInitializeWithVersion: false,
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	helper.PanicIfError(err)

	sqlDB, err := db.DB()
	helper.PanicIfError(err)

	maxIdleCons, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONS"))
	helper.PanicIfError(err)

	maxOpenCons, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONS"))
	helper.PanicIfError(err)

	maxIdleTime, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_TIME_CONS"))
	helper.PanicIfError(err)

	maxConnMaxLifeTime, err := strconv.Atoi(os.Getenv("DB_MAX_LIFE_TIME"))
	helper.PanicIfError(err)

	sqlDB.SetMaxIdleConns(maxIdleCons)
	sqlDB.SetMaxOpenConns(maxOpenCons)
	sqlDB.SetConnMaxIdleTime(time.Duration(maxIdleTime) * time.Minute)
	sqlDB.SetConnMaxLifetime(time.Duration(maxConnMaxLifeTime) * time.Minute)

	return db
}
