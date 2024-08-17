package art

import (
	"fmt"
	"os"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
func fileNotFound(fileName string) string {
	_, err := os.Stat("./art/" + fileName)
	if err == nil {
		return "true" // File exists
	}
	if os.IsNotExist(err) {
		return "false" // File does not exist
	}
	// Some other error occurred
	return "false"

}

// ReadFromFile gets the filename and splits it into an array of strings.
func ReadFromFile(fileName string) ([]string, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	if string(file) == "" {
		return nil, nil
	}

	var lines []string
	if fileName == "./art/thinkertoy.txt" {
		lines = strings.Split(string(file), "\r\n")
	} else {
		lines = strings.Split(string(file), "\n")
	}

	return lines, nil
}
