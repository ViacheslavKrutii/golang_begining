package textRedactor

import (
	"flag"
	"fmt"
	"os"
	"regexp"
)

var txtFilename *string

func openTxtFile() *os.File {

	txtFilename = flag.String("txt", "example.txt", "	a txt file")
	flag.Parse()

	file, err := os.OpenFile(*txtFilename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		exit(fmt.Sprintf("Failed to open txt file: %s\n", *txtFilename))
	}
	return file
}

func writeTxtToFile(file *os.File) *os.File {
	var inputText string

	fmt.Println("type string for writing")
	fmt.Scan(&inputText)

	inputText += "\n"

	_, err := file.WriteString(inputText)
	if err != nil {
		exit(fmt.Sprintf("Failed to write txt file: %s\n", *txtFilename))
	}

	return file

}

func searchTxt(file *os.File) {
	var inputSearchedText string
	fmt.Println("type searched text")
	fmt.Scan(&inputSearchedText)

	content, err := os.ReadFile(*txtFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to read txt file: %s\n", *txtFilename))

	}

	var matches []string
	matches = searchLines(string(content), inputSearchedText)

	if len(matches) == 0 {
		exit(fmt.Sprintf("String is not found: %s\n", *txtFilename))
	}

	fmt.Printf("Matches: %v\n", matches)

}

func searchLines(text, searchText string) []string {
	re := regexp.MustCompile(`(?m).*` + regexp.QuoteMeta(searchText) + `.*`)
	matches := re.FindAllString(text, -1)
	return matches
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func TextRedactor() {
	file := openTxtFile()
	defer file.Close()
	writeTxtToFile(file)
	searchTxt(file)

}
