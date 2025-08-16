package main

import (
	"fmt"
)

func main() {
	dbName, err := initialize()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Initialized JSON DB connection - %v\n", dbName)

}
