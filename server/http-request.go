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
	method      string
	route       string
	httpVersion string
	body        interface{}
	metadata    map[string]string
	query       map[string]string
}

func parseRequestLine(req *HttpRequest, reqStr string) {
	parsed := strings.Split(reqStr, " ")
	fmt.Println(parsed)
	method, routeStr, version := parsed[0], parsed[1], parsed[2]

	parsedRoute := strings.Split(routeStr, "?")

	var route, queryLine string

	if len(parsedRoute) > 1 {
		route, queryLine = parsedRoute[0], parsedRoute[1]
		queries := strings.Split(queryLine, "&")

		for _, q := range queries {
			queryParse := strings.Split(q, "=")
			name, value := queryParse[0], queryParse[1]

			req.query[name] = value
		}
	} else {
		route, queryLine = parsedRoute[0], ""
	}

	req.route = route
	req.method = method
	req.httpVersion = version
}

func parseHeader(req *HttpRequest, reqStr string) {
	lines := strings.SplitAfter(reqStr, HEADER_END_LINE)
	length := len(lines)
	parseRequestLine(req, lines[0])

	idx := 1
	for lines[idx] != HEADER_END_LINE && idx < length {
		//fmt.Print(lines[idx])
		c := strings.Split(lines[idx], ":")
		name, value := strings.ToLower(c[0]), c[1][1:] //c[1][1:] parses out value while removing the space

		req.metadata[name] = value
		idx += 1
	}

	idx += 1
	if idx >= length {
		return
	}

	//parse body if there is one
	var body interface{}
	err := json.Unmarshal([]byte(lines[idx]), &body)
	if err != nil {
		log.Println(err)
		return
	}

	req.body = body
}

func CreateHttpRequest(requestString string) *HttpRequest {
	//_, route := parseMethodAndRoute(requestString)
	req := HttpRequest{
		method:   "",
		route:    "",
		body:     nil,
		metadata: make(map[string]string),
		query:    make(map[string]string),
	}

	parseHeader(&req, requestString)

	return &req
}
