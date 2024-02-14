package server

import (
	"net"
)

type Client struct {
	conn     net.Conn
	httpServ *HttpServer
}

func (client *Client) handleRequest() {
	buffer := make([]byte, 2048*2) //where request is stored
	for {
		data, err := client.conn.Read(buffer)
		if err != nil {
			//fmt.Println("Read Error: ", err)
			continue
		}

		requestHeaderString := string(buffer[:data])

		//fmt.Println(requestHeaderString)

		request := CreateHttpRequest(requestHeaderString)

		response := createHttpResponse(client.conn)
		client.httpServ.Find(request.method, request.route).function(*response)

		client.conn.Close()
	}
}
