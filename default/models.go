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