package file

import (
	"log"
	"os"
)

// parses out file with respect from root directory.
func ParseFile(parentDir string, name string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	data, err := os.ReadFile(wd + "/" + parentDir + "/" + name)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(data), nil
}
