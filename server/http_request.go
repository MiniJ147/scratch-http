package server

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

/*
GET /?meep=asdfsad HTTP/1.1
content-length: 19
accept-encoding: gzip, deflate, br
Accept:
User-Agent: Thunder Client (https://www.thunderclient.com)
Content-Type: application/json
Host: localhost:3000
Connection: close

	{
	  "msg
*/

type HttpRequest struct {
	Method      string
	Route       string
	HttpVersion string
	Body        interface{}
	Metadata    *HeaderData
	Query       *HeaderData
}

// parses out the first line of the request.
func parseRequestLine(req *HttpRequest, reqStr string) {
	parsed := strings.Split(reqStr, " ")
	//fmt.Println(parsed)
	method, routeStr, version := parsed[0], parsed[1], parsed[2]

	parsedRoute := strings.Split(routeStr, "?")

	var route, queryLine string

	if len(parsedRoute) > 1 {
		route, queryLine = parsedRoute[0], parsedRoute[1]
		queries := strings.Split(queryLine, "&")

		for _, q := range queries {
			queryParse := strings.Split(q, "=")
			name, value := queryParse[0], queryParse[1]

			req.Query.Insert(name, value)
		}
	} else {
		route, queryLine = parsedRoute[0], ""
	}

	req.Route = route
	req.Method = method
	req.HttpVersion = version
}

// parses out the headers infromation and puts it into a http Request.
func parseHeader(req *HttpRequest, reqStr string) {
	lines := strings.Split(reqStr, HEADER_END_LINE)
	length := len(lines)
	parseRequestLine(req, lines[0])

	idx := 1
	for lines[idx] != HEADER_END_LINE && idx < length-1 {
		//fmt.Println(lines[idx])
		c := strings.Split(lines[idx], ":")
		if len(c) > 1 {
			name, value := strings.ToLower(c[0]), c[1][1:]
			req.Metadata.Insert(name, value)
		}
		idx += 1
	}

	if idx >= length {
		fmt.Println("out of bounds")
		return
	}

	//parse body if there is one
	var body interface{}
	err := json.Unmarshal([]byte(lines[idx]), &body)
	if err != nil {
		log.Println(err)
		return
	}

	req.Body = body
	//fmt.Println(body)
}

// creates a http request
func createHttpRequest(requestString string) *HttpRequest {
	req := HttpRequest{
		Method:   "",
		Route:    "",
		Body:     nil,
		Metadata: createHeaderData(),
		Query:    createHeaderData(),
	}

	parseHeader(&req, requestString)

	return &req
}
