package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.OutputFilePath)

	if err != nil {
		return nil, errors.New("ERROR: Fail to open File")
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("ERROR: Fail to reading File content")
	}

	file.Close()
	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.InputFilePath)

	if err != nil {
		return errors.New("ERROR: Fail to create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("ERROR: Fail to convert data to JSON")
	}

	file.Close()
	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
