package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/rupinjairaj/snippet/entity"
)

var url string

func getFileContentBase64Encoded(fileName string) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Printf("Failed to open the file: %v\n", fileName)
		return "", err
	}

	reader := bufio.NewReader(f)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Printf("Failed to read contents of file: %v\n", fileName)
		return "", nil
	}

	encoded := base64.StdEncoding.EncodeToString(content)
	return encoded, nil
}

func upload(data *entity.SnippetClient) {
	method := "POST"

	reqByte, err := json.Marshal(data)
	if err != nil {
		log.Printf("%v", err)
		return
	}

	payload := bytes.NewReader(reqByte)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Printf("%v", err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("%v", err)
		return
	}

}

func main() {

	baseUrl := os.Getenv("BASE_URL")
	snippetUrl := os.Getenv("SNIPPET_URL")

	if baseUrl == "" || snippetUrl == "" {
		log.Println("Environment variables 'baseUrl' and 'snippetUrl' are not set")
		return
	}

	url = baseUrl + snippetUrl

	// publish snippet command set
	publishCommand := flag.NewFlagSet("publish", flag.ExitOnError)

	filePtr := publishCommand.String("file", "", "File to publish.")
	tagsPtr := publishCommand.String("tags", "", "Space separated tags in quotes. Eg: -tags \"tag1 tag2 tag3\".")

	if len(os.Args) < 2 {
		fmt.Println("Insufficient arguments passed")
		publishCommand.PrintDefaults()
		return
	}

	switch os.Args[1] {
	case "publish":
		publishCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		return
	}

	if *filePtr == "" {
		publishCommand.PrintDefaults()
		return
	}

	if *tagsPtr == "" {
		publishCommand.PrintDefaults()
		return
	}

	// read input file
	log.Printf("Input file: %s \nTaget topic: %v\n", *filePtr, *&tagsPtr)
	data, err := getFileContentBase64Encoded(*filePtr)
	if err != nil {
		log.Println("Failed to get file content base64 encoded.")
		return
	}

	// send the data to the api
	uploadData := entity.SnippetClient{
		Content: data,
		Tags:    strings.Fields(*tagsPtr),
		Name:    *filePtr,
	}

	upload(&uploadData)
}
