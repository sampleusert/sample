package main

import (
	//"fmt"
	"github.com/mjibson/goon"
	"github.com/zenazn/goji/web"
	//"google.golang.org/appengine"
	//"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"html/template"
	"net/http"
	//"strconv"
)

func info(c web.C, w http.ResponseWriter, r *http.Request) {
	n := goon.NewGoon(r)
	g := &Title{Id: r.FormValue("id")}
	//err := n.Get(g)
	n.Get(g)

	ctx := newContext(r)
	log.Infof(ctx, g.Name)

	//render("views/info.html", w, titlesViews)
	tmpl := template.Must(template.ParseFiles("views/layout.html", "views/info.html"))
	tmpl.Execute(w, g)
	//fmt.Fprintf(w, g.Name)

}

func comment(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r)
	log.Infof(ctx, "Index")
	//fmt.Fprintf(w, "okok")

	//シーケンスにしたい

	//g := goon.NewGoon(r)
	//title := Title{Id: r.FormValue("input_text"), Name: r.FormValue("input_text"), Propose: r.FormValue("textarea1"), User: "test", Update: time.Now()}
	//post := Post{Title: "タイトル", Body: "本文です..."}

	//g.Put(&post)
	/*if _, err := g.Put(&title); err != nil {
		u := Status{Id: "ng", Balance: "ng"}
		json.NewEncoder(w).Encode(u)
		return
	}*/

	name := r.FormValue("titleId")
	info := r.FormValue("comment")
	u := Status{Id: name, Balance: info}
	json.NewEncoder(w).Encode(u)
}
