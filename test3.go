package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type data interface {
	Update()
	Get(string) ([]info, error)
	Delete(int, string) (string, error)
}
type info struct {
	id                int64
	jobtitle          string
	emailaddress      string
	firstnamelastname string
	phonenumber       string
	country           string
}

func (info) Get(db string) ([]info, error) {

	ss, err := sql.Open("postgres", db)
	if err != nil {
		log.Panicf("%s\n%s", err, "Error while opening db")
	}
	defer ss.Close()
	myqueary := `SELECT firstnamelastname,country,emailaddress,jobtitle FROM mockdata ORDER BY firstnamelastname ASC; `
	rows, err := ss.Query(myqueary)
	if err != nil {
		log.Panicf("%s\n%s", err, "Error while get info from db")
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			log.Panicf("%s\n%s", err, "Error while closing rows from db")
		}
	}(rows)
	var inf []info
	for rows.Next() {
		get := info{}
		err := rows.Scan(&get.firstnamelastname, &get.country, &get.emailaddress, &get.jobtitle)
		if err != nil {
			fmt.Println(err)
			continue
		}
		inf = append(inf, get)
	}
	return inf, err
}

func main() {
	mybase := "user=venom password=112233 dbname=venom sslmode=disable"
	var w info
	get, err := w.Get(mybase)
	if err != nil {
		panic(err)
	}
	for _, v := range get {
		fmt.Println(v.firstnamelastname)
	}

}
func (info) Delete(id int, db string) (string, error) {
	ss, err := sql.Open("postgres", db)
	if err != nil {
		log.Panicf("%s\n%s", err, "Error while opening db")
	}
	defer ss.Close()
	myquery := (`DELETE FROM mockdata where id $1`)
}
