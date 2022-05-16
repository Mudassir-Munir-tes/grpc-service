package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Mudassir-Munir-tes/grpc-service/models"
)

var (
	Db  *gorm.DB
	Err error
)

const (
	constr = "host=172.17.0.2 user=postgres password=postgres port=5432 dbname=postgres sslmode=disable"
)

func Connect() {

	Db, Err = gorm.Open(postgres.Open(constr), &gorm.Config{})
	if Err != nil {
		panic("failed connection")
	}

	Db.AutoMigrate(&models.User{})
	Db.AutoMigrate(&models.Driver{})
	Db.AutoMigrate(&models.Truck{})

}
