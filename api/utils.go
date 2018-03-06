package api

import (
	"encoding/csv"
	"log"
)

// BaseURL requests to other microservice
const BaseURL = "http://0.0.0.0:3000/api"

// W writer to produce CSV
var W *csv.Writer

// WriteCSV writes values to CSV
func WriteCSV(column int, values []string) {
	record := make([]string, column)
	for i, value := range values {
		record[i] = value
	}
	if err := W.Write(record); err != nil {
		log.Fatalln("error writing record to csv: ", err)
	}
}
