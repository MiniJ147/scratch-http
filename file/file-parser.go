package file

import (
	"log"
	"os"
)

// with respect from root
func ParseFile(name string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	data, err := os.ReadFile(wd + "/views/" + name)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(data), nil
}
