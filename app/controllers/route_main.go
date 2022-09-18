package controllers

import (
	"net/http"
)

// func top(w http.ResponseWriter, r *http.Request) {
// layout topを実行
func top(w http.ResponseWriter, _ *http.Request) {
	generateHTML(w, "Hello", "layout", "top")
}
