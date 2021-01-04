package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	var url string
	fmt.Print("Write url: ")
	fmt.Scanf("%s\n", &url)

	sep_url := strings.Split(url, "/")

	var file_name string
	for i := 1; file_name == "" && i < len(sep_url); i++ {
		file_name = sep_url[len(sep_url)-i]
	}
	if file_name == "" {
		panic("incorrect url")
	}

	err := DownloadFile(file_name, url)
	if err != nil {
		panic(err)
	}
}

func DownloadFile(filepath string, url string) error {

	if !strings.Contains(filepath, ".") {
		filepath += ".html"
	}
	if strings.Contains(filepath, "\\") ||
		strings.Contains(filepath, ":") ||
		strings.Contains(filepath, "*") ||
		strings.Contains(filepath, "?") ||
		strings.Contains(filepath, "\"") ||
		strings.Contains(filepath, "<") ||
		strings.Contains(filepath, ">") ||
		strings.Contains(filepath, "|") {
		fmt.Println("Your URL contains the wrong symbols. File will be saved as temp_file.html")
		filepath = "temp_file.html"
	}

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	var buf bytes.Buffer
	endFile := false
	r := io.TeeReader(resp.Body, &buf)
	fmt.Println("Begin download. Already installed:")
	go func() {
		for !endFile {
			fmt.Println(buf.Len()/1024, "Kb")
			time.Sleep(time.Second)
		}
	}()
	_, err = io.Copy(out, r)
	endFile = true

	fmt.Println(buf.Len()/1024, "Kb")
	fmt.Println("Downloaded as " + filepath)
	return err

	//https://learnxinyminutes.com/docs/ru-ru/go-ru/ - small html file
	//https://i.imgur.com/wgOUusQ.mp4 - mp4 file
	//https://tour.golang.org/static/img/gopher.png - png file
	//https://www.google.com/search?q=Wget&oq=Wget&aqs=chrome..69i57j0l6j69i61.1410j0j7&sourceid=chrome&ie=UTF-8
	// - url with incorrect name for filename
}
