package lib

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Read_line_by_line(path string) map[string]int {

	filePath := path
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	domain := strings.Trim(strings.ToLower(filePath), ".txt")
	urls := make(map[string]int)
	for _, line := range fileLines {
		line_arr := strings.Split(line, " ")
		status_code, err := strconv.Atoi(line_arr[1])

		if err != nil {
			log.Fatalf("Error during conversion\n")
		}
		url := "https://" + domain + line_arr[0]
		urls[url] = status_code
	}
	return urls
}
