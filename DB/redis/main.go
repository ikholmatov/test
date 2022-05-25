package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
)

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	err := client.Set("name", "erlll", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := client.Get("name").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(val)
	asd, err := json.Marshal(Author{Name: "Elliot", Age: 25})
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set("id1234", asd, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err = client.Get("id1234").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}
