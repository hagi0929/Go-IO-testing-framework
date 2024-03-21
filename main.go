package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type IOFile struct {
	input  string `json:"input"`
	output string `json:"output"`
}

type ManualPairType struct {
	TestFormat
	IOPairs []IOFile `json:"io-pairs"`
}

func (t ManualPairType) GetTestFiles(path string) []IOFile {
	return t.IOPairs
}

type PatternType struct {
	TestFormat
	inputPattern  string `json:"input-pattern"`
	outputPattern string `json:"output-pattern"`
}

func (t PatternType) GetTestFiles(path string) []IOFile {
	return []IOFile{{t.inputPattern, t.outputPattern}}
}

type TestFormat interface {
	GetTestFiles(path string) []IOFile
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
		var manualPair ManualPairType
		err = json.Unmarshal(testFormatData, &manualPair)
		config.TestFormat = manualPair
	case "manual-pair":
		var pattern PatternType
		err = json.Unmarshal(testFormatData, &pattern)
		config.TestFormat = pattern
	default:
		fmt.Println("Invalid test format")
	}
	return *config
}

func (tc TestConfig) Print() {
	fmt.Println(tc)
}

func main() {
	fmt.Println("Hello, world.")
	configObj := NewTestConfig("./setup.json")
	configObj.Print()
}
