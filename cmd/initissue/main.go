package main

import (
	"fmt"

	"github.com/marcpires/rss/internal/db"
	redis "github.com/marcpires/rss/internal/db/redis"
)

var f = func() int {
	fmt.Println("A function")
	return 0
}()

func init() {
	fmt.Println("init function from main.go")
}

func init() {
	fmt.Println("init function 2 from main.go")
}

func main() {
	fmt.Println("main function")
	err := db.Connect("mysql://user:password@tcp(localhost:3306)/dbname")

	// Do something with the error, otherwise compiler will complain about it being unused
	if err != nil {
		fmt.Println("Error connecting to MySQL:", err)
		return
	}

	err = redis.Connect("redis://localhost:6379/0")
}
