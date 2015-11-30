package main

import (
	"fmt"
	"log"
	"os"

	r "github.com/dancannon/gorethink"
)

func main() {
	fmt.Println("Connecting to rethinkdb", os.Getenv("RETHINKDB_ADDR"), os.Getenv("RETHINKDB_DB"))
	sess, err := r.Connect(r.ConnectOpts{
		Address:  os.Getenv("RETHINKDB_ADDR"),
		Database: os.Getenv("RETHINKDB_DB")})
	if err != nil {
		log.Fatal(err)
	}

	dat := map[string]interface{}{
		"id":    "124",
		"Name":  "David Davidson",
		"Place": "Somewhere",
	}
	insert(sess, dat)
}

func insert(sess *r.Session, dat interface{}) {
	res, err := r.Table("people").Insert(dat).RunWrite(sess)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
