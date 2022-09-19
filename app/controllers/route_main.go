package controllers

import (
	"log"
	"net/http"
)

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
	sess, err := session(w, r)
	// エラーがある場合、
	if err != nil {
		// トップページにリダイレクト
		http.Redirect(w, r, "/", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		todos, _ := user.GetTodoByUser()
		user.Todos = todos
		// sessionが存在する場合、indexを表示
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}

func todoNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}
}

func todoSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")
		if err := user.CreateTodo(content); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", 302)
	}
}
