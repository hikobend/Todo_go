package controllers

import (
	"net/http"
)

// func top(w http.ResponseWriter, r *http.Request) {
// layout topを実行
func top(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "Hello", "layout", "public_navbar", "top")
}
