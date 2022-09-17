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

	// 作成したidが1番を取ってくる
	u, _ := models.GetUser(1)

	// ユーザーを出力する
	fmt.Println(u)

	// ユーザー更新
	u.Name = "kohei2"
	u.Email = "kohei2@test.com"
	// ユーザー更新
	u.UpdateUser()
	// ユーザー取得
	u, _ = models.GetUser(1)
	// ユーザー表示
	fmt.Println(u)

	// ユーザー削除
	u.DeleteUser()
	// ユーザー表示
	u, _ = models.GetUser(1)
	fmt.Println(u)
}
