package server

import (
	"assignment1/queue"
	"log"
	"net"
	"sync"
)

/* On success, ConnReqHandler returns nil as error */
type ConnReqHandler func(request *Message, response *Message) (err error)
type IsServerStopRequest func(server *Server) bool

func IsServerStopRequestNever(server *Server) bool {
	return false
}

type Message struct {
	Node string					/* Node which sends or receives the message */
	Buff []byte					/* Buffeer to store the received data */
}

type Connection struct {
	Stream net.Conn
	Msg *Message
	Err error
}

type ServerConfig struct {
	Id int64						/* unique id to distingish between different server objects */
	NetworkType string				/* Valid Values: "tcp", "udp" */
	Port string						/* Example: ":1728" */
	MaxReq int64					/* max connection requests that the server can handle in parallel */
	WaitConnMax int64				/* max waiting connections allowed to queue */
	ReqMsgHandler ConnReqHandler	/* the function that runs when a conenction request is received by the server */
	IsStop IsServerStopRequest		/* the function that checks if server needs to stop/halt */
}

type serverFd struct {
	Id int64									/* from config */
	NetworkType string							/* from config */
	Port string									/* from config */
	MaxReq int64								/* from config */
	WaitConnMax int64							/* from config */
	ReqMsgHandler ConnReqHandler				/* from config */
	IsStop IsServerStopRequest					/* from config */
	Addr string									/* IP address: "127.0.0.1:1728" */
	ReqCount int64								/* current number of request that are currently handled in parallel */
	ReqCountLock sync.Mutex						/* since ReqCount is shared and modified between various connection (gorountines) it needs to be protected (atomic operations) */
	// WaitConnBuffer []Connection				/* all waiting connections are stored here */
	Listner net.Listener						/* the socket to listen */
	WaitConnQueue *queue.Queue[*Connection]		/* connection FIFO that are waiting to be handle */
}

type Server struct {
	fd *serverFd
}

func Create(config *ServerConfig) (serv Server, err error) {
	listner, err := net.Listen(config.NetworkType, config.Port)
	if err != nil {
		return serv, err
	}

	serv.fd = &serverFd{
		Id: config.Id,
		NetworkType: config.NetworkType,
		Port: config.Port,
		MaxReq: config.MaxReq,
		WaitConnMax: config.WaitConnMax,
		ReqMsgHandler: config.ReqMsgHandler,
		IsStop: config.IsStop,
		Addr: "" + config.Port,			/* TOOD: In later versions it can be any generic IP */
		WaitConnQueue: queue.New[*Connection](int(config.WaitConnMax)),
		Listner: listner,
	}
	/*
		ReqCount: already initialized to default value
		ReqCountLock: already initialized to default value
	*/

	return serv, nil
}

/* 
	Always prefer the use of log.Println() instead of log.Fatal() to avoid sudden shut-down of server.
	This will cause sudden termination the program without closing the connection gracefully.
*/
func (serv *Server) __connectionRequestHandler() {
	/* get the connection to handle from waiting queue and defer to close the connection stream */
	connection, err := serv.fd.WaitConnQueue.Dequeue()
	if err != nil {
		log.Println(err)
		return
	}
	defer connection.Stream.Close()

	/* read the request */
	connection.Msg.Node = connection.Stream.RemoteAddr().String()
	_, err = connection.Stream.Read(connection.Msg.Buff)
	if err != nil {
		log.Println(err)
		return
	}
	
	/* create a response packet */
	response := Message{
		Node: connection.Msg.Node,
		Buff: make([]byte, 0, cap(connection.Msg.Buff)),
	}
	
	/* handle the request */
	err = serv.fd.ReqMsgHandler(connection.Msg, &response)
	if (err != nil) {
		log.Println(err)
		return
	}

	/* write the response */
	_, err = connection.Stream.Write(response.Buff[:len(response.Buff)])
	if (err != nil) {
		log.Println(err)
		return
	}

	// atomic.AddInt64(&serv.fd.ReqCount, -1)
	serv.fd.ReqCountLock.Lock()
	serv.fd.ReqCount--
	serv.fd.ReqCountLock.Unlock()
}

func (serv *Server) Run() (err error) {
	// var wgroup sync.WaitGroup	/* used to wait for all the goroutines launched here to finish */
	for !serv.fd.IsStop(serv) {
		/* if waiting queue is NOT full then wait for the new connection and enqueue the it to the waiting queue else don't wait for new connections */
		if int64(serv.fd.WaitConnQueue.Size()) < serv.fd.WaitConnMax {
			/* wait for the new connection and enqueue it to waiting queue */
			newConn, err := serv.fd.Listner.Accept()
			if (err != nil) {
				log.Println(err)
				return err
			}
			var newConnection = &Connection {
				Stream: newConn,
				Msg: &Message {
					Buff: make([]byte, 1024),
				},
				Err: nil,
			}
			
			serv.fd.WaitConnQueue.Enqueue(newConnection)
		} 
		
		/* 	
			if a connection is in waiting queue AND 
			ReqCount < MaxReq then 
			dequeue a connection from the waiting queue and handle it 
		*/
		if (serv.fd.WaitConnQueue.Size() > 0) && (serv.fd.ReqCount < serv.fd.MaxReq) {
			// atomic.AddInt64(&serv.fd.ReqCount, 1)
			serv.fd.ReqCountLock.Lock()
			serv.fd.ReqCount++
			serv.fd.ReqCountLock.Unlock()

			go serv.__connectionRequestHandler()
		}
	}

	// /* always wait for all connections to close before stopping the server. */
	// wgroup.Wait()
	return err
}


