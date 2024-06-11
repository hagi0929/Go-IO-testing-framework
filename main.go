package main

import (
	"fmt"
	"your_project_path/config"
)

func main() {
	fmt.Println("Hello, world.")
	configObj := config.NewTestConfig("./setup.json")
	configObj.Print()
}
