package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"your_project_path/models"
	"your_project_path/test"
	"your_project_path/utils"
)

type TestFormat interface {
	GetTestFiles(path string) []models.IOFile
}

type TestConfig struct {
	SourcePath   string     `json:"source-folder"`
	TestCasePath string     `json:"test-folder"`
	TestFormat   TestFormat `json:"-"`
}

func NewTestConfig(configPath string) TestConfig {
	jsonFile, err := os.Open(configPath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	var result map[string]interface{}
	json.Unmarshal(byteValue, &result)

	testFormatType := result["test-format-type"].(string)

	testFormatData, err := json.Marshal(result["test-format"])
	if err != nil {
		fmt.Println(err)
	}

	config := &TestConfig{
		SourcePath:   result["source-folder"].(string),
		TestCasePath: result["test-folder"].(string),
	}

	switch testFormatType {
	case "pattern":
		var pattern test.PatternType
		err = json.Unmarshal(testFormatData, &pattern)
		config.TestFormat = pattern
	case "manual-pair":
		var manualPair test.ManualPairType
		err = json.Unmarshal(testFormatData, &manualPair)
		config.TestFormat = manualPair
	default:
		fmt.Println("Invalid test format")
	}
	return *config
}

func (tc TestConfig) Print() {
	fmt.Println(tc)
}
