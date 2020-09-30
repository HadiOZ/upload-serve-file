package handler

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Upload(w http.ResponseWriter, r *http.Request) {

	//memeriksa method request
	if r.Method != http.MethodPost {
		http.Error(w, "Just only allow Method POST", http.StatusInternalServerError)
		return
	}

	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), 500)
		log.Panicf("[x] Can't parse multipart form : %v", err)
		return
	}

	fileUpload, handler, err := r.FormFile("file")
	defer fileUpload.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Panicf("[x] File is not found in request form file : %v", err)
		return
	}

	log.Println("New file have Uploded : ", handler.Filename)

	path, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Panicf("[x] Can't get info of this directory : %v", err)
		return
	}

	fileLocation := filepath.Join(path, "file-received", handler.Filename)

	newFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	defer newFile.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Panicf("[x] Can't open or make file : %v", err)
		return
	}

	if _, err := io.Copy(newFile, fileUpload); err != nil {
		http.Error(w, err.Error(), 500)
		log.Panicf("[x] Can't copy file : %v", err)
		return
	}

	log.Println("file was seved to directory : ", fileLocation)

	w.Write([]byte("Seved"))
}

func MultiUpload(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Just only allow Method POST", http.StatusInternalServerError)
		return
	}

	path, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Panicf("[x]  Can't get info of this directory : %v", err)
		return
	}

	reader, err := r.MultipartReader()
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Panicf("[x] Can't open reader data : %v", err)
		return
	}

	for {
		fileUpload, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		fileLocation := filepath.Join(path, "file-received", fileUpload.FileName())

		newFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
		defer newFile.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			log.Panicf("[x] Can't open or make file : %v", err)
			return
		}

		if _, err := io.Copy(newFile, fileUpload); err != nil {
			http.Error(w, err.Error(), 500)
			log.Panicf("[x] Can't copy file : %v", err)
			return
		}

		log.Println("file was seved to directory : ", fileLocation)
	}

	w.Write([]byte("Seved"))
}
