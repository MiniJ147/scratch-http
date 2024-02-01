package main

import (
	"fmt"

	"github.com/minij147/scratch-http/server"
)

func main() {
	fmt.Println("Hello From Server")

	serv := server.CreateServer("localhost", "3000")
	serv.Run()
}
