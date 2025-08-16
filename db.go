package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const emptyDb string = `{}`

func initialize() string {

	var dbPath string = "database/database.json"

	if len(os.Args) > 1 {
		dbName := os.Args[1]

		if !strings.HasSuffix(dbName, ".json") {
			log.Fatalf("DB name must be a .json file, [%v] passed instead\n", dbName)
		}

		// check if file exists
		dbPath = fmt.Sprintf("database/%v", dbName)
		if _, err := os.Open(dbPath); err != nil {
			fmt.Printf("DB [%v] does not exist, creating new\n", dbPath)

			err = os.WriteFile(dbPath, []byte(emptyDb), 0o644)
			if err != nil {
				log.Fatalf("Failed creating DB [%v] - %v\n", dbPath, err)
			}
		}
	} else {
		log.Fatalln("Database name is not passed as argument")
	}

	return dbPath
}
