package main

import (
	"fmt"
	"hello/app/controllers"
	"hello/app/models"
)

func main() {
	// init関数を呼びだす
	fmt.Println(models.Db)

	// 関数呼び出し
	controllers.StartMainServer()

	// GetUserByEmail動作確認
	// user, _ := models.GetUserByEmail("kohei@kohei.com")
	// fmt.Println(user)

	// CreateSession動作確認
	// sessionを利用して作成
	// session, err := user.CreateSession()
	// if err != nil {
	// 	log.Panicln(err)
	// }
	// fmt.Println(session)

	// CheckSession動作確認
	// valid, _ := session.CheckSession()
	// fmt.Println(valid)
}
