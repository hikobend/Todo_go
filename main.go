package main

import (
	"fmt"
	"hello/app/models"
)

func main() {
	// init関数を呼びだす
	fmt.Println(models.Db)

	// ポインタ呼び出し
	// u := &models.User{}
	// u.Name = "kohei"
	// u.Email = "kohei@test.com"
	// u.PassWord = "koheidesu"
	// uを表示
	// fmt.Println(u)

	// ユーザーを作成
	// u.CreateUser()

	u, _ := models.GetUser(1)

	fmt.Println(u)

	u.Name = "kohei2"
	u.Email = "kohei2@test.com"
	u.UpdateUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)

	u.DeleteUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)
}
