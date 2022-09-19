package controllers

import (
	"net/http"
)

// func top(w http.ResponseWriter, r *http.Request) {
// layout topを実行
func top(w http.ResponseWriter, r *http.Request) {
	// sessionを使って、cookieを取得
	_, err := session(w, r)
	// 存在しない場合
	if err != nil {
		// 全体で読み込むtemplateたち
		generateHTML(w, "Hello", "layout", "public_navbar", "top")
		// ログインしている
	} else {
		// todosにRedirect
		http.Redirect(w, r, "/todos", 302)
	}
}

// indexを表示する関数作成
func index(w http.ResponseWriter, r *http.Request) {
	// sessionを使って、ログインしているか判定を取得
	_, err := session(w, r)
	// エラーがある場合、
	if err != nil {
		// トップページにリダイレクト
		http.Redirect(w, r, "/", 302)
	} else {
		// sessionが存在する場合、indexを表示
		generateHTML(w, nil, "layout", "private_navbar", "index")
	}
}
