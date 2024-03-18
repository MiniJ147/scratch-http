package main

/*
TODO:
 - add support for super long request (split tcp reading)
 - make sure all types of bodys are accepted
 - simplfy json parsing
 - Redirect
 - middleware support
*/

import (
	"fmt"

	"github.com/minij147/scratch-http/server"
)

func main() {
	fmt.Println("Hello From Server")

	serv := server.CreateHttpServer()
	serv.Get("/", func(req *server.HttpRequest, res server.HttpResponse) {
		res.SendFile("html/test.html")
	})

	serv.Post("/submit", func(req *server.HttpRequest, res server.HttpResponse) {
		type example struct {
			Price string `json:"price"`
			Name  string `json:"name"`
		}

		e := example{}
		req.FormatBodyToStruct(&e)
		fmt.Printf("Name: %v | Price: %v \n", e.Name, e.Price)
		res.Send("Hey")
	})

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
