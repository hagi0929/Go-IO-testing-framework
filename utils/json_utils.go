package utils

import (
	"encoding/json"
	"fmt"
)

func PrettyPrintJSON(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}
