package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"log"
)

var MyBase = "user=venom password=112233 dbname=gorilla sslmode=disable"

type user struct {
	Id          string  `json:"id"`
	FistName    string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	Birthday    string  `json:"birthday"`
	PhoneNumber string  `json:"phoneNumber"`
	Address     address `json:"address"`
}
type address struct {
	Id         string `json:"id"`
	UserId     string `json:"userId"`
	Country    string `json:"country"`
	City       string `json:"city"`
	District   string `json:"district"`
	Apartment  string `json:"apartment"`
	PostalCode int64  `json:"postalCode"`
}

func (r *user) Create(info user) (bool, error) {

	db, err := sql.Open("postgres", MyBase)
	if err != nil {
		log.Panicf("%s\n%s", err, "Error while opening db")
	}
	UserID := uuid.NewV4()
	CreateQuery := `INSERT INTO users (id,first_name,last_name,birthday,phone_number) VALUES ($1,$2,$3,$4,$5)`
	_, err = db.Exec(CreateQuery, UserID, info.FistName, info.LastName, info.Birthday, info.PhoneNumber)
	if err != nil {
		return false, err
	}
	AddressID := uuid.NewV4()
	AddressQuery := `INSERT INTO addresses (id,user_id,country,city,district,apartment,postal_code) VALUES ($1,$2,$3,$4,$5,$6,$7)`
	Address := info.Address
	_, err = db.Exec(AddressQuery, AddressID, UserID, Address.Country, Address.City, Address.District, Address.Apartment, Address.PostalCode)
	if err != nil {
		return false, err
	}

	return true, err
}
func (r *user) Get(Id string) (user, error) {
	db, err := sql.Open("postgres", MyBase)
	if err != nil {
		log.Panicf("%s\n%s", err, "Error while opening db")
	}
	GetQuery := `SELECT first_name,last_name,birthday,phone_number FROM users WHERE id = $1`
	UserInf := user{}
	err = db.QueryRow(GetQuery, Id).Scan(&UserInf.FistName, &UserInf.LastName, &UserInf.Birthday, &UserInf.PhoneNumber)
	if err != nil {
		return UserInf, err
	}
	AddrGetQuery := `SELECT country, city, district, apartment, postal_code FROM addresses WHERE user_id = $1`
	err = db.QueryRow(AddrGetQuery, Id).Scan(&UserInf.Address.Country, &UserInf.Address.City, &UserInf.Address.District, &UserInf.Address.Apartment, &UserInf.Address.PostalCode)
	return UserInf, err
}
func (r *user) Update(Id string, num string) (bool, error) {
	db, err := sql.Open("postgres", MyBase)
	if err != nil {
		return false, err
	}
	UpdateQuery := `UPDATE users SET phone_number = $2 WHERE id = $1 `
	_, err = db.Exec(UpdateQuery, Id, num)
	if err != nil {
		return false, err
	}
	return true, err
}
func (r *user) Delete(Id string) (bool, error) {
	db, err := sql.Open("postgres", MyBase)
	if err != nil {
		return false, err
	}
	DeleteQuery := `DELETE FROM users WHERE id = $1 `
	_, err = db.Exec(DeleteQuery, Id)
	if err != nil {
		return false, err
	}

	return true, err

}

