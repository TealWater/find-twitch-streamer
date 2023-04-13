package main

import (
	"fmt"
	"log"
	"net/http"
	control "twitch_auth/webServ/controller"
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
