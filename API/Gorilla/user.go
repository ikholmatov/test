package main

import "time"

type user struct {
	id          string
	firstName   string
	lastName    string
	birthday    time.Time
	phoneNumber string
	address     addresses
}
type addresses struct {
	id         string
	userId     string
	country    string
	city       string
	district   string
	apartment  string
	postalCode int64
}
