package test

import (
	"your_project_path/models"
)

type ManualPairType struct {
	IOPairs []models.IOFile `json:"io-pairs"`
}

func (t ManualPairType) GetTestFiles(path string) []models.IOFile {
	return t.IOPairs
}
