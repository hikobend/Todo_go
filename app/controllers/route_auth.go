package controllers

import (
	"hello/app/models"
	"log"
	"net/http"
)

// サインアップ関数
func signup(w http.ResponseWriter, r *http.Request) {
	// GETのときの処理
	if r.Method == "GET" {
		// エラーが存在したら
		_, err := session(w, r)
		if err != nil {
			// signup.htmlテンプレートファイルだけを出力するようにする
			generateHTML(w, nil, "layout", "public_navbar", "signup")
			// sessionが存在したら
		} else {
			// todosにRediret
			http.Redirect(w, r, "/todos", 302)
		}
		// POSTのときの処理
	} else if r.Method == "POST" {
		// 入力フォームの値を取得する
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		// Userを作成
		user := models.User{
			// Nameを入力フォームから受け取る
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		// user作成メソッドを使用
		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}
		// ユーザーの登録が成功したら、トップページにリダイレクトさせる
		http.Redirect(w, r, "/", 302)
	}
}

// 関数作成
func login(w http.ResponseWriter, r *http.Request) {
	// エラーが存在したら
	_, err := session(w, r)
	if err != nil {
		// 表示するtemplateを設定
		generateHTML(w, nil, "layout", "public_navbar", "login")
		// sessioonが存在したら
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

// ユーザーの認証認証
func authenticate(w http.ResponseWriter, r *http.Request) {
	// フォームから値を取得
	err := r.ParseForm()
	// フォームからメールを取得
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Fatalln(err)
		// ログインページにリダイレクト
		http.Redirect(w, r, "/login", 302)
	}
	// パスワードが一致しているか
	if user.PassWord == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}
		// cookie作成
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)

		// ログインに成功
		// リダイレクト
		http.Redirect(w, r, "/", 302)
	} else {
		// パスワード不一致
		http.Redirect(w, r, "/login", 302)
	}
}

// ログアウト関数作成
func logout(w http.ResponseWriter, r *http.Request) {
	// cookie取得
	cookie, err := r.Cookie("_cookie")

	// エラーハンドリング
	if err != nil {
		log.Println(err)
	}

	// エラーがErrNoCookieと異なる場合
	if err != http.ErrNoCookie {
		// session struct作成
		session := models.Session{UUID: cookie.Value}
		// UUIDと一致するものを削除
		session.DeleteSessionByUUID()
	}
	// 削除したらRedirectする
	http.Redirect(w, r, "/login", 302)
}
