package model

type CustomCSV struct {
	Value string `json:"value"`
}

func (CustomCSV) TableName() string { return "data" }
