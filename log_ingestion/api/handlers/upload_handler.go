package handlers

import (
	"fmt"
	"logprocessing/log-ingestion/internal/service/log"
	"net/http"
)

// HTTP handler to receive log files via HTTP
func UploadLogFileHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Handle file upload from the request and call ingestLogFile
	log.IngestLogFile("")
	fmt.Fprintf(w, "Log file received and ingested.")
}
