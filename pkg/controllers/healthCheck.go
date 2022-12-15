package controllers

import (
	"time"
	"log"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
)

type healthCheckResponse struct {
	Uptime   float64
	Message  string
	Timestamp time.Time
}

func HealthCheck(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	startTime := time.Now()
	response := healthCheckResponse{
		Uptime:   time.Since(startTime).Seconds(),
		Message:  "OK",
		Timestamp: time.Now(),
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		log.Printf("Call to health check endpoint failed with error '%s'", err)
		response.Message = err.Error()
		w.WriteHeader(http.StatusServiceUnavailable)
		encoder.Encode(response)
	}
}
