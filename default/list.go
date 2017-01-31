package main

import (
	"github.com/mjibson/goon"
	"github.com/zenazn/goji/web"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"html/template"
	"net/http"
	"strconv"
)

func list(c web.C, w http.ResponseWriter, r *http.Request) {
	/*data := map[string]interface{}{
		"Name": "home",
	}*/
	//titles := []Title{}
	titles := make([]Title, 0)
	//titlesViews := make([]Title, 2)
	cc := appengine.NewContext(r)
	q := datastore.NewQuery("Title")
	count, err := q.Count(cc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	titlesViews := make([]Title, count)

	g := goon.NewGoon(r)
	g.GetAll(q, &titles)

	ctx := newContext(r)

	for pos, title := range titles {
		log.Infof(ctx, "test:"+title.Name)
		log.Infof(ctx, "count:"+strconv.Itoa(pos))

		//member_views[pos].Id = keys[pos].IntID()
		//member_views[pos].Name = fmt.Sprintf("%s", member.Name)
		//member_views[pos].Gender = member.Gender
		titlesViews[pos].Id = title.Id
		titlesViews[pos].Name = title.Name
		titlesViews[pos].Propose = title.Propose
		titlesViews[pos].User = title.User
		log.Infof(ctx, "testok:")
	}
	log.Infof(ctx, "testok1:")

	/*data := map[string]interface{}{
		"Name":   "users/index",
		"Titles": titles,
	}*/

	//render("views/list.html", w, titlesViews)
	tmpl := template.Must(template.ParseFiles("views/layout.html", "views/list.html"))
	tmpl.Execute(w, titlesViews)

}
