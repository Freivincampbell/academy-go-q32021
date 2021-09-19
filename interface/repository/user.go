package repository

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type user struct {
}

type Data struct {
	Value string `json:"value"`
}

type User interface {
	ReadUsers(f string) (string, error)
	ReadUsersByKey(f string) (string, error)
}

func NewUserRepository() User {
	return &user{}
}

func (ur *user) ReadUsers(f string) (string, error) {
	csvFile, err := openFile(f)
	if err != nil {
		return "", err
	}

	err = validateCSV(csvFile)
	if err != nil {
		return "", err
	}

	return f, nil
}

func (ur *user) ReadUsersByKey(k string) (string, error) {
	csvFile, err := openFile("./public/data.csv")
	if err != nil {
		return "", err
	}

	defer csvFile.Close()

	csvData, err := readFile(csvFile)
	if err != nil {
		return "", err
	}

	jsonData, err := transformData(csvData, k)
	if err != nil {
		return "", err
	}

	f, err := writeFile(jsonData)
	if err != nil {
		return "", err
	}

	return f, nil
}

func openFile(f string) (*os.File, error) {
	csvFile, err := os.Open(f)
	if err != nil {
		err = fmt.Errorf("path provided was not found")
		return nil, err
	}
	return csvFile, nil
}

func validateCSV(csvFile *os.File) error {
	r := csv.NewReader(csvFile)
	for {
		_, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			err = fmt.Errorf("wrong number of fields or wrong format")
			return err
		}
	}

	return nil
}

func readFile(csvFile *os.File) (records [][]string, err error) {
	reader := csv.NewReader(csvFile)

	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()

	return csvData, err
}

func writeFile(jsonData []Data) (string, error) {
	var filePath = "./public/data_by_key.csv"
	csvFile, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)

	for _, temp := range jsonData {
		var row []string
		row = append(row, temp.Value)
		writer.Write(row)
	}

	writer.Flush()

	return filePath, nil
}

func transformData(csvData [][]string, k string) ([]Data, error) {

	var oneRecord Data
	var allRecords []Data

	for i, _ := range csvData[0] {
		if strings.ToLower(strings.TrimSpace(csvData[0][i])) == strings.ToLower(k) {
			for _, e := range csvData {
				oneRecord.Value = string(e[i])
				allRecords = append(allRecords, oneRecord)
			}
		}
	}

	r, err := json.Marshal(allRecords)
	var jsonData []Data
	err = json.Unmarshal(r, &jsonData)

	return jsonData, err
}
