package server

import (
	"fmt"
	"net"
	"strings"
)

type Client struct {
	conn     net.Conn
	httpServ *HttpServer
}

func (client *Client) handleRequest() {
	buffer := make([]byte, 2048*4) //where request is stored
	for {
		data, err := client.conn.Read(buffer)
		if err != nil {
			//fmt.Println("Read Error: ", err)
			continue
		}

		requestHeaderString := string(buffer[:data])

		//fmt.Println(requestHeaderString)

		request := CreateHttpRequest(requestHeaderString)

		fmt.Println("client: ", request.route)
		response := createHttpResponse(client.conn)

		//redirects request to render css
		if strings.Contains(request.route, "/css/") {
			fmt.Println("need to intercept")
			route, err := client.httpServ.Find("GET", "/css")
			if err != nil {
				fmt.Println(err)
				response.WriteStatus(500, "INTERNAL SERVER BREAK")
				response.Send("")
				client.conn.Close()
				return
			}
			route.function(request, *response)
		} else {
			route, err := client.httpServ.Find(request.method, request.route)
			if err != nil {
				fmt.Println(err)
				response.WriteStatus(404, "COULD NOT FIND ROUTE")
				response.Send("")
				client.conn.Close()
				return
			}
			route.function(request, *response)
		}

		client.conn.Close()
	}
}
