package server

import (
	"fmt"
	"net"
	"strings"
)

type client struct {
	conn     net.Conn
	httpServ *HttpServer
}

// handles intercept and returns true if intercept happend and false if not.
func (client *client) handleIntercept(req *HttpRequest, res *HttpResponse, fileType string) bool {
	if strings.Contains(req.Route, "/"+fileType+"/") {
		fmt.Println("need to intercept")
		route, err := client.httpServ.find("GET", "/"+fileType)
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

func (client *client) handleCalls(route Route, req *HttpRequest, res HttpResponse) {
	fmt.Println("handling calls")
	fmt.Println(route)
	for _, middleware := range route.middlewares {
		returnFlag := middleware(req, res) //runs the middleware then checks if we need to return out
		if returnFlag {
			return
		}
	}

	route.function(req, res)
}

// handles an incoming request from the tcp server.
func (client *client) handleRequest() {
	buffer := make([]byte, 2048*4) //where request is stored
	for {
		data, err := client.conn.Read(buffer)
		if err != nil {
			//fmt.Println("Read Error: ", err)
			continue
		}

		requestHeaderString := string(buffer[:data])

		//fmt.Println(requestHeaderString)

		request := createHttpRequest(requestHeaderString)

		fmt.Println("client: ", request.Route)
		response := createHttpResponse(client.conn)

		//checks css or js request from html file
		isIntercept := client.handleIntercept(request, response, "css") || client.handleIntercept(request, response, "js") ||
			client.handleIntercept(request, response, "assets")

		if !isIntercept {
			route, err := client.httpServ.find(request.Method, request.Route)
			if err != nil {
				fmt.Println(err)
				response.SendError(404, "Could not find Route")
				client.conn.Close()
				return
			}
			client.handleCalls(route, request, *response)
			//route.function(request, *response)
		}

		client.conn.Close()
	}
}
