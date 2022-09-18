package controllers

import (
	"fmt"
	"hello/config"
	"html/template"
	"net/http"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

// サーバーの立ち上げコード作成
func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	// URLの登録
	// 第二引数はハンドラ
	http.HandleFunc("/", top)
	// ポート作成
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
