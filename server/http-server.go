package server

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"time"
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

func (r *HttpResponse) initHeader() {
	r.writeHeader("Content-Type", "text/html")
	r.writeHeader("Date", time.Now().Format("01-02-2006 15:04:05"))
	r.writeHeader("Connection", "Keep-Alive")
	r.writeHeader("Keep-Alive", "timeout=5")
}

func (r *HttpResponse) sendJSON(payload interface{}) {
	payLoadData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("ERROR")
		payLoadData = []byte("<h1>ERROR</h1")
	}

	r.writeHeader("Content-Type", "application/json")
	r.writeHeader("Content-Length", strconv.Itoa(len(payLoadData)))

	headerData := r.compileHeader()

	fmt.Println(string(headerData))

	r.conn.Write(headerData)
	r.conn.Write(payLoadData)
}

func (r *HttpResponse) send(payload string) {
	payLoadData := r.compilePayload(payload)

	r.writeHeader("Content-Length", strconv.Itoa(len(payLoadData)))

	headerData := r.compileHeader()

	fmt.Println(string(headerData))
	fmt.Println(string(payLoadData))

	r.conn.Write(headerData)
	r.conn.Write(payLoadData)
}

func createHttpResponse(conn net.Conn) *HttpResponse {
	r := HttpResponse{
		statusLine: "",
		header:     make(map[string]string),
		conn:       conn,
	}

	r.initHeader()
	return &r
}
