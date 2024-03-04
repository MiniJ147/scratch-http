# Project Design Doc

## Server

**JOB:**
To manage incoming request and responses. The main idea is for it to direct the binary traffic over to the http handler. From there the http handler will parse the data and execute the proper commands.

## HTTP Handler

**JOB:**
To parse the incoming headers and provide the proper functions that typical http libaries would have. EX: .get, .post, custom routing

### HTTP ROUTING:

Have a map hold the avaible request which will then hold routes mapped to that request. From there routes will have information such as a function. Once we find that function we will run it.

below is example theory code

    struct HTTP-SERVER{
        Map string[REQUEST] -> [ROUTES]
    }

    struct Route{
        method string
        route string
        func *func
    }

    //map["GET"] -> "/","/book"
    avaiableRoutes := map["GET"]
    for i := range(avaiableRoutes){
        if(avaiableRoutes[i]=="REQUESTED_ROUTE"){
            route = avaiableRoutes[i]
            route.func(HTTP_REQUEST, HTTP_RESPONSE)
        }
    }

## rendering html and css
### file dir must look like this
    
    Root Dir/
        views/
            html/
                index.html
            css/
                index.css
        /...    