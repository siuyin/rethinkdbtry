package main

import (
	"encoding/json"
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

	res, err := r.DB(os.Getenv("RETHINKDB_DB")).TableCreate("people").RunWrite(sess)
	if err != nil {
		log.Fatal(err)
	}
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}
