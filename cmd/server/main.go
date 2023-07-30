package main

import (
	"fmt"
	"image"
	"log"
	"net/http"

	"github.com/toodemhard/asciify/internal/asciify-lib"
)

func main() {
	port := "8080"
	publicDir := "public"
	fs := http.FileServer(http.Dir(publicDir))

	http.Handle("/", fs)

	imgDefault, err := asciify.DecodeImageFile("/home/toodemhard/Pictures/other/1687701362667506.png")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {

		options := asciify.Options{CharSetName: "standard", Invert: false, ScaleHeight: 20}
		fmt.Fprint(w, asciify.ImageToText(imgDefault, options))
	})

	http.HandleFunc("/api/image", func(w http.ResponseWriter, r *http.Request) {
		img, _, err := image.Decode(r.Body)
		if err != nil {
			return
		}
		options := asciify.Options{CharSetName: "simple", Invert: false, ScaleHeight: 20}
		text := asciify.ImageToText(img, options)
		fmt.Fprint(w, text)
		fmt.Print(text)
	})

	fmt.Println("serving on : http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
