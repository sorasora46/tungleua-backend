package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func main() {
	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		responseValue := map[string]any{"Message": "Hello, World"}
		response, _ := json.Marshal(responseValue)
		w.Write(response)
	})

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
