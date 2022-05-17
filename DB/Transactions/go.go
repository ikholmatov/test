package Transactions

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type info struct {
	id             int
	FullName       string
	BankcardNumber string
	AmountInCard   string
}

func main() {
	mybase := "user=venom password=112233 dbname=venom sslmode=disable"
}

func (info) Update(mybase string, data info) (string, error) {
	db, err := sql.Open("postgres", mybase)
	if err != nil {
		panic("err while opening data base")
	}

	tx, err := db.Begin()
	if err != nil {
		panic("Err while beginning Tx")
	}
	myquery := ""
}
