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

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
