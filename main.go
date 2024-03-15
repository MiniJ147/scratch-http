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

	serv := server.CreateHttpServer()
	serv.Get("/json", func(req *server.HttpRequest, res server.HttpResponse) {
		type Person struct {
			Name    string   `json:"name"`
			Age     uint8    `json:"age"`
			Friends []string `json:"friends"`
		}

		MyStruct := Person{}
		req.FormatBodyToStruct(&MyStruct)
		res.Send(MyStruct.Name)
	})

	serv.Listen("localhost", "3000")
}
