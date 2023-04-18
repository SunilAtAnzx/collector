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

	archiveFile, err := os.Open("coverageReportsArchive.zip") //Open the file to be downloaded later
	defer archiveFile.Close()                                 //Close after function return

	if err != nil {
		http.Error(response, "File not found.", 404) //return 404 if file is not found
		return
	}

	tempBuffer := make([]byte, 512)                       //Create a byte array to read the file later
	archiveFile.Read(tempBuffer)                          //Read the file into  byte
	FileContentType := http.DetectContentType(tempBuffer) //Get file header

	fmt.Println("File content")
	FileStat, _ := archiveFile.Stat()                  //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	Filename := "coverageReportsArchive"

	//Set the headers
	response.Header().Set("Content-Type", FileContentType+";"+Filename)
	response.Header().Set("Content-Length", FileSize)

	archiveFile.Seek(0, 0) //We read 512 bytes from the file already so we reset the offset back to 0
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
