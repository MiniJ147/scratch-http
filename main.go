package main

/*
TODO:
 - add support for super long request (split tcp reading)
 - add post and delete
*/

import (
	"fmt"

	"github.com/minij147/scratch-http/server"
)

func main() {
	fmt.Println("Hello From Server")

	type msg struct {
		Message string `json:"Message"`
	}
	serv := server.CreateHttpServer()
	serv.Get("/", func(req *server.HttpRequest, res server.HttpResponse) {
		res.SendJSON(msg{
			Message: "HELLO FROM BASE ROUTE!",
		})
	})

	serv.Get("/html", func(req *server.HttpRequest, res server.HttpResponse) {
		res.SendFile("html/test.html")
	})

	serv.Get("/test", func(req *server.HttpRequest, res server.HttpResponse) {
		v, _ := req.Query.Find("name")
		res.Send(v)
	})

	serv.Patch("/patch", func(req *server.HttpRequest, res server.HttpResponse) {
		res.Send("patch")
	})

	serv.Listen("localhost", "3000")
}
