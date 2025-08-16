package main

import (
	"fmt"
	"json_db/config"
	"json_db/operations"
)

func main() {
	dbName, err := initialize()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Initialized JSON DB - %v\n", dbName)

	// read operation
	dbPath := fmt.Sprintf("%v%v", config.DB_BASE_DIR, dbName)
	allData, err := operations.ReadAll(dbPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(allData)

}
