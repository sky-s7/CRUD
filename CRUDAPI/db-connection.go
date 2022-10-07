package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *gorm.DB

var urlDSN = "root:admin@tcp(127.0.0.1:3306)/mydb"

//clientOption := options.Client().ApplyURL("mongodb://localhost:27017")

var err error

func DataMigaration() {

	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})

	//client, err := mongo.Connect(context.TODO(),clientOption)

	if err != nil {

		fmt.Print(err.Error())
		panic("Connectinon Error :(")

	}
	Database.AutoMigrate(&Employee{})
}
