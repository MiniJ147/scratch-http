# READ ME

### important
    this libary is currently not finished and still being developed.
    if you have any errors please report to make this libary better.

## About

    Simple Http libary made in go-lang from scratch. As of 3/4/24 it supports:
        routing,
        GET, POST
        JSON request
        Full header access
        html rendering
        css rendering
    
    The goal of this project is to be useable in an actual application. 

    Future Plans:
        99% success rate 
        JS script support
        Easy json responses
        Large request acceptance 
        Support for large raw data streaming
        Easy searching of request & responses
        and more... 


## installation

    go get github.com/minij147/scratch-http/

## Demo

    package main

    import "github.com/minij147/scratch-http/server"

    func main(){
        type msg struct {
		    Message string `json:"Message"`
	    }

        serv := server.CreateHttpServer()

        serv.Get("/", func(req *server.HttpRequest, res server.HttpResponse) {
            res.WriteStatus(200, "OK")
            res.SendJSON(msg{
                Message: "HELLO FROM BASE ROUTE!",
            })
        })

        serv.Listen("localhost", "3000")
    }

## Creating Server / running
    
    initilizes the http server. This will be used to communicate with http.

    //returns *server.HttpServer{}
    serv := server.CreateHttpServer() 

    //turns on server to listen on ip and port
    serv.Listen(host, port) 

## Method Functions

    //this is how to create routes assigned to specfic methods. 
    //Func will be called when route is called.
    
    serv.Get("route", func)

    serv.Post("route", func)


## Response
    type HttpResponse struct {
        statusLine string  
        header     map[string]string
        conn       net.Conn
    }   

    //return parsed file
    httpResponse.SendFile(filename)

    //sends string to browser
    httpResponse.Send(string)

    //sends json to browser
    httpResponse.SendJSON(interface{})

    //sets header with the key and gives it the value
    httpResponse.WriteHeader("key", "value")

    //writes response code to the browser
    httpResponse.WriteStatus(code int, msg string)

## Request

    //requests data will automatically be prased
    //this is how to access

    type HttpRequest struct {
        method      string
        route       string
        httpVersion string
        body        interface{}
        metadata    map[string]string
        query       map[string]string
    }

## Rendering html 

    Root/..
        views/
            css/
                index.css
            html/
                index.html
        ...

    httpResponse.SendFile("html/{filename}.html")