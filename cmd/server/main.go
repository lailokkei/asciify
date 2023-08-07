package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/toodemhard/asciify/internal/asciify-lib"
)

var upgrader = websocket.Upgrader{}

func updateImage(message []byte) image.Image {
	img, _, err := image.Decode(bytes.NewReader(message))
	if err != nil {
		log.Println(err)
	}
	return img
}

func connect(w http.ResponseWriter, r *http.Request) {
	img, err := asciify.DecodeImageFile("/home/toodemhard/Pictures/other/1687701362667506.png")
	if err != nil {
		log.Println(err)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		if mt == websocket.TextMessage {
			log.Println("message type : text")
			log.Println(string(message))

			options := asciify.Options{}
			json.Unmarshal(message, &options)
			err := conn.WriteMessage(websocket.TextMessage, []byte(asciify.ImageToText(img, options)))
			if err != nil {
				log.Println(err)
				return
			}
		}

		if mt == websocket.BinaryMessage {
			log.Println("message type : binary")
			img = updateImage(message)
		}
	}
}

func main() {
	port := "8080"
	publicDir := "public"
	fs := http.FileServer(http.Dir(publicDir))

	http.Handle("/", fs)

	http.HandleFunc("/connect", connect)

	fmt.Println("serving on : http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
