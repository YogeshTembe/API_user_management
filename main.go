package main

import (
	"net/http"

	"github.com/YogeshTembe/golang_project/controller"

	li "github.com/YogeshTembe/golang_project/logwrapper"
	"github.com/gorilla/mux"
)

var StandardLogger = li.NewLogger()

func main() {

	StandardLogger.Info("server starting...")
	controller.ConnectToDB()
	router := mux.NewRouter()

	router.HandleFunc("/api", controller.PostFilepath).Methods("POST")
	router.HandleFunc("/api/users", controller.GetUsers).Methods("GET")

	StandardLogger.Fatal(http.ListenAndServe(":8000", router).Error())
}
