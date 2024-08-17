package art

import (
	"fmt"
	"path/filepath"
	"strings"
)

// CheckFileName helps get a file with the write extension and the file that is present in the directory.
func CheckFileName(fileName string) string {
	if filepath.Ext(fileName) != ".txt" {
		return ""
	}
	if fileName == "standard.txt" || fileName == "shadow.txt" || fileName == "thinkertoy.txt" {
		return fileName
	}
	return ""
}

// AsciiArt gives the art of the given words in the format required.
func AsciiArtOut(input string, inputFile []string) string {
	var result strings.Builder
	var newLinesOnly strings.Builder
	// input = strings.ReplaceAll(input, "\n", "\\n")
	input = strings.ReplaceAll(input, "\r\n", "\\n")
	input = strings.ReplaceAll(input, "\\t", "    ")

	correctInput := CheckInputString(input)
	if !correctInput {
		fmt.Println("combination of required and non required characters")
		return "internal server error"

	}

	sepInputString := strings.Split(input, "\\n")
	newLines := OnlyNewLines(sepInputString)
	if newLines != "false" {
		newLinesOnly.WriteString(newLines)
		return newLinesOnly.String()

	} else {
		for _, words := range sepInputString {
			if words == "" {
				result.WriteString("\n")
			} else {
				for i := 0; i < len(words); {
					for j := 0; j < 8; {
						start := (int(words[i]-32) * 9) + 1
						result.WriteString(inputFile[start+j])
						i++
						if i == len(words) {
							if j == 7 {
								result.WriteString("\n")
								break
							}
							result.WriteString("\n")
							j++
							i = 0

						}
					}
				}
			}
		}
	}
	return result.String()
}

// OnlyNewLines checks if the input string contains only new lines.
func OnlyNewLines(sepInputString []string) string {
	empty := ""
	for i, words := range sepInputString {
		if words != "" {
			return "false"
		}
		if words == "" && i == 0 {
			continue
		}
		empty += "\n"

	}

	return empty
}
