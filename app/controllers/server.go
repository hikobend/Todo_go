package controllers

import (
	"fmt"
	"hello/app/models"
	"hello/config"
	"html/template"
	"net/http"
)

// 関数を作成して共通化
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	// ファイル格納
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	// templatesを事前にキャッシュしておいて、効率的に処理
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

// cookieを取得する関数作成
func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	// HTTPリクエストからcookieを取得する
	cookie, err := r.Cookie("_cookie")
	// エラーがnil
	if err == nil {
		// sessionのstruct作成
		// ユーザーを作成してそのUUIDを保存
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			// 存在しない場合、エラーを生成
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

// サーバーの立ち上げコード作成
func StartMainServer() error {
	// CSS, jsファイル読み込み
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	// URLの登録
	// 第二引数はハンドラ
	http.HandleFunc("/", top)
	// URLに登録
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	// ポート作成
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
