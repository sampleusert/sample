package main

import (
	"time"
)

type User struct {
	Name     string
	Role     string
	HireDate time.Time
}

type Post struct {
	Id    string `datastore:"-" goon:"id"`
	Title string
	Body  string
}

type Title struct {
	Id      string `datastore:"-" goon:"id"`
	Name    string
	Propose string
	User    string
	Update  time.Time
}
