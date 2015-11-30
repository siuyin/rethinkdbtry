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
	fmt.Println(sess)
}
