package file_handler

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	li "github.com/YogeshTembe/golang_project/logwrapper"
	"github.com/YogeshTembe/golang_project/model"
	"github.com/YogeshTembe/golang_project/validation"
	uuid "github.com/satori/go.uuid"
)

var StandardLogger = li.NewLogger()

func OpenCSVFile(fileDir string) *os.File {
	csvFile, err := os.Open(fileDir)
	if err != nil {
		StandardLogger.Fatal(err)
	}
	fmt.Println("Successfully Opened CSV file")
	return csvFile
}

func ReadCSVFile(csvFile *os.File) []model.User {
	var users []model.User

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		StandardLogger.Fatal("Unable to read the CSV file-" + err.Error())
	}

	for _, line := range csvLines {
		id, _ := uuid.FromString(line[0])
		phoneNo, _ := strconv.Atoi(line[3])
		isActive, _ := strconv.ParseBool(line[4])

		user := model.User{
			Id:          id,
			Name:        line[1],
			Email:       line[2],
			PhoneNumber: phoneNo,
			IsActive:    isActive,
		}
		isValid := validation.Validate(&user)

		if isValid {
			users = append(users, user)
			validation.UserIds[user.Id.String()] = struct{}{}
		}
	}
	StandardLogger.Info("CSV file reading and data validation is done.")
	return users
}

func WriteJSONFile(fileDir string, users []model.User) {
	file, _ := json.MarshalIndent(users, "", " ")
	_ = ioutil.WriteFile(fileDir, file, 0644)
	StandardLogger.Info("JSON file writing is done.")
}
