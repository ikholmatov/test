package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type product struct {
	id     int
	title  string
	author string
	price  string
	amount int
}

func main() {

	mybase := "user=venom password=112233 dbname=venom sslmode=disable"
	db, err := sql.Open("postgres", mybase)
	if err != nil {
		log.Panicf("%s\n%s", err, "Error while opening db")
	}
	openquery := `CREATE TABLE IF NOT EXISTS book(book_id SERIAL PRIMARY KEY ,title 
				VARCHAR(50),author VARCHAR(50),price DECIMAL(8, 2),amount INT);`

	ans, err := db.Exec(openquery)
	if err != nil {
		log.Panicf("%s\n%s", err, "Error while creating table")
	}
	fmt.Println(ans.RowsAffected())

	defer func(db *sql.DB) {
		if err := db.Close(); err != nil {
			log.Panicf("%s\n%s", err, "Error while closeing db")
		}
	}(db)
	myquery := `SELECT author,title,
       			DO $$
       			BEGIN
       			IF author = "Булгаков М.А."
       				RAISE PRICES price * 1.10
       			END IF;
       			END $$;
           as new_price from book;`

	rows, err := db.Query(myquery)
	if err != nil {
		log.Panicf("%s\n%s", err, "Error while geting row from db")
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			log.Panicf("%s\n%s", err, "Error while closing rows from db")
		}
	}(rows)
	var products []product
	for rows.Next() {
		get := product{}
		err := rows.Scan(&get.author, &get.title, &get.price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, get)
	}
}
