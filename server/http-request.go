package server

import (
	"fmt"
	"strconv"
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
type Query struct {
	name  string
	value string
	len   int
}

type HttpRequest struct {
	method      string
	route       string
	httpVersion string
	metadata    string
	body        string
	contentSize int
	query       []Query
}

// might refactor later
func parseMethodAndRoute(headerString string) (string, string) {
	requestMethod := ""
	requestRoute := ""
	count := 0
	i := 0

	for {
		currChar := string(headerString[i])

		if count >= 2 {
			break
		}
		if currChar == " " {
			count += 1
			i += 1
			continue
		}

		if count < 1 {
			requestMethod += currChar
		} else {
			requestRoute += currChar
		}

		i += 1
	}

	return requestMethod, requestRoute
}

func parseQuery(req *HttpRequest, queryString string) {
	stage := 0 //0 adding to name | 1: add to value | 2: reset

	currName, currValue := "", ""
	for i := 0; i < len(queryString); i++ {
		currChar := string(queryString[i])
		switch stage {
		case 0:
			if currChar == "=" {
				stage++
				continue
			}

			currName += currChar
		case 1:
			if currChar == "&" {
				req.query = append(req.query, Query{
					name:  currName,
					value: currValue,
				})

				currName = ""
				currValue = ""
				stage = 0

				continue
			}
			currValue += currChar
		}
	}

	//checking if we have to do it one more time
	if currName != "" && currValue != "" {
		req.query = append(req.query, Query{
			name:  currName,
			value: currValue,
		})
	}
}

func parseRequest(req *HttpRequest, requestString string) {
	stage := 0 //0: Method 1: route | 2:query-name | 3: http-version | 4: content-length | 5: metadata | 6: body

	query := Query{
		name:  "",
		value: "",
		len:   0,
	}
	queryStartIndex := -1
	contentLenStr := ""

	for i := 0; i < len(requestString); i++ {
		currentChar := string(requestString[i])
		switch stage {
		case 0:
			if currentChar == " " {
				stage++
				continue
			}

			req.method += currentChar
		case 1:
			if currentChar == " " {
				stage += 2
				continue
			} else if currentChar == "?" {
				queryStartIndex = i + 1
				stage += 1
				continue
			}

			req.route += currentChar
		case 2:
			if currentChar == " " {
				parseQuery(req, requestString[queryStartIndex:queryStartIndex+query.len])
				stage += 1
				continue
			}

			query.len += 1
		case 3:
			if currentChar == "\n" {
				if requestString[i+1:i+16] != "content-length: " {
					//fmt.Println(requestString[i+1 : i+16])
					stage += 1
				} else {
					stage += 2
				}
				continue
			}
			req.httpVersion += currentChar
		case 4:
			if currentChar == "\n" {
				size := len(contentLenStr)
				value, err := strconv.Atoi(contentLenStr[16 : size-1])
				if err != nil {
					fmt.Println(err)
					stage += 1
					continue
				}

				req.contentSize = value
				stage += 1
				continue
			}

			contentLenStr += currentChar
		case 5:
			//TDODO PARSE OUT BODY
			req.body = ""
			req.metadata = requestString[i:]
			stage += 1
		}
		//fmt.Println(string(requestString[i]))
	}
}

func CreateHttpRequest(requestString string) *HttpRequest {
	_, route := parseMethodAndRoute(requestString)
	req := HttpRequest{
		method: "",
		route:  route,
	}

	parseRequest(&req, requestString)
	fmt.Println(req)

	return &req
}
