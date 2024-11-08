package main

import (
	"net/http"

	web "github.com/yeungon/corpora/cmd"
)

func main() {
	http.HandleFunc("/", web.Index)
	http.HandleFunc("/dashboard", web.Dashboard)
	http.HandleFunc("/profile/show", web.ProfileShow)
	http.HandleFunc("/profile/edit", web.ProfileEdit)
	http.ListenAndServe(":9999", nil)
}
