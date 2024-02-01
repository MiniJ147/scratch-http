package server

import (
	"fmt"
	"net"
	"strconv"
)

const HEADER_END_LINE string = "\r\n"

type HttpResponse struct {
	statusLine string
	header     map[string]string
	conn       net.Conn
}

// write things to header
func (r *HttpResponse) writeHeader(key string, value string) {
	r.header[key] = value
}

// write status to response
func (r *HttpResponse) writeStatus(code int, msg string) {
	r.statusLine = "HTTP/1.1 " + strconv.Itoa(code) + " " + msg + HEADER_END_LINE
}

func (r *HttpResponse) compileHeader() []byte {
	headerString := ""
	headerString += r.statusLine

	for k, v := range r.header {
		headerString += k + ": " + v + HEADER_END_LINE
	}

	headerString += HEADER_END_LINE

	return []byte(headerString)
}

func (r *HttpResponse) compilePayload(data string) []byte {
	return []byte(data)
}

func (r *HttpResponse) send(payload string) {
	headerData := r.compileHeader()
	payLoadData := r.compilePayload(payload)

	fmt.Println(string(headerData))
	fmt.Println(string(payLoadData))

	r.conn.Write(headerData)
	r.conn.Write(payLoadData)
}

func createHttpResponse(conn net.Conn) *HttpResponse {
	return &HttpResponse{
		statusLine: "",
		header:     make(map[string]string),
		conn:       conn,
	}
}
