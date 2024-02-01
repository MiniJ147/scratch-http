package main

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	host string
	port string
}

type Client struct {
	conn net.Conn
}

func createServer(host string, port string) *Server {
	return &Server{
		host: host,
		port: port,
	}
}

func (client *Client) handleRequest() {
	buffer := make([]byte, 2048*2)
	for {
		data, err := client.conn.Read(buffer)
		if err != nil {
			//fmt.Println("Read Error: ", err)
			continue
		}

		msg := buffer[:data]
		fmt.Println(string(msg))

		body := "<html><h1>hey</h1></html"
		fmt.Println(len(body))

		response := []byte("HTTP/1.1 200 OK\r\nDate: Wed, 10 Aug 2016 13:17:18GMT\r\nConnection: Keep-Alive\r\nKeep-Alive: timeout=5\r\nContent-Length: 24\r\nContent-Type: text/html\r\n\r\n")
		client.conn.Write(response)

		client.conn.Write([]byte("<html><h1>hey</h1></html>"))
		client.conn.Close()
	}
}

func (server *Server) run() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", server.host, server.port))
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		client := &Client{
			conn: conn,
		}

		go client.handleRequest()
	}
}

func main() {
	fmt.Println("Hello From Server")

	serv := createServer("localhost", "3000")
	serv.run()

}
