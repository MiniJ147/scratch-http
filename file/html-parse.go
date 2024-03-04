package file

import (
	"log"
	"os"
)

// must include .html
func ReadFile(name string) string {
	wd, err_ := os.Getwd()
	if err_ != nil {
		log.Println(err_)
		return "<h1>ERROR PARSING</h1>"
	}

	data, err := os.ReadFile(wd + "/html/test.html")
	if err != nil {
		log.Println(err)
		return "<h1>ERROR PARSING</h1>"
	}
	return string(data)
}
