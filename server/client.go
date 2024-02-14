package server

import (
	"fmt"
	"net"
)

type Client struct {
	conn     net.Conn
	httpServ *HttpServer
}

// might refactor later
func parseMethodAndRoute(headerString string) (string, string) {
	requestMethod := ""
	requestRoute := ""
	count := 0
	i := 0

	for {
		currChar := string(headerString[i])

		if count >= 2 {
			break
		}
		if currChar == " " {
			count += 1
			i += 1
			continue
		}

		if count < 1 {
			requestMethod += currChar
		} else {
			requestRoute += currChar
		}

		i += 1
	}

	return requestMethod, requestRoute
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

		method, route := parseMethodAndRoute(requestHeaderString)

		fmt.Println("Found: ", method, route)

		response := createHttpResponse(client.conn)
		client.httpServ.Find(method, route).function(*response)

		client.conn.Close()
	}
}
