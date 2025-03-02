package handler

import "net/http"

func HandleRoom(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./html/room.html")
}
