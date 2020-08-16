package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func readFile(path string) (inputData, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return inputData{}, err
	}

	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return inputData{}, err
	}

	var input inputData

	err = json.Unmarshal(b, &input)
	if err != nil {
		return inputData{}, err
	}

	return input, nil
}

func writeFile(output outputData, path string) error {
	b, err := json.Marshal(output)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path, b, 0644); err != nil {
		return err
	}

	return nil
}
