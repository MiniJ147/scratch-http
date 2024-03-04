package main

/*
TODO:
 - add support for super long request (split tcp reading)
 - add parsing queries and body support
 - add error handling
 - add post and delete
 - add html file support
 - add css support
 - add js support
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
	serv.Get("/", func(res server.HttpResponse) {

		res.WriteStatus(200, "OK")
		res.SendJSON(msg{
			Message: "HELLO FROM BASE ROUTE!",
		})
	})

	serv.Get("/book", func(res server.HttpResponse) {
		fmt.Println("book route")
		res.WriteStatus(200, "OK")
		res.SendJSON(msg{
			Message: "HELLO FROM BOOK!",
		})
	})

	serv.Get("/html", func(res server.HttpResponse) {
		res.WriteStatus(200, "OK")
		//res.Send("<h1>JOHN</h1>")
		res.SendFile("test.html")
	})

	serv.Listen("localhost", "3000")
}
