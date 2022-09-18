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

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
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
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/todos", index)
	// ポート作成
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
