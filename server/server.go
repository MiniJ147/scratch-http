package server

/*
	Handles the tcp connections behind the scenes. No fancy http handling here
*/

import (
	"fmt"
	"log"
	"net"
)

type server struct {
	host string
	port string
}

type config struct {
	HttpVersion string
}

func createServer(host string, port string) *server {
	return &server{
		host: host,
		port: port,
	}
}

func (serv *server) run(httpServ *HttpServer) {
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

		client := &client{
			conn:     conn,
			httpServ: httpServ,
		}

		go client.handleRequest()
	}
}
