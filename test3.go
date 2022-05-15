package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	//mybase := "user=venom password=112233 dbname=venom sslmode=disable"
	//var w info
	//Code for use get
	//get, err := w.Get(33, mybase)
	//if err != nil {
	//	panic(err)
	//}
	//for _, v := range get {
	//	fmt.Println(v.id, v.firstnamelastname, v.jobtitle, v.emailaddress, v.country, v.phonenumber)
	//}
	//Code for use delete
	//s, err := w.Delete(99, mybase)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(s)
	//Code for use Upadta
	//update, err := w.Update(33, "Devops Engineer", mybase)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(update)
}

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

func (info) Get(id int, db string) ([]info, error) {

	ss, err := sql.Open("postgres", db)
	if err != nil {
		log.Panicf("%s\n%s", err, "Error while opening db")
	}
	defer ss.Close()
	myqueary := `SELECT phonenumber,id,firstnamelastname,country,emailaddress,jobtitle FROM mockdata WHERE id = $1; `
	rows, err := ss.Query(myqueary, id)
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
		err := rows.Scan(&get.phonenumber, &get.id, &get.firstnamelastname, &get.country, &get.emailaddress, &get.jobtitle)
		if err != nil {
			fmt.Println(err)
			continue
		}
		inf = append(inf, get)
	}
	return inf, err
}

func (info) Delete(id int, db string) (string, error) {
	ss, err := sql.Open("postgres", db)
	if err != nil {
		log.Panicf("%s\n%s", err, "Error while opening db")
	}
	defer ss.Close()
	myquery := `DELETE FROM mockdata where id = $1`
	_, err = ss.Exec(myquery, id)
	if err != nil {
		log.Panicf("%s\n%s", err, "Error while deleting info from table")
	}
	return "Ok", err
}
func (info) Update(id int, nam string, db string) (string, error) {
	ss, err := sql.Open("postgres", db)
	if err != nil {
		log.Panicf("%s\n%s", err, "Error while opening db")
	}
	defer ss.Close()
	myquery := `UPDATE mockdata SET jobtitle = $2 WHERE id = $1`
	_, err = ss.Exec(myquery, id, nam)
	if err != nil {
		panic(err)
	}
	return "Ok", err
}
