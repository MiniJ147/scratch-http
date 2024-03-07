package server

import (
	"fmt"
	"log"
	"strings"
)

type Route struct {
	route    string
	function func(req *HttpRequest, res HttpResponse)
}

type HttpServer struct {
	tcpServer *Server
	methods   map[string][]Route
}

func (serv *HttpServer) Find(method string, route string) (Route, error) {
	routes := serv.methods[method]

	for i := range routes {
		//fmt.Println(i)
		//fmt.Println(routes[i].route, route)
		if routes[i].route == route {
			fmt.Println("Found")
			return routes[i], nil
		}
	}

	fmt.Println("error")
	return routes[0], fmt.Errorf("could not find route")
}

// method: GET, POST, PATCH, DELETE ...
func helperCreateMethod(serv *HttpServer, method string, urlPath string, function func(req *HttpRequest, res HttpResponse)) {
	newRoute := Route{
		route:    urlPath,
		function: function,
	}

	currentRoutes := serv.methods[method]
	currentRoutes = append(currentRoutes, newRoute)

	serv.methods[method] = currentRoutes
}

func (serv *HttpServer) Get(urlPath string, function func(req *HttpRequest, res HttpResponse)) {
	helperCreateMethod(serv, "GET", urlPath, function)
}

func (serv *HttpServer) Post(urlPath string, function func(req *HttpRequest, res HttpResponse)) {
	helperCreateMethod(serv, "POST", urlPath, function)
}

func (serv *HttpServer) Listen(host string, port string) {
	serv.tcpServer = CreateServer(host, port)
	serv.tcpServer.Run(serv)
}

func CreateHttpServer() *HttpServer {
	serv := &HttpServer{
		methods: make(map[string][]Route),
	}

	//important routes that are standard
	serv.Get("/css", func(req *HttpRequest, res HttpResponse) {
		fmt.Println("csss: " + req.route)
		fileName := strings.Join(strings.SplitAfter(req.route, "/")[2:], "")
		res.WriteHeader("Content-Type", "text/css")
		res.SendFile("css/" + fileName)
	})

	serv.Get("/js", func(req *HttpRequest, res HttpResponse) {
		fmt.Println("JS: " + req.route)
		fileName := strings.Join(strings.SplitAfter(req.route, "/")[2:], "")
		log.Println(fileName)
		res.WriteHeader("Content-Type", "text/js")
		res.SendFile("js/" + fileName)
	})
	return serv
}
