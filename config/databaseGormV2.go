package config

import (
	"fmt"
	"os"

	logger "github.com/davidyap2002/user-go/logger"

	postgres "gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
)

//Generated By github.com/davidyap2002/GormCrudGenerator

//ConnectGorm Database Connection to Gorm V2
func ConnectGorm() *gorm.DB {

	dsn := os.Getenv("POSTGRE_URI")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), logger.InitConfig())

	if err != nil {
		fmt.Println(err)
		panic("Fail To Connect Database")
	}

	return db
}
