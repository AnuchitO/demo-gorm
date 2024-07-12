package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// _ "modernc.org/sqlite"
)

type Profile struct {
	gorm.Model
	PID       string `gorm:"column:pid"`
	Email     string
	FirstName string
	LastName  string
}

func main() {
	fmt.Println("Hi, gorm")
	db, err := gorm.Open(sqlite.Open("demo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("database:", err)
	}

	db.AutoMigrate(&Profile{})

	p := Profile{
		PID:       "123",
		Email:     "demo@anuchito.com",
		FirstName: "AnuchitO",
		LastName:  "Nong",
	}

	//  Create
	db.Create(&p)

	// Read
	var pro Profile
	db.First(&pro, "pid = ?", "123")
	fmt.Printf("pro: % #v\n", pro)

	// Update
	db.Model(&pro).Updates(Profile{FirstName: "AAAAA"})
	fmt.Printf("pro: % #v\n", pro)

}
