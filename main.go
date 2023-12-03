package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"net/http"
	"os"

	"github.com/strukturag/libheif/go/heif"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func convertHandler(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("filename")

	bytes, err := getBytes("./HEIC/" + name)

	if err == nil {
		w.Write(bytes)
		return
	}

	fmt.Println(err)
	w.WriteHeader(500)
}

func getBytes(filename string) ([]byte, error) {
	var out bytes.Buffer

	fmt.Printf("Performing highlevel conversion of %s\n", filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Could not read file %s: %s\n", filename, err)
		return out.Bytes(), err
	}
	defer file.Close()

	img, magic, err := image.Decode(file)
	if err != nil {
		fmt.Printf("Could not decode image: %s\n", err)
		return out.Bytes(), err
	}

	fmt.Printf("Decoded image of type %s: %s\n", magic, img.Bounds())

	if err := png.Encode(&out, img); err != nil {
		fmt.Printf("Could not encode image as PNG: %s\n", err)
		return out.Bytes(), err
	}

	return out.Bytes(), nil
}

func main() {
	fmt.Printf("libheif version: %v\n", heif.GetVersion())
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/convert", convertHandler)
	http.ListenAndServe(":8090", nil)
}
