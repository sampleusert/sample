package main

import (
	"encoding/json"
	"github.com/zenazn/goji/web"
	"google.golang.org/appengine/log"
	"net/http"
)

/*type Status struct {
	code string
	info []string
}*/
type Status struct {
	Id      string
	Balance string
}

func title(c web.C, w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Name": "home",
	}
	render("views/title.html", w, data)
}

func titleCreate(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r)
	log.Infof(ctx, "Index")
	//fmt.Fprintf(w, "okok")

	/*status := Status{"Alex", []string{"snowboarding", "programming"}}
	js, err := json.Marshal(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)*/

	name := r.FormValue("input_text")
	info := r.FormValue("textarea1")
	u := Status{Id: name, Balance: info}
	json.NewEncoder(w).Encode(u)
}
