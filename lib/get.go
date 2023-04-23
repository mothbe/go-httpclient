package lib

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Get_text_files(dir string) []string {
	file, err := filepath.Abs(dir)
	fmt.Printf("File: %s\n", file)
	if err != nil {
		log.Printf("[ERROR] Open file: %s", err)
		os.Exit(1)
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	var txt_files []string
	for _, f := range files {
		name := f.Name()
		_, txt := strings.CutSuffix(strings.ToLower(name), ".txt")
		if txt {
			txt_files = append(txt_files, name)
		}
	}
	return txt_files
}
