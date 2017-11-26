package web

import (
	"net/http"

	"web/box"
	"web/timeline"
)

//Start http://localhost:4444/box
func Start() {
	http.Handle("/", http.FileServer(http.Dir("../res/public")))
	http.HandleFunc("/pic/", func(w http.ResponseWriter, r *http.Request) {
		p := box.DataPath + r.URL.Path[4:]
		http.ServeFile(w, r, p)
	})
	box.RegisterPath("/box")
	timeline.RegisterPath("/tl")
	http.ListenAndServe(":4444", nil)
}
