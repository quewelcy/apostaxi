package web

import (
	"net/http"

	"web/box"
	"web/timeline"
)

//Start http://localhost:4444/box
func Start() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../res/public"))))
	http.HandleFunc("/pic/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.FormValue("p"))
	})
	box.RegisterPath()
	timeline.RegisterPath("/tl")
	http.ListenAndServe(":4444", nil)
}
