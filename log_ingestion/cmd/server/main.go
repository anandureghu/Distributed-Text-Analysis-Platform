package main

import (
	"logprocessing/log-ingestion/api/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/upload", handlers.UploadLogFileHandler)
	http.ListenAndServe(":8080", nil)
}
