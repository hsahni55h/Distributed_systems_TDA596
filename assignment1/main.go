package main

import (
	"assignment1/server"
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

/* TODO: re-think of generalizing if possible (not attach it to main.go scope) */
var fileMutex sync.Mutex

func connectionRequestHandler(request *server.Message, response *server.Message) (err error) {
	/*	reads and parses an incoming HTTP request from buffer
		bufio wraps an existing bytes.Reader type (that implements io.Reader interface's Read() function) and provide buffering for it. 
		Buffering imporves performance since it reduces the number of actual reads from the underlying io.Reader (bytes.Reader in our case)
	*/
	httpRequest, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(request.Buff)))
	if err != nil {
		// Error parsing the request
		response.Buff = []byte("HTTP/1.1 400 Bad Request\r\nContent-Type: text/plain\r\n\r\nInvalid HTTP request format.\n")
		return
	}

	/* Extract the HTTP method and path from the request */
	method := httpRequest.Method
	path := httpRequest.URL.Path

	/* Process the request based on the HTTP method and path */
	switch method {
	case "GET":
		handleGetRequest(path, httpRequest, response)
	case "POST":
		handlePostRequest(path, httpRequest, response)
	default:
		/* Unsupported HTTP method */
		response.Buff = []byte("HTTP/1.1 501 Not Implemented\r\nContent-Type: text/plain\r\n\r\nUnsupported HTTP method.\n")
	}

	return nil
}

func handleGetRequest(path string, httpRequest *http.Request, response *server.Message) {
	/* Validate the file extension */
	ext := filepath.Ext(path)
	contentType := ""

	switch ext {
	case ".html":
		contentType = "text/html"
	case ".txt":
		contentType = "text/plain"
	case ".gif":
		contentType = "image/gif"
	case ".jpeg", ".jpg":
		contentType = "image/jpeg"
	case ".css":
		contentType = "text/css"
	default:
		/* Invalid file extension */
		response.Buff = []byte("HTTP/1.1 400 Bad Request\r\nContent-Type: text/plain\r\n\r\nInvalid file extension.\n")
		return
	}

	/* Read the file and send the response */
	filePath := filepath.Join("files", path[1:])
	content, err := os.ReadFile(filePath)

	if err != nil {
		/* File not found */
		response.Buff = []byte("HTTP/1.1 404 Not Found\r\nContent-Type: text/plain\r\n\r\nFile not found.\n")
		return
	}

	/* Successful response for GET request */
	response.Buff = []byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: %s\r\n\r\n", contentType))
	response.Buff = append(response.Buff, content...)
}


func handlePostRequest(path string, httpRequest *http.Request, response *server.Message) {
	/* Validate the file extension */
	ext := filepath.Ext(path)
	filePath := filepath.Join("files", path[1:])

	switch ext {
	case ".html", ".css":
		/* Not allowed to change these file formats */
		response.Buff = []byte("HTTP/1.1 403 Forbidden\r\nContent-Type: text/plain\r\n\r\nCannot modify HTML and CSS files.\n")
		return
	case ".txt":
		/* Append data to existing file or return 404 if the file doesn't exist */
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			response.Buff = []byte("HTTP/1.1 404 Not Found\r\nContent-Type: text/plain\r\n\r\nFile not found.\n")
			return
		}

		/* Open the file in append mode */
		file, err := os.OpenFile(filePath, os.O_APPEND | os.O_WRONLY, 0644)
		if err != nil {
			response.Buff = []byte("HTTP/1.1 500 Internal Server Error\r\nContent-Type: text/plain\r\n\r\nError opening file.\n")
			return
		}
		defer file.Close()
		
		/* Read the request body and append data to the existing file */
		bodyBytes, err := io.ReadAll(httpRequest.Body)
		if err != nil {
			response.Buff = []byte("HTTP/1.1 500 Internal Server Error\r\nContent-Type: text/plain\r\n\r\nError reading request body.\n")
			return
		}
		bodyBytes = append(bodyBytes, '\n')
		
		fileMutex.Lock()
		_, err = file.Write(bodyBytes)
		if err != nil {
			response.Buff = []byte("HTTP/1.1 500 Internal Server Error\r\nContent-Type: text/plain\r\n\r\nError appending to file.\n")
			return
		}
		fileMutex.Unlock()

		/* Successful response for POST request */
		response.Buff = []byte("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n\r\nData appended to file successfully.\n")
	default:
		/* Replace the complete file with the received data without the header */
		bodyBytes, err := io.ReadAll(httpRequest.Body)
		if err != nil {
			response.Buff = []byte("HTTP/1.1 500 Internal Server Error\r\nContent-Type: text/plain\r\n\r\nError reading request body.\n")
			return
		}

		err = os.WriteFile(filePath, bodyBytes, 0644)
		if err != nil {
			response.Buff = []byte("HTTP/1.1 500 Internal Server Error\r\nContent-Type: text/plain\r\n\r\nError writing to file.\n")
			return
		}

		/* Successful response for POST request */
		response.Buff = []byte("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n\r\nFile replaced successfully.\n")
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
	
	/* create and start the server */
	a1Serv, err := server.Create(config)
	logAndExitIfError(err, nil)
	
	go func() {
		err = a1Serv.Run()
		logAndExitIfError(err, nil)
	}()
	fmt.Println("Server is running. Waiting for client(s)...")
	
	/* sleep infinitely since there is nothing else to do */
	for {
		time.Sleep(time.Second)
	}
}

