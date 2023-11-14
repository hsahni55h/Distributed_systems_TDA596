/* A  TCP server is a simple process that runs in a machine that listens to a port.
Any machine that wants to talk to a server has to connect over the port and establish the connection.

Step 1: Start Listening on the port

when your process starts, pick a process and start listening to it
net.Listen("tcp", ":1729")

Step 2: Wait for a client to connect

invoke the accept system call and wait for a client to connect.
This is a blocking call and your server would not procedd until some client connects.


Step 3: Once the connection is established
1. Invoke the 'read' system call to read the request		- conn.Read()
2. Invoke the 'write' system call to read the request		- conn.Write()
3. close the connection										- conn.close()
Note - Read and Write are blocking calls. When we invoke them, unless the client sends something the process is blocked

Step 4: Do this over and over again
Put this entire thing in an infinite loop

For loop ....const
1. Continuously waiting for client to connect
2. Readign the request
3. Writing the request
4. Closing the connection

Step 5: Multiple Request Concurrently  (accepting one, processing and then accepting another one)
Sequential Execution and Handling

*/
// The server that we have written here cannot handle concurrent requests coming in, it handles only one request at at time.
// This is a single thread server.

package main

import (
	"log"
	"net"
	"time"
)

func do(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf) // number of byter that I read (we are skipping it here)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Processing the request")
	time.Sleep(8 * time.Second)

	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello World!\r\n"))
	conn.Close()
}

func main() {

	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal(err)
	}

	for {
		log.Println("Waiting for the client to connect")
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Client connected")
		do(conn)

	}

}
