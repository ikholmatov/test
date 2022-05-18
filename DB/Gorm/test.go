package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID 		 int
	Name     string
	Age      int
	Birthday time.Time
}

func main() {
	mybase := "user=venom password=112233 dbname=venom sslmode=disable"
	db, err := gorm.Open(postgres.Open(mybase), &gorm.Config{})
	if err != nil {
		fmt.Println(err, "POINT-1")
		return
	}
	//user := User{Name: "Sardor", Age: 22, Birthday: time.Now().AddDate(2002,05,26)}
	// result := db.Create(&user)
	// fmt.Println(user.ID,result.RowsAffected,result.Error)
	// db.Select("Name","Age").Create(&user)
	var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	db.Create(&users)

	for _, user := range users {
  		fmt.Println(user.ID)

}
}