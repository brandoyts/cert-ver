package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	_PORT string = ":7070"
)

func requestHandler(writer http.ResponseWriter, request *http.Request) {
	// read body
	body, err := io.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}

	fmt.Printf("Request Body: %v\n", string(body))

	// read headers
	xRealIP := request.Header.Get("X-Real-IP")
	xForwaredIP := request.Header.Get("X-Forwarded-For")
	verified := request.Header.Get("VERIFIED")
	issuer := request.Header.Get("DN")

	fmt.Printf("X-Real-IP: %v\n", xRealIP)
	fmt.Printf("X-Forwared-For: %v\n", xForwaredIP)
	fmt.Printf("Verified: %v\n", verified)
	fmt.Printf("Certificate Issuer: %v\n", issuer)

	// send response
	response := `{"status": "authenticated"}`
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	io.WriteString(writer, string(response))
}

func main() {
	http.HandleFunc("/", requestHandler)
	fmt.Printf("starting server on port %v\n", _PORT)
	err := http.ListenAndServe(_PORT, nil)
	if err != nil {
		println(err.Error)
		os.Exit(1)
	}
}
