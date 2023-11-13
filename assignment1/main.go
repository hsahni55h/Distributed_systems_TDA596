package main

import (
	"assignment1/server"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/* TODO: NEED A GENERIC IMPLEMENTATION OF HTTP REQUEST HANDLER */
func connectionRequestHandler(request *server.Message, response *server.Message) (err error) {
	/* Extract the HTTP method and path from the request */
	requestLine := string(request.Buff)
	parts := strings.Split(requestLine, " ")

	if len(parts) < 3 {
		return fmt.Errorf("invalid HTTP request format")
	}

	method := parts[0]
	path := parts[1]

	/* Process the request based on the HTTP method and path */
	switch method {
	case "GET":
		handleGetRequest(path, response)
	case "POST":
		handlePostRequest(path, request, response)
	default:
		return fmt.Errorf("unsupported HTTP method: %s", method)
	}

	return nil
}

func handlePostRequest(path string, request *server.Message, response *server.Message) {
	/* Simple response for a POST request */
	if path == "/submit" {
		/* Assuming data is sent in the request body */
		data := string(request.Buff)

		// Process the data as needed
		fmt.Println("Received data:", data)

		response.Buff = []byte("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n\r\nThanks for submitting data!\n")
	} else {
		response.Buff = []byte("HTTP/1.1 404 Not Found\r\nContent-Type: text/plain\r\n\r\n404 - Not Found\n")
	}
}

func handleGetRequest(path string, response *server.Message) {
	/* Simple response for a GET request */
	if path == "/" {
		response.Buff = []byte("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n\r\nHello, this is a GET request!\n")
	} else {
		response.Buff = []byte("HTTP/1.1 404 Not Found\r\nContent-Type: text/plain\r\n\r\n404 - Not Found\n")
	}
}

func logAndExitIfError(err error, expected error) {
	if err != expected {
		log.Fatal(err)
	}
}

func main() {
	/* check if user arg exist */
	if len(os.Args) < 2 {
		log.Fatal("Port number missing.\nExample: `-p 1278`")
	}

	config := &server.ServerConfig {
		Id: 0,
		NetworkType: "tcp",
		Port: "",
		MaxReq: 10,
		WaitConnMax: 10,
		ReqMsgHandler: connectionRequestHandler,
		IsStop: server.IsServerStopRequestNever,
	}
	
	/* flags declaration and parsing using flag package for user arg extraction */
	flag.StringVar(&config.Port, "p", "", "Specify port number. Example: -p 1278")
	flag.Parse()
	config.Port = ":" + config.Port
	
	a1Serv, err := server.Create(config)
	logAndExitIfError(err, nil)
	
	err = a1Serv.Run()
	logAndExitIfError(err, nil)
}

// command to run n programs in parallel on single command prompt without blocking it
// for /L %i in (1, 1, 20) do start /b curl http://127.0.0.1:1278
