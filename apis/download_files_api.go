package apis

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func DownloadFiles(response http.ResponseWriter, request *http.Request) {

	files, err := os.ReadDir("coverageReports")
	if err != nil {
		log.Fatal(err)
	}
	archive, err := os.Create("coverageReportsArchive.zip")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func(archive *os.File) {
		err := archive.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(archive)
	zipWriter := zip.NewWriter(archive)

	for _, file := range files {
		err := copyFileToZip(zipWriter, "coverageReports/"+file.Name())
		if err != nil {
			panic(err)
		}
		fmt.Println(file.Name())
	}
	zipWriter.Close()

	archiveFile, err := os.Open("coverageReportsArchive.zip")
	defer archiveFile.Close()

	if err != nil {
		http.Error(response, "File not found.", 404)
		return
	}

	tempBuffer := make([]byte, 512)
	archiveFile.Read(tempBuffer)
	FileContentType := http.DetectContentType(tempBuffer)

	fmt.Println("File content")
	FileStat, _ := archiveFile.Stat()
	FileSize := strconv.FormatInt(FileStat.Size(), 10)

	Filename := "coverageReportsArchive"

	//Set the headers
	response.Header().Set("Content-Type", FileContentType+";"+Filename)
	response.Header().Set("Content-Length", FileSize)

	archiveFile.Seek(0, 0)
	io.Copy(response, archiveFile)
}

func copyFileToZip(zf *zip.Writer, filename string) error {
	r, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer r.Close()

	w, err := zf.Create(r.Name())
	if err != nil {
		return err
	}
	_, err = io.Copy(w, r)
	if err != nil {
		return err
	}
	return nil
}
