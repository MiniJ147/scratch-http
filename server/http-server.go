package server

import "fmt"

type Route struct {
	route    string
	function func(res HttpResponse)
}

type HttpServer struct {
	tcpServer *Server
	methods   map[string][]Route
}

func (serv *HttpServer) Find(method string, route string) Route {
	routes := serv.methods[method]

	for i := range routes {
		fmt.Println(i)
		fmt.Println(routes[i].route, route)
		if routes[i].route == route {
			fmt.Println("Found")
			return routes[i]
		}
	}

	//TODO ADD ERROR HANDLING
	fmt.Println("error")
	fmt.Println(serv.methods, method, route)
	return routes[0]
}

// method: GET, POST, PATCH, DELETE ...
func helperCreateMethod(serv *HttpServer, method string, urlPath string, function func(res HttpResponse)) {
	newRoute := Route{
		route:    urlPath,
		function: function,
	}

	currentRoutes := serv.methods[method]
	currentRoutes = append(currentRoutes, newRoute)

	serv.methods[method] = currentRoutes
}

func (serv *HttpServer) Get(urlPath string, function func(res HttpResponse)) {
	helperCreateMethod(serv, "GET", urlPath, function)
}

func (serv *HttpServer) Post(urlPath string, function func(res HttpResponse)) {
	helperCreateMethod(serv, "POST", urlPath, function)
}

func (serv *HttpServer) Listen(host string, port string) {
	serv.tcpServer = CreateServer(host, port)
	serv.tcpServer.Run(serv)
}

func CreateHttpServer() *HttpServer {
	return &HttpServer{
		methods: make(map[string][]Route),
	}
}
