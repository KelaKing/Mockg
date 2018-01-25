package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("Usage mockg DIR_PATH\n\n")
		fmt.Printf("    eg: mockg ./Desktop\n\n")
		return
	}
	root := flag.Arg(0)
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		relativePath, _ := filepath.Rel(root, path)
		if strings.HasSuffix(relativePath, ".json") {
			handleJSON(path, filepath.Join("/", relativePath))
		}
		return nil
	})
	fmt.Println("Start listen on port: 8080...")
	http.ListenAndServe(":8080", nil)
}

func handleJSON(absPath string, relPath string) {
	fmt.Printf("Listen: http://localhost:8080%v\n", relPath)
	http.HandleFunc(relPath, func(writer http.ResponseWriter, request *http.Request) {
		file, error := ioutil.ReadFile(absPath)
		if error != nil {
			fmt.Printf("error: %v\n", error)
		}
		writer.Header().Set("Content-Type", "application/json")
		io.WriteString(writer, string(file))
	})
}
