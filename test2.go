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
	price  float64
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

	_, err = db.Exec(openquery)
	if err != nil {
		log.Panicf("%s\n%s", err, "Error while creating table")
	}

	defer func(db *sql.DB) {
		if err := db.Close(); err != nil {
			log.Panicf("%s\n%s", err, "Error while closeing db")
		}
	}(db)
	insertquery := `INSERT INTO book (title,author,price,amount) VALUES ('Белая гвардия','Булгаков М.А.',540.50,5), ('Идиот',
	'Достоевский Ф.М.',460.00,10), ('Братья Карамазовы','Достоевский Ф.М.',799.01,2); `

	_, err = db.Exec(insertquery)
	if err != nil {
		log.Panicf("%s\n%s", err, "Error while inserting values to the table")
	}

	myquery := `SELECT * FROM book WHERE amount > 3;`

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
		err := rows.Scan(&get.id, &get.title, &get.author, &get.price, &get.amount)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, get)
	}
	for _, v := range products {
		fmt.Println(v)
	}
}
