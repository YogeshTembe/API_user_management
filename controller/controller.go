package controller

import (
	"encoding/json"

	"net/http"

	"github.com/YogeshTembe/golang_project/file_handler"
	li "github.com/YogeshTembe/golang_project/logwrapper"
	"github.com/YogeshTembe/golang_project/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var StandardLogger = li.NewLogger()

const (
	dbUser     = "root"
	dbPassword = "Panda@19"
	dbName     = "users"
	dbHost     = "127.0.0.1:3306"
)

func ConnectToDB() {
	var err error
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		StandardLogger.Fatal("failed to connect database-" + err.Error())
	} else {
		StandardLogger.Info("connected to database")
	}
}

func PostFilepath(w http.ResponseWriter, r *http.Request) {
	var file model.File
	json.NewDecoder(r.Body).Decode(&file)
	csvFile := file_handler.OpenCSVFile(file.FilePath)
	defer csvFile.Close()
	users := file_handler.ReadCSVFile(csvFile)
	file_handler.WriteJSONFile("users.json", users)
	CreateUsers(users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Succesfully inserted in DB.")
}

func CreateUsers(users []model.User) {
	for _, user := range users {
		if db.Model(&user).Where("id = ?", user.Id).Updates(&user).RowsAffected == 0 {
			db.Create(&user)
		}
	}
	StandardLogger.Info("Succesfully added users data in database.")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []model.User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}
