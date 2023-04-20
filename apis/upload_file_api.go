package apis

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func UploadFile(response http.ResponseWriter, request *http.Request) {

	//Max file size of 10MB
	err := request.ParseMultipartForm(10 * 1024 * 1204)

	if err != nil {
		return
	}

	file, handler, err := request.FormFile("coverageReport")

	if err != nil {
		log.Fatal(err)
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(file)

	fmt.Println("File Info")
	fmt.Println("File Name:", handler.Filename)
	fmt.Println("File Size:", handler.Size)

	n := fmt.Sprintf("coverageReport-%d.out", time.Now().UTC().Unix())
	dst := fmt.Sprintf("%s/%s", "coverageReports", n)
	out, err := os.Create(dst)

	_, err = io.Copy(out, file)

	if err != nil {
		log.Fatal(err)
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(out)

	fmt.Println("Upload Complete")
}
