package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func initialize() string {

	var dbName string = "database.json"

	if len(os.Args) > 1 {
		dbName = os.Args[1]

		if !strings.HasSuffix(dbName, ".json") {
			log.Fatalf("DB name must be a .json file, [%v] passed instead\n", dbName)
		}

		// check if db exists, create new if not
		dbPath := fmt.Sprintf("database/%v", dbName)
		file, err := os.Open(dbPath)
		if err != nil {
			fmt.Printf("DB [%v] does not exist, creating new...\n", dbPath)

			os.Mkdir("database", 0o755)
			if err = os.WriteFile(dbPath, []byte(`{}`), 0o644); err != nil {
				log.Fatalf("Failed creating DB [%v] - %v\n", dbPath, err)
			}
		}

		// check existing db for corruption
		var sb strings.Builder

		_, err = io.Copy(&sb, file)
		if err != nil {
			log.Fatalf("unable to retrieve data from file [%v]; file is corrupt!\n", dbName)
		}

		fileString := sb.String()
		jsonMap := map[string]any{}

		err = json.Unmarshal([]byte(fileString), &jsonMap)
		if err != nil {
			log.Fatalf("Unable to parse JSON - %v; file is corrupt\n", err)
		}

	} else {
		log.Fatalln("Database name is not passed as argument")
	}
	return dbName
}
