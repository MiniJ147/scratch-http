package server

import (
	"fmt"
	"net"
)

type Client struct {
	conn     net.Conn
	httpServ *HttpServer
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

		response := createHttpResponse(client.conn)
		client.httpServ.methods["GET"][0].function(*response)

		client.conn.Close()
	}
}
