package main

import (
	"encoding/json"
	"fmt"
	"json_db/config"
	"os"
	"strings"
)

func initialize() (string, error) {

	var dbName string = "database.json"

	if len(os.Args) > 1 {
		dbName = os.Args[1]

		if !strings.HasSuffix(dbName, ".json") {
			return "", fmt.Errorf("DB name must be a .json file, [%v] passed instead", dbName)
		}

		// check if db exists, create new if not
		dbPath := fmt.Sprintf("%v%v", config.DB_BASE_DIR, dbName)
		fileData, err := os.ReadFile(dbPath)
		if err != nil {
			fmt.Printf("DB [%v] does not exist, creating new...\n", dbName)

			os.Mkdir("database", 0o755)
			if err = os.WriteFile(dbPath, []byte(`{}`), 0o644); err != nil {
				return "", fmt.Errorf("failed creating DB [%v] - %v", dbName, err)
			}

			return dbName, nil
		}

		// check existing db for corruption
		jsonMap := map[string]any{}

		err = json.Unmarshal(fileData, &jsonMap)
		if err != nil {
			return "", fmt.Errorf("unable to parse JSON - %v; file is corrupt", err)
		}

	} else {
		return "", fmt.Errorf("database name is not passed as argument")
	}
	return dbName, nil
}
