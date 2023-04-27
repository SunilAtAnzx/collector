package main

import (
	"flag"
	"fmt"
	"github.com/SunilAtAnzx/collector/apis"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	var port int
	flag.IntVar(&port, "p", 8080, "Provide a port number")
	flag.Parse()

	router := mux.NewRouter()

	router.HandleFunc("/api/upload", apis.UploadFile).Methods("POST")
	router.HandleFunc("/api/download", apis.DownloadFiles).Methods("GET")

	err := http.ListenAndServe(fmt.Sprint(":", port), router)

	if err != nil {
		fmt.Println(err)
		return
	}

}
