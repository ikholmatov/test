package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Dbmethods struct {
}

func main() {
	connStr := "user=postgres password=1 dbname=hydorlife sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic("error while connecting to db")
	}
	defer db.Close()
	var (
		id      int64
		name    string
		email   string
		comment string
	)
	fmt.Println("1")
	fmt.Scan(&id)
	fmt.Println("2")
	fmt.Scan(&name)
	fmt.Println("3")
	fmt.Scan(&email)
	fmt.Println("4")
	fmt.Scan(&comment)

	dbs := Dbmethods{}

	usr, err := dbs.Create(Ecoreq{
		id:            id,
		name:          name,
		email_address: email,
		comment:       comment,
	})
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(usr)

}

func (h Dbmethods) Create(req Ecoreq) (Ecoreq, error) {
	user := &Ecoreq{}
	connStr := "user=postgres password=1 dbname=hydorlife sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("error")
		return Ecoreq{}, err
	}

	query := `INSERT INTO  eco (name,email_address,comments) values ($1,$2,$3) returning id,name,email_address,comments`
	err = db.QueryRow(query, req.name, req.email_address, req.comment).Scan(
		&user.id,
		&user.name,
		&user.email_address,
		&user.comment,
	)
	if err != nil {
		fmt.Println("error")
		return Ecoreq{}, err
	}
	return *user, nil

}

type Methods interface {
	Create(Ecoreq, db string) (Ecoreq, error)
}

type Ecoreq struct {
	id            int64
	name          string
	email_address string
	comment       string
}
