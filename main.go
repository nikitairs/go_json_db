package main

import (
	"fmt"
)

func main() {
	dbName := initialize()
	fmt.Printf("Initialized JSON DB connection - %v\n", dbName)

}
