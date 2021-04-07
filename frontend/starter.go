package main

import (
	"net/http"

	"crawler/frontend/controller"
)

func main() {
	// w := view.CreateSearchResultView("view/template.html")
	http.Handle("/", http.FileServer(http.Dir("frontend/view")))
	http.Handle("/search", controller.CreateSearchResultHandler("frontend/view/template.html"))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
