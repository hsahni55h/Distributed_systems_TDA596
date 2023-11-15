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
	"time"
)

func forwardRequest(request *server.Message, response *server.Message) error {
	/*	reads and parses an incoming HTTP request from buffer
		bufio wraps an existing bytes.Reader type (that implements io.Reader interface's Read() function) and provide buffering for it. 
		Buffering imporves performance since it reduces the number of actual reads from the underlying io.Reader (bytes.Reader in our case)
	*/
	httpRequest, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(request.Buff)))
	if err != nil {
		response.Buff = []byte("HTTP/1.1 400 Bad Request\r\nContent-Type: text/plain\r\n\r\nInvalid HTTP request format")
		return err
	}

	/* Only handle GET requests */
	if httpRequest.Method != http.MethodGet {
		response.Buff = []byte("HTTP/1.1 501 Not Implemented\r\nContent-Type: text/plain\r\n\r\nOnly GET requests are supported")
		return nil
	}

	/* Create a connection to the remote server */
	remoteConn, err := net.Dial("tcp", httpRequest.URL.Host)
	if err != nil {
		response.Buff = []byte("HTTP/1.1 502 Bad Gateway\r\nContent-Type: text/plain\r\n\r\nError connecting to the remote server")
		return err
	}
	defer remoteConn.Close()

	/* Forward the GET request to the remote server */
	if err := httpRequest.Write(remoteConn); err != nil {
		response.Buff = []byte("HTTP/1.1 502 Bad Gateway\r\nContent-Type: text/plain\r\n\r\nError forwarding the request")
		return err
	}

	/* Read the response from the remote server */
	remoteResp, err := http.ReadResponse(bufio.NewReader(remoteConn), httpRequest)
	if err != nil {
		response.Buff = []byte("HTTP/1.1 502 Bad Gateway\r\nContent-Type: text/plain\r\n\r\nError reading the remote server response")
		return err
	}
	defer remoteResp.Body.Close()

	/* Forward the response to the client */
	responseBuf, err := httputil.DumpResponse(remoteResp, true)
	if err != nil {
		response.Buff = []byte("HTTP/1.1 502 Bad Gateway\r\nContent-Type: text/plain\r\n\r\nError reading the remote server response")
		return err
	}
	
	response.Buff = responseBuf
	return nil
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
	logAndExitIfError(err, nil)

	go func() {
		err = proxyServer.Run()
		logAndExitIfError(err, nil)
	}()
	fmt.Println("Proxy server is running. Waiting for client requests...")

	/* sleep infinitely since there is nothing else to do */
	for {
		time.Sleep(time.Second)
	}
}
