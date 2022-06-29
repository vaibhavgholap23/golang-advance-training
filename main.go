package main

import (
	"encoding/csv"
	"encoding/json"
	"net/http"
	"os"

	"vaibhavgholap23/golang-advance-training.git/model"

	"go.uber.org/zap"
)

func main() {
	var userData model.User
	resp, err := http.Get("https://random-data-api.com/api/users/random_user")
	if err != nil {
		zap.S().Error(err.Error())
	}

	json.NewDecoder(resp.Body).Decode(&userData)
	csvFile, err := os.Create("user_data.csv")
	if err != nil {
		zap.S().Error(err.Error())
	}

	csvWrite := csv.NewWriter(csvFile)
	var row []string
	csvWrite.Write(row)
}
