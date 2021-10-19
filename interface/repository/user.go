package repository

import (
	"academy-go-q32021/domain/model"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var lock = new(sync.Mutex)

type user struct{}

type User interface {
	ReadUsers() ([]*model.User, error)
	ReadUsersByKey(f string) ([]*model.CustomCSV, error)
	GetUsers(u []*model.User) ([]*model.User, error)
	GetUserById(id int) (*model.User, error)
	GetUsersConcurrently(itemType string, items, itemsWorker int) ([]*model.User, error)
}

var URL = "https://jsonplaceholder.typicode.com/users/"
var CSVFILE = "./public/data.csv"

func NewUserRepository() User {
	return &user{}
}

func (ur *user) ReadUsers() ([]*model.User, error) {
	csvFile, err := openFile(CSVFILE)
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

// ReadUsersByKey read users from CSV by key
func (ur *user) ReadUsersByKey(k string) ([]*model.CustomCSV, error) {
	csvFile, err := openFile(CSVFILE)
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

// GetUsers read users from CSV
func (ur *user) GetUsers(u []*model.User) ([]*model.User, error) {
	response, err := http.Get(URL)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(response.Body).Decode(&u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// GetUserById get user by id from API
func (ur *user) GetUserById(id int) (*model.User, error) {
	endPoint := fmt.Sprint(URL, id)
	response, err := http.Get(endPoint)
	if err != nil {
		return nil, err
	}

	var u *model.User

	err = json.NewDecoder(response.Body).Decode(&u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// GetUsersConcurrently get users concurrently
func (ur *user) GetUsersConcurrently(itemType string, items, itemsWorker int) ([]*model.User, error) {
	var result [][]string
	content := make(chan []string, items)
	workers := items / itemsWorker

	file, err := openFile(CSVFILE)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	r := csv.NewReader(file)

	for w := 0; w <= workers; w++ {
		wg.Add(1)
		go worker(r, itemType, itemsWorker, content)
	}

	go func(rows chan []string) {
		wg.Wait()
		close(rows)
	}(content)

	for row := range content {
		result = append(result, row)
	}

	jsonData, err := transformData(result)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

// openFile to open files with a path name
func openFile(f string) (*os.File, error) {
	csvFile, err := os.Open(f)
	if err != nil {
		err = fmt.Errorf("path provided was not found")
		return nil, err
	}
	return csvFile, nil
}

// readFile to read files with a path name
func readFile(csvFile *os.File) (records [][]string, err error) {
	reader := csv.NewReader(csvFile)

	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()

	return csvData, err
}

// transformDataByKey to transform all data and return it as model
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

// transformData to transform all data and return it as model
func transformData(csvData [][]string) ([]*model.User, error) {
	var oneRecord model.User
	var allRecords []model.User

	for i, each := range csvData {
		if i != 0 {
			oneRecord.Id, _ = strconv.Atoi(strings.TrimSpace(each[0]))
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

// worker to read all lines and write data on the channel
func worker(r *csv.Reader, itemType string, itemsWorker int, content chan<- []string) {
	defer wg.Done()
	counter := 0

	for {
		if counter == itemsWorker {
			break
		}

		lock.Lock()
		row, err := r.Read()

		if err == io.EOF {
			break
		}

		if len(row) == 0 {
			break
		}

		if validateType(itemType, row) {
			content <- row
		}

		lock.Unlock()
		counter++
	}
}

// validateType to return valid data to storage into channel
func validateType(itemType string, row []string) bool {
	id, _ := strconv.Atoi(strings.TrimSpace(row[0]))
	switch itemType {
	case "odd":
		return id%2 != 0
	case "even":
		return id%2 == 0
	default:
		return false
	}
}
