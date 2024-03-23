package main

/*
TODO:
 - add support for super long request (split tcp reading)
 - make sure all types of bodys are accepted
 - simplfy json parsing
 - middleware support
 - support for routers
 - cookies
*/
//RandyGG

import (
	"fmt"

	"github.com/minij147/scratch-http/server"
)

func exampleMiddleware(req *server.HttpRequest, res server.HttpResponse) bool {
	fmt.Println("From Middleware")
	return false
}

func main() {
	fmt.Println("Hello From Server")

	serv := server.CreateHttpServer()
	serv.Get("/", func(req *server.HttpRequest, res server.HttpResponse) {
		res.SendFile("html/test.html")
	})

	serv.Listen("localhost", "3000")
}
