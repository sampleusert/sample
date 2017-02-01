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
	"encoding/json"
	"google.golang.org/appengine/memcache"
	//"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"time"
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

	g := goon.NewGoon(r)
	comment := Comment{Id: time.Now().String(), TitleId: r.FormValue("titleId"), Comment: r.FormValue("comment"), User: "test", Update: time.Now()}
	//post := Post{Title: "タイトル", Body: "本文です..."}

	//g.Put(&post)
	if _, err := g.Put(&comment); err != nil {
		u := Status{Id: "ng", Balance: "ng"}
		json.NewEncoder(w).Encode(u)
		return
	}

	name := r.FormValue("titleId")
	info := r.FormValue("comment")
	u := Status{Id: name, Balance: info}
	json.NewEncoder(w).Encode(u)
}

func commentList(w http.ResponseWriter, r *http.Request) {

	//comments := make([]Comment, 0)

	//gneral
	/*q := datastore.NewQuery("Comment")
	g := goon.NewGoon(r)
	g.GetAll(q, &comments)*/
	/*var s CommentList
	for _, comment := range comments {
		//commentViews[pos].Id = comment.Id
		//commentViews[pos].Comment = comment.Comment
		s.Comment = append(s.Comment, comment)
	}*/

	ctx := newContext(r)
	q := datastore.NewQuery("Comment").Limit(2)
	g := goon.NewGoon(r)

	// If the application stored a cursor during a previous request, use it.
	item, err := memcache.Get(ctx, "person_cursor")
	if err == nil {
		cursor, err := datastore.DecodeCursor(string(item.Value))
		if err == nil {
			q = q.Start(cursor)
		}
	}

	// Iterate over the results.
	t := g.Run(q)
	var s CommentList
	for {
		var p Comment
		_, err := t.Next(&p)
		if err == datastore.Done {
			break
		}
		s.Comment = append(s.Comment, p)
		if err != nil {
			log.Errorf(ctx, "fetching next Person: %v", err)
			break
		}
		// Do something with the Person p
	}

	// Get updated cursor and store it for next time.
	if cursor, err := t.Cursor(); err == nil {
		memcache.Set(ctx, &memcache.Item{
			Key:   "person_cursor",
			Value: []byte(cursor.String()),
		})
	}

	json.NewEncoder(w).Encode(s)
	//fmt.Println(string(b))

	//var s Serverslice
	//str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	//json.Unmarshal([]byte(str), &s)
	//fmt.Println(s)
	//json.NewEncoder(w).Encode(s)

	/*var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		//fmt.Println("json err:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(string(b))*/
}