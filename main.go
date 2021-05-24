package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	PORT = ":" + PORT

	echoHandler := func(w http.ResponseWriter, req *http.Request) {

		xRealIp := req.Header.Get("x-real-ip")
		if xRealIp == "" {
			xRealIp = "Empty"
		}
		remoteAddress := req.RemoteAddr
		response := fmt.Sprintf("Your IP address is %s, and your 'X-Real-IP' header is %s\n", remoteAddress, xRealIp)
		_, err := io.WriteString(w, response)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("[Remote address]: %s \t[X-REAL-IP]: %s\n", remoteAddress, xRealIp)
	}

	http.HandleFunc("/", echoHandler)
	log.Fatal(http.ListenAndServe(PORT, nil))

}
