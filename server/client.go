package server

import (
	"fmt"
	"net"
	"time"
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

		dt := time.Now()
		res := createHttpResponse(client.conn)

		res.writeStatus(200, "OK")
		res.writeHeader("Date", dt.Format("01-02-2006 15:04:05"))
		res.writeHeader("Connection", "Keep-Alive")
		res.writeHeader("Keep-Alive", "timeout=5")
		res.writeHeader("Content-Length", "24")
		res.writeHeader("Content-Type", "text/html")

		res.send("<html><h1>SUP</h1></html>")

		client.conn.Close()
	}
}
