package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

const (
	csvPath  = "./data.csv"
	jsonPath = "data.json"
)

func WriteDataToFileAsJSON(data interface{}, filedir string) error {
	//write data as buffer to json encoder
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "\t")

	err := encoder.Encode(data)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filedir, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	_, err = file.Write(buffer.Bytes())
	if err != nil {
		return err
	}

	fmt.Println("Create and write file successfully !")
	return nil
}

func ReadCsv(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Unable to read input file %s %e\n", filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Unable to parse file as CSV for %s %e\n", filePath, err)
	}

	return records, nil
}
