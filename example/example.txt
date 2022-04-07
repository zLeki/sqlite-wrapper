package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	data "liquidgen/sql"
	"log"
	"time"
)

func main() {
	x := time.Now()
	db, _ := sql.Open("sqlite3", "./database.db")
	feed := data.NewTable(db, "no")
	log.Println(feed.Query())
	feed.Edit(data.Item{
		ID: 2,
	}, data.Item{
		Content: "Bye World",
	})
	log.Println(feed.Query())

	log.Println(time.Since(x))

}
