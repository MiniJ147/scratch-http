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

func (serv *HttpServer) Display() {
	for k, v := range serv.methods {
		fmt.Print(k + " ")
		for i := range v {
			fmt.Println(v[i].route)
			v[i].function(*createHttpResponse(nil))
		}
	}
}

func (serv *HttpServer) Get(url_path string, function func(res HttpResponse)) {
	newRoute := Route{
		route:    url_path,
		function: function,
	}

	currentRoutes := serv.methods["GET"]
	currentRoutes = append(currentRoutes, newRoute)

	serv.methods["GET"] = currentRoutes
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
