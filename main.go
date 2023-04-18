package main

import (
	"fmt"
	"github.com/SunilAtAnzx/collector/apis"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/upload", apis.UploadFile).Methods("POST")
	router.HandleFunc("/api/download", apis.DownloadFiles).Methods("GET")

	err := http.ListenAndServe(":5000", router)

	if err != nil {
		fmt.Println(err)
		return
	}

}
