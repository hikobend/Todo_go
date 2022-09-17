package main

import (
	"fmt"
	"hello/app/models"
)

func main() {
	fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "kohei"
	// u.Email = "kohei@test.com"
	// u.PassWord = "koheidesu"
	// fmt.Println(u)

	// u.CreateUser()

	u, _ := models.GetUser(1)

	fmt.Println(u)
}
