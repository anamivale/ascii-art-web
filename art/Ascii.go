package art

// checkErr helps in Printing errors that occur when you use some package methods.

func AsciiArt(input string, file string) string {
	var file1 string


	if file == "standard" {
		file1 = "standard.txt"
	} else if file == "thinkertoy" {
		file1 = "thinkertoy.txt"
	} else if file == "shadow" {
		file1 = "shadow.txt"
	}

	fileExists := fileNotFound(file1)
	if fileExists == "false" {
		return "Not Found"
		
	}

	// check valid file extension and file name.
	validFileNameAndExtension := CheckFileName(file1)
	if validFileNameAndExtension == "" {
		return "Not Found"
	}
	// stores data in slices of lines
	lines, err := ReadFromFile("./art/" + validFileNameAndExtension)
	if lines == nil {
		return "internal server error"
	}
	if err != nil {
		return "internal server error"

	}

	// inpuString is the string that needs to be converted to an art.
	inpuString := input

	// call Trial function to implement its functioality.
	result := AsciiArtOut(inpuString, lines)
	return result

}
