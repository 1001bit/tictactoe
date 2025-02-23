package handler

import "net/http"

func HandleHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./html/home.html")
}
