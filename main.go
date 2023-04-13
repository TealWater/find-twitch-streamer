package main

import (
	control "find-twitch-streamer/controller"
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/home", control.GetHomePageHandler)
	http.HandleFunc("/notFound", control.GetNotFoundHandler)
	http.HandleFunc("/findStreamer", control.FindStreamHandler)
	http.HandleFunc("/notFoundRedirect", control.NotFoundRedirectHandler)

	fmt.Printf("Starting server at port 5500\n")
	if err := http.ListenAndServe(":5500", nil); err != nil {
		log.Fatal(err)
	}
}
