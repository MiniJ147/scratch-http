package server

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
		if routes[i].route == route {
			return routes[i]
		}
	}

	//TODO ADD ERROR HANDLING
	return routes[0]
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
