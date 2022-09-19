package controllers

import (
	"hello/app/models"
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
		//sessionを取り出す
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		// todosをByUserから取り出す
		todos, _ := user.GetTodoByUser()
		user.Todos = todos
		// userの情報を渡す
		// sessionが存在する場合、indexを表示
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}

// todo作成関数
func todoNew(w http.ResponseWriter, r *http.Request) {
	// ログイン確認
	_, err := session(w, r)
	// エラーハンドリング
	if err != nil {
		// ログインページログイン
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}
}

// todo更新関数
func todoSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		// 入力フォームの値を取得する
		err = r.ParseForm()
		// ハンドリング
		if err != nil {
			log.Println(err)
		}
		// ユーザーの取得
		user, err := sess.GetUserBySession()
		// エラーハンドリング
		if err != nil {
			log.Println(err)
		}
		// フォームの値を取得
		content := r.PostFormValue("content")
		// 取得したコンテントを渡す
		// エラーハンドリング
		if err := user.CreateTodo(content); err != nil {
			log.Println(err)
		}
		// todos一覧画面に遷移
		http.Redirect(w, r, "/todos", 302)
	}
}

func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		t, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}
		generateHTML(w, t, "layout", "private_navbar", "todo_edit")
	}
}

func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		log.Println(err)
	}
	user, err := sess.GetUserBySession()
	if err != nil {
		log.Println(err)
	}
	content := r.PostFormValue("content")
	t := &models.Todo{ID: id, Content: content, UserID: user.ID}
	if err := t.UpdateTodo(); err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/todos", 302)
}

func todoDelete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		t, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}
		if err := t.DeleteTodo(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", 302)
	}
}
