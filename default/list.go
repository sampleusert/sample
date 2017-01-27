package main

import (
	"github.com/zenazn/goji/web"
	"net/http"
)

func list(c web.C, w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Name": "home",
	}
	render("views/list.html", w, data)
}
