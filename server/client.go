package server

import (
	"fmt"
	"net"
)

type Client struct {
	conn net.Conn
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

		res := createHttpResponse(client.conn)

		//test
		type Msg struct {
			Message string `json:"Message"`
		}

		res.writeStatus(200, "OK")
		res.sendJSON(Msg{
			Message: "HELLO1",
		})

		client.conn.Close()
	}
}
