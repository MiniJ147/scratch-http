# READ ME

### important
    this libary is currently not finished and still being developed.
    if you have any errors please report to make this libary better.

## About

    Simple Http libary made in go-lang from scratch. As of 3/4/24 it supports:
        routing,
        GET, POST, DELETE, PATCH, PUT
        JSON request
        Full header access
        html rendering
        css rendering
        js support
        Image support
    
    The goal of this project is to be useable in an actual application. 

    Future Plans:
        99% success rate
        hot reloading 
        unit test
        [somewhat done] Easy json responses
        Large request acceptance 
        Support for large raw data streaming
        [somewhat done] Easy searching of request & responses
        and more... 


## installation
```bash
go get github.com/minij147/scratch-http/
```
## Demo
```go
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
```
## Creating Server / running
```go
initilizes the http server. This will be used to communicate with http.

//returns *server.HttpServer{}
serv := server.CreateHttpServer() 

//turns on server to listen on ip and port
serv.Listen(host, port) 
```

## Method Functions
```go
//this is how to create routes assigned to specfic methods. 
//Func will be called when route is called.
    
serv.Get("route", func, middleware ...func)

serv.Post("route", func, middleware ...func)

serv.Delete("route",func, middleware ...func)

serv.Put("route",func, middleware ...func)

serv.Patch("route",func, middleware ...func)
```

## Response
```go
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

//sends error to the browser with code and message
httpResponse.SendError(code int, msg string)

//sets header with the key and gives it the value
httpResponse.WriteHeader("key", "value")

//writes response code to the browser
httpResponse.WriteStatus(code int, msg string)

//will redirect to a new route after function calling
httpResponse.Redirect("route")
```
## Request
```go
//requests data will automatically be prased
//this is how to access

type HttpRequest struct {
	Method      string
	Route       string
	HttpVersion string
	Body        interface{}
	Metadata    *HeaderData
	Query       *HeaderData
}

```
## Accessing Query or Metadata using HeaderData
```go
type HeaderData struct {
	data map[string]string
}

//value and if it passed 
//up to user to handle bool
//searches map for you 
//returns "" and False if not found
//returns value and True if found
func Find("key") (string, bool)

//inserts into map based off value to set
func Insert("key", "value")
```

## Middleware
```go
/*
Important note
middleware will return a bool this is the status of the function call.
False -> means no redirect we can continue to next middleware
True -> means redirect and we must stop the middleware and all future calls

Middleware gets executed in order 0 --> n
*/

//To declare
func exampleMiddleware(req *server.HttpRequest, res server.HttpResponse) bool {
	fmt.Println("From Middleware")
	return false
}

//passing it in
serv.Get("/",func, exampleMiddleware)

//example flow of middleware
//the flow would look like
//mid1 -> mid2 -> mid3 -> logic
serv.Get("/",logic func, mid1 func, mid2, func, mid3 func)
```

## Rendering html 
```go
Root/..
    views/
        assets/
            .png, .svg, ...
        css/
            index.css
        html/
            index.html
        js/
            index.js
    ...

 httpResponse.SendFile("html/{filename}.html")
```