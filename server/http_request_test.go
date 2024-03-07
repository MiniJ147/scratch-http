package server

import (
	"testing"
)

func TestStatusLine(t *testing.T) {
	const testData1 = "GET /?meep=asdfsad HTTP/1.1 \r\n"

	master := HttpRequest{
		Method:      "GET",
		Route:       "/",
		HttpVersion: "HTTP/1.1",
	}

	tester := createHttpRequest(testData1)

	if master.Method != tester.Method {
		t.Fatalf("TEST 1: FAILED | Method: %v != %v", master.Method, tester.Method)
	}
	if master.Route != tester.Route {
		t.Fatalf("TEST 1: FAILED | Routes: %v != %v", master.Route, tester.Route)
	}
	if master.HttpVersion != tester.HttpVersion {
		t.Fatalf("TEST 1: FAILED | Version: %v != %v", master.HttpVersion, tester.HttpVersion)
	}
}
