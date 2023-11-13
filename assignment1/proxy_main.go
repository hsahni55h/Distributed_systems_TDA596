package main

import (
	"assignment1/server"
	"bufio"
	"bytes"
	"fmt"
	"flag"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func forwardRequest(request *server.Message, response *server.Message) error {
	// Parse the HTTP request
	httpRequest, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(request.Buff)))
	if err != nil {
		response.Buff = []byte("HTTP/1.1 400 Bad Request\r\nContent-Type: text/plain\r\n\r\nInvalid HTTP request format")
		return err
	}

	// Only handle GET requests
	if httpRequest.Method != http.MethodGet {
		response.Buff = []byte("HTTP/1.1 501 Not Implemented\r\nContent-Type: text/plain\r\n\r\nOnly GET requests are supported")
		return nil
	}

	// Extract the host and port from the request
	host := httpRequest.URL.Host
	port := "80" // Default port for HTTP

	// Split host and port
	if strings.Contains(host, ":") {
		parts := strings.Split(host, ":")
		host = parts[0]
		port = parts[1]
	}

	// Create a connection to the remote server
	remoteConn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		response.Buff = []byte("HTTP/1.1 502 Bad Gateway\r\nContent-Type: text/plain\r\n\r\nError connecting to the remote server")
		return err
	}
	defer remoteConn.Close()

	// Forward the GET request to the remote server
	if err := httpRequest.Write(remoteConn); err != nil {
		response.Buff = []byte("HTTP/1.1 502 Bad Gateway\r\nContent-Type: text/plain\r\n\r\nError forwarding the request")
		return err
	}

	// Read the response from the remote server
	remoteResp, err := http.ReadResponse(bufio.NewReader(remoteConn), httpRequest)
	if err != nil {
		response.Buff = []byte("HTTP/1.1 502 Bad Gateway\r\nContent-Type: text/plain\r\n\r\nError reading the remote server response")
		return err
	}
	defer remoteResp.Body.Close()

	// Forward the response to the client
	responseBuf, err := httputil.DumpResponse(remoteResp, true)
	if err != nil {
		response.Buff = []byte("HTTP/1.1 502 Bad Gateway\r\nContent-Type: text/plain\r\n\r\nError reading the remote server response")
		return err
	}
	
	response.Buff = responseBuf
	return nil
}

func main() {
	/* check if user arg exist */
	if len(os.Args) < 2 {
		log.Fatal("Port number missing.\nExample: `-p 1278`")
	}

	config := &server.ServerConfig{
		Id:           0,
		NetworkType:  "tcp",
		Port:         "",
		MaxReq:       10,
		WaitConnMax:  10,
		ReqMsgHandler: forwardRequest,
		IsStop:       server.IsServerStopRequestNever,
	}
	
	/* flags declaration and parsing using flag package for user arg extraction */
	flag.StringVar(&config.Port, "p", "", "Specify port number. Example: -p 1278")
	flag.Parse()
	config.Port = ":" + config.Port

	proxyServer, err := server.Create(config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Proxy server is running. Waiting for client requests...")
	err = proxyServer.Run()
	if err != nil {
		log.Fatal(err)
	}
}
