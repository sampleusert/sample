package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"golang.org/x/net/context"

	//"appengine"
	"google.golang.org/appengine/user"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"html/template"
	"io"

	"google.golang.org/appengine/datastore"
	"time"

	"github.com/mjibson/goon"
)

func init() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/welcome", welcome)

	http.Handle("/html", goji.DefaultMux)
	goji.Get("/html", homeHandler)

	http.Handle("/datastore", goji.DefaultMux)
	goji.Get("/datastore", homeHandler2)

	http.Handle("/dataread", goji.DefaultMux)
	goji.Get("/dataread", usersIndexHandler)

	http.Handle("/goon", goji.DefaultMux)
	goji.Get("/goon", PutWorktime)

}

func PutWorktime(c web.C, w http.ResponseWriter, r *http.Request) {
	g := goon.NewGoon(r)
	post := Post{Id: "abc", Title: "タイトル", Body: "本文です"}
	//post := Post{Title: "タイトル", Body: "本文です..."}
	//g.Put(&post)
	if _, err := g.Put(&post); err != nil {
		//ctx := newContext(r)
		//log.Infof(ctx, err)
		fmt.Fprintf(w, "NG")
	}
	fmt.Fprintf(w, "goon")
}

func usersIndexHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	context := appengine.NewContext(r)
	q := datastore.NewQuery("users").Limit(10)

	users := make(map[string]User, 0)

	for t := q.Run(context); ; {
		var user User
		key, err := t.Next(&user)
		if err == datastore.Done {
			break
		}
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		users[key.String()] = user
	}

	data := map[string]interface{}{
		"Name":  "users/index",
		"Users": users,
	}
	render("views/users/index.html", w, data)
}

func homeHandler2(c web.C, w http.ResponseWriter, r *http.Request) {

	context := appengine.NewContext(r)

	el := User{
		Name:     "Joe",
		Role:     "Manager",
		HireDate: time.Now(),
	}

	key, err := datastore.Put(context, datastore.NewIncompleteKey(context, "users", nil), &el)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, key.String())

	// ...
}

func render(v string, w io.Writer, data map[string]interface{}) {
	tmpl := template.Must(template.ParseFiles("views/layout.html", v))
	tmpl.Execute(w, data)
}

func homeHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Name": "home",
	}
	render("views/home.html", w, data)
}

func handle(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r)
	log.Infof(ctx, "Index")
	fmt.Fprintf(w, "ok")
}

func newContext(r *http.Request) context.Context {
	return appengine.NewContext(r)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	// Contextの生成
	ctx := appengine.NewContext(r)
	// [2]
	u := user.Current(ctx)
	// [3]
	if u == nil {
		// [4]
		loginUrl, _ := user.LoginURL(ctx, "/welcome")
		fmt.Fprintf(w, `<a href="%s">Sign in</a>`, loginUrl)
		return
	}
	// [5]
	logoutUrl, _ := user.LogoutURL(ctx, "/welcome")
	fmt.Fprintf(w, `Welcome, %s! (<a href="%s">Sign out</a>)`, u, logoutUrl)
}
