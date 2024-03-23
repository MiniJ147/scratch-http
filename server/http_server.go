package server

import (
	"fmt"
	"strings"
)

type Route struct {
	route       string
	middlewares []func(req *HttpRequest, res HttpResponse) bool
	function    func(req *HttpRequest, res HttpResponse)
}

type HttpServer struct {
	tcpServer *server
	methods   map[string][]Route
}

// search for a given route
func (serv *HttpServer) find(method string, route string) (Route, error) {
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
// helper function to cut down on repeated code.
// creates a method for us.
func helperCreateMethod(serv *HttpServer, method string, urlPath string, function func(req *HttpRequest, res HttpResponse),
	middlewares []func(req *HttpRequest, res HttpResponse) bool) {
	newRoute := Route{
		route:       urlPath,
		middlewares: middlewares,
		function:    function,
	}

	fmt.Println(middlewares)

	currentRoutes := serv.methods[method]
	currentRoutes = append(currentRoutes, newRoute)

	serv.methods[method] = currentRoutes
}

// create a get request
func (serv *HttpServer) Get(urlPath string, function func(req *HttpRequest, res HttpResponse),
	middlewares ...func(req *HttpRequest, res HttpResponse) bool) {
	helperCreateMethod(serv, "GET", urlPath, function, middlewares)
}

// create a post request
func (serv *HttpServer) Post(urlPath string, function func(req *HttpRequest, res HttpResponse),
	middlewares ...func(req *HttpRequest, res HttpResponse) bool) {
	helperCreateMethod(serv, "POST", urlPath, function, middlewares)
}

// create a delete request
func (serv *HttpServer) Delete(urlPath string, function func(req *HttpRequest, res HttpResponse),
	middlewares ...func(req *HttpRequest, res HttpResponse) bool) {
	helperCreateMethod(serv, "DELETE", urlPath, function, middlewares)
}

// create a PATCH request
func (serv *HttpServer) Patch(urlPath string, function func(req *HttpRequest, res HttpResponse),
	middlewares ...func(req *HttpRequest, res HttpResponse) bool) {
	helperCreateMethod(serv, "PATCH", urlPath, function, middlewares)
}

// create a PUT request
func (serv *HttpServer) Put(urlPath string, function func(req *HttpRequest, res HttpResponse),
	middlewares ...func(req *HttpRequest, res HttpResponse) bool) {
	helperCreateMethod(serv, "PUT", urlPath, function, middlewares)
}

// opens server on the port so it can take connections.
func (serv *HttpServer) Listen(host string, port string) {
	serv.tcpServer = createServer(host, port)
	serv.tcpServer.run(serv)
}

// creates standard route and plugs in the repeated data.
func helperCreateStandardRoute(fileType string, accepts string) func(req *HttpRequest, res HttpResponse) {
	return func(req *HttpRequest, res HttpResponse) {
		fileName := strings.Join(strings.SplitAfter(req.Route, "/")[2:], "")
		res.WriteHeader("Content-Type", accepts)
		res.SendFile(fileType + "/" + fileName)
	}
}

// creates http server that you can use to communicate with http.
func CreateHttpServer() *HttpServer {
	serv := &HttpServer{
		methods: make(map[string][]Route),
	}

	//important routes that are standard.
	serv.Get("/css", helperCreateStandardRoute("css", "text/css"))
	serv.Get("/js", helperCreateStandardRoute("js", "text/js"))
	serv.Get("/assets", helperCreateStandardRoute("assets", "*/*"))

	return serv
}
