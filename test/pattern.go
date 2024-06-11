package test

import (
	"your_project_path/models"
)

type PatternType struct {
	InputPattern  string `json:"input-pattern"`
	OutputPattern string `json:"output-pattern"`
}

func (t PatternType) GetTestFiles(path string) []models.IOFile {
	return []models.IOFile{{Input: t.InputPattern, Output: t.OutputPattern}}
}
