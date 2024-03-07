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

// handles intercept and returns true if intercept happend and false if not
func (client *Client) handleIntercept(req *HttpRequest, res *HttpResponse, fileType string) bool {
	if strings.Contains(req.route, "/"+fileType+"/") {
		fmt.Println("need to intercept")
		route, err := client.httpServ.Find("GET", "/"+fileType)
		if err != nil {
			fmt.Println(err)
			res.SendError(500, "COULD NOT FIND FILE TYPE")
			client.conn.Close()
			return true
		}
		route.function(req, *res)
		return true
	}
	return false
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

		//checks css or js request from html file
		isIntercept := client.handleIntercept(request, response, "css") || client.handleIntercept(request, response, "js") ||
			client.handleIntercept(request, response, "assets")
		if !isIntercept {
			route, err := client.httpServ.Find(request.method, request.route)
			if err != nil {
				fmt.Println(err)
				response.SendError(404, "Could not find Route")
				client.conn.Close()
				return
			}
			route.function(request, *response)
		}

		client.conn.Close()
	}
}
