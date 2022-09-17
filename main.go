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
	// u.Name = "kohei2"
	// u.Email = "kohei2@test.com"
	// u.PassWord = "koheidesu"
	// uを表示
	// fmt.Println(u)

	// ユーザーを作成
	// u.CreateUser()

	// // 作成したidが1番を取ってくる
	// u, _ := models.GetUser(1)

	// // ユーザーを出力する
	// fmt.Println(u)

	// // ユーザー更新
	// u.Name = "kohei2"
	// u.Email = "kohei2@test.com"
	// // ユーザー更新
	// u.UpdateUser()
	// // ユーザー取得
	// u, _ = models.GetUser(1)
	// // ユーザー表示
	// fmt.Println(u)

	// // ユーザー削除
	// u.DeleteUser()
	// // ユーザー表示
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// id : 2にfirst todoを追加
	// user, _ := models.GetUser(2)
	// user.CreateTodo("first todo")

	// first todoを取得
	// t, _ := models.GetTodo(1)
	// tを表示
	// fmt.Println(t)

	// user, _ := models.GetUser(3)
	// user.CreateTodo("Third Todo")

	// todos, _ := models.GetTodos()
	// 複数表示
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// user2, _ := models.GetUser(3)
	// todos, _ := user2.GetTodoByUser()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// t, _ := models.GetTodo(1)
	// t.Content = "Update Todo"
	// t.UpdateTodo()

	t, _ := models.GetTodo(4)
	t.DeleteTodo()
}
