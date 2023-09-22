package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/receive-iot-data", receiveIoTData)

	port := 8080 
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Listening on port %d...\n", port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Println("Error starting HTTP server:", err)
	}
}

func receiveIoTData(w http.ResponseWriter, r *http.Request) {
	for {
		if r.Method == http.MethodPost {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Error reading request body", http.StatusInternalServerError)
				return
			}

			fmt.Println("Received IoT data:", string(body))

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Data received successfully"))
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}
	time.Sleep(2 * time.Second)
}
