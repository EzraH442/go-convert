package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"net/http"

	"github.com/strukturag/libheif/go/heif"
)

func convertHandler(w http.ResponseWriter, req *http.Request) {
	image, magic, err := image.Decode(req.Body)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}

	bytes, err := convertToPng(image, magic)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}

	w.Write(bytes)
	return
}

func convertToPng(img image.Image, magic string) ([]byte, error) {
	var out bytes.Buffer

	fmt.Printf("Decoded image of type %s: %s\n", magic, img.Bounds())

	if err := png.Encode(&out, img); err != nil {
		fmt.Printf("Could not encode image as PNG: %s\n", err)
		return out.Bytes(), err
	}

	return out.Bytes(), nil
}

func main() {
	fmt.Printf("libheif version: %v\n", heif.GetVersion())
	http.HandleFunc("/convert", convertHandler)
	http.ListenAndServe(":8090", nil)
}
