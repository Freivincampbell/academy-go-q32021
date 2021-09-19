package repository

import (
	"academy-go-q32021/domain/model"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type user struct {
}

type User interface {
	ReadUsers(f string) ([]*model.User, error)
	ReadUsersByKey(f string) ([]*model.CustomCSV, error)
}

func NewUserRepository() User {
	return &user{}
}

func (ur *user) ReadUsers(f string) ([]*model.User, error) {
	csvFile, err := openFile(f)
	if err != nil {
		return nil, err
	}

	defer csvFile.Close()

	data, err := readFile(csvFile)
	if err != nil {
		return nil, err
	}

	jsonData, err := transformData(data)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func (ur *user) ReadUsersByKey(k string) ([]*model.CustomCSV, error) {
	csvFile, err := openFile("./public/data.csv")
	if err != nil {
		return nil, err
	}

	defer csvFile.Close()

	csvData, err := readFile(csvFile)
	if err != nil {
		return nil, err
	}

	jsonData, err := transformDataByKey(csvData, k)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func openFile(f string) (*os.File, error) {
	csvFile, err := os.Open(f)
	if err != nil {
		err = fmt.Errorf("path provided was not found")
		return nil, err
	}
	return csvFile, nil
}

func readFile(csvFile *os.File) (records [][]string, err error) {
	reader := csv.NewReader(csvFile)

	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()

	return csvData, err
}

func transformDataByKey(csvData [][]string, k string) ([]*model.CustomCSV, error) {
	var oneRecord model.CustomCSV
	var allRecords []model.CustomCSV

	for i, _ := range csvData[0] {
		if strings.ToLower(strings.TrimSpace(csvData[0][i])) == strings.ToLower(k) {
			for j, e := range csvData {
				if j != 0 {
					oneRecord.Value = strings.TrimSpace(string(e[i]))
					allRecords = append(allRecords, oneRecord)
				}
			}
		}
	}

	r, err := json.Marshal(allRecords)
	var jsonData []*model.CustomCSV
	err = json.Unmarshal(r, &jsonData)

	return jsonData, err
}

func transformData(csvData [][]string) ([]*model.User, error) {
	var oneRecord model.User
	var allRecords []model.User

	for i, each := range csvData {
		if i != 0 {
			oneRecord.Id , _ = strconv.Atoi(strings.TrimSpace(each[0]))
			oneRecord.Name = strings.TrimSpace(each[1])
			oneRecord.Username = strings.TrimSpace(each[2])
			oneRecord.Email = strings.TrimSpace(each[3])
			oneRecord.Phone = strings.TrimSpace(each[4])
			oneRecord.Website = strings.TrimSpace(each[5])
			allRecords = append(allRecords, oneRecord)
		}
	}

	r, err := json.Marshal(allRecords) // convert to JSON
	var jsonData []*model.User
	err = json.Unmarshal(r, &jsonData)
	return jsonData, err
}
