package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"
)

var (
	filePath string
	addr     string
)

func init() {
	flag.StringVar(&filePath, "file", "", "the file to upload")
	flag.StringVar(&addr, "addr", "localhost:8080", "the addr of file server")
	flag.Parse()
}

func main() {
	if filePath == "" {
		fmt.Println("file must not be empty")
		return
	}

	err := doUpload(addr, filePath)
	if err != nil {
		fmt.Printf("upload file [%s] error: %s", filePath, err)
		return
	}
	fmt.Printf("upload file [%s] ok\n", filePath)
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

func createReqBody(filePath string) (string, io.Reader, error) {
	var err error
	pr, pw := io.Pipe()
	bw := multipart.NewWriter(pw) // body writer
	f, err := os.Open(filePath)
	if err != nil {
		return "", nil, err
	}

	go func() {
		defer f.Close()
		// text part1
		p1w, _ := bw.CreateFormField("name")
		p1w.Write([]byte("Tony Bai"))

		// text part2
		p2w, _ := bw.CreateFormField("age")
		p2w.Write([]byte("15"))

		// file part1
		_, fileName := filepath.Split(filePath)
		h := make(textproto.MIMEHeader)
		h.Set("Content-Disposition",
			fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
				escapeQuotes("file1"), escapeQuotes(fileName)))
		h.Set("Content-Type", "application/pdf")
		fw1, _ := bw.CreatePart(h)
		cnt, _ := io.Copy(fw1, f)
		log.Printf("copy %d bytes from file %s in total\n", cnt, fileName)
		bw.Close() //write the tail boundry
		pw.Close()
	}()
	return bw.FormDataContentType(), pr, nil
}

func doUpload(addr, filePath string) error {
	// create body
	contType, reader, err := createReqBody(filePath)
	if err != nil {
		return err
	}

	log.Printf("createReqBody ok\n")
	url := fmt.Sprintf("http://%s/upload", addr)
	req, err := http.NewRequest("POST", url, reader)

	//add headers
	req.Header.Add("Content-Type", contType)

	client := &http.Client{}
	log.Printf("upload %s...\n", filePath)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("request send error:", err)
		return err
	}
	resp.Body.Close()
	log.Printf("upload %s ok\n", filePath)
	return nil
}
