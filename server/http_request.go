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
	Body        map[string]any
	Metadata    *HeaderData
	Query       *HeaderData
}

// parses out the first line of the request.
func parseRequestLine(req *HttpRequest, reqStr string) {
	parsed := strings.Split(reqStr, " ") //
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

	log.Println(lines[idx:])
	parseBody(req, lines[idx])

	fmt.Println(req.Body)
}

func parseBody(req *HttpRequest, bodyString string) {
	//parse body if there is one [assuming json]
	if bodyString == "" {
		log.Println("No Body")
		return
	}

	//the curly brace detects if it is a json or not
	if bodyString[0] == '{' {
		err := json.Unmarshal([]byte(bodyString), &req.Body)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		fmt.Println("not vaild json time to parse normally")
		inputs := strings.Split(bodyString, "&")
		for _, input := range inputs {
			vals := strings.Split(input, "=")
			req.Body[vals[0]] = vals[1]
		}
	}
}

// pass in the ptr to the struct you wish the data to enter.
// attempts to format request body based of template passed in.
// will return err if not successful.
func (req *HttpRequest) FormatBodyToStruct(template any) error {
	dbByte, err := json.Marshal(req.Body)
	if err != nil {
		log.Fatalf("Failed to marshal request body: %v\n", err)
		return err
	}
	err = json.Unmarshal(dbByte, &template)
	if err != nil {
		log.Fatalf("Failed to unmarshal json body: %v\n", err)
		return err
	}
	return nil
}

// creates a http request
func createHttpRequest(requestString string) *HttpRequest {
	req := HttpRequest{
		Method:   "",
		Route:    "",
		Body:     make(map[string]any),
		Metadata: createHeaderData(),
		Query:    createHeaderData(),
	}

	parseHeader(&req, requestString)

	return &req
}
