package operations

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadAll(dbPath string) (map[string]any, error) {

	fileData, err := os.ReadFile(dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed reading db file [%v] - %v", dbPath, err)
	}

	var data map[string]any
	if err = json.Unmarshal(fileData, &data); err != nil {
		return nil, fmt.Errorf("failed parsing db file [%v] - %v", dbPath, err)
	}

	return data, nil
}
