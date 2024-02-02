package main

/*
TODO:
 - add request support
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

		res.WriteStatus(200, "OK")
		res.SendJSON(msg{
			Message: "HELLO FROM BOOK!",
		})
	})

	serv.Get("/html", func(res server.HttpResponse) {
		res.WriteStatus(200, "OK")
		res.Send("<h1>JOHN</h1>")
	})

	serv.Listen("localhost", "3000")
}
