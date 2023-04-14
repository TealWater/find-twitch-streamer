package main

import (
	control "find-twitch-streamer/controller"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("can't load env file")
	}
}

func main() {

	http.HandleFunc("/", control.GetHomePageHandler)
	http.HandleFunc("/notFound", control.GetNotFoundHandler)
	http.HandleFunc("/findStreamer", control.FindStreamHandler)
	http.HandleFunc("/notFoundRedirect", control.NotFoundRedirectHandler)

	PORT := os.Getenv("PORT")
	fmt.Printf("Starting app on port :" + PORT)
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Fatal(err)
	}
}
