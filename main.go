package main

import (
	"encoding/csv"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"vaibhavgholap23/golang-advance-training.git/model"

	"go.uber.org/zap"
)

func main() {

	var userData model.User
	csvFile, err := os.Create("user_data.csv")
	if err != nil {
		zap.S().Error(err.Error())
	}
	csvWrite := csv.NewWriter(csvFile)
	headerLevel1 := []string{"Sr No.", "ID", "UID", "Password", "First Name", "Last Name", "Username", "Email", "Avatar", "Gender", "Phone Number", "Social Insurance Number", "DOB", "Employment", "", "Address", "", "", "", "", "", "", "", "CreditCard", "Subscription"}
	headerLevel2 := []string{"", "", "", "", "", "", "", "", "", "", "", "", " ", "Title", "Key Skill", "City", "Street Name", "Street Address", "Zip Code", "State", "Country", "Coordinates", "", "CreditCard", "Plan", "Status", "Payment Method", "Term"}
	headerLevel3 := []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "Lat", "Lng"}

	csvWrite.Write(headerLevel1)
	csvWrite.Write(headerLevel2)
	csvWrite.Write(headerLevel3)
	for i := 1; i <= 5; i++ {
		resp, err := http.Get("https://random-data-api.com/api/users/random_user")
		if err != nil {
			zap.S().Error(err.Error())
		}

		json.NewDecoder(resp.Body).Decode(&userData)

		var row []string
		row = append(row, strconv.Itoa(int(i)))
		row = append(row, strconv.Itoa(int(userData.ID)))
		row = append(row, userData.UID)
		row = append(row, userData.Password)
		row = append(row, userData.FirstName)
		row = append(row, userData.LastName)
		row = append(row, userData.Username)
		row = append(row, userData.Email)
		row = append(row, userData.Avatar)
		row = append(row, userData.Gender)
		row = append(row, userData.PhoneNumber)
		row = append(row, userData.SocialInsuranceNumber)
		row = append(row, userData.DateOfBirth)

		csvWrite.Write(row)
	}
	defer csvWrite.Flush()

}
