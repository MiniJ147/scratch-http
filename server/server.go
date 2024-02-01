package server

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	host string
	port string
}

type Config struct {
	HttpVersion string
}

func CreateServer(host string, port string) *Server {
	return &Server{
		host: host,
		port: port,
	}
}

func (serv *Server) Run() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", serv.host, serv.port))
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
