package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type TwitchUsers struct {
	Data []struct {
		ID           string    `json:"id"`
		UserID       string    `json:"user_id"`
		UserLogin    string    `json:"user_login"`
		UserName     string    `json:"user_name"`
		GameID       string    `json:"game_id"`
		GameName     string    `json:"game_name"`
		Type         string    `json:"type"`
		Title        string    `json:"title"`
		ViewerCount  int       `json:"viewer_count"`
		StartedAt    time.Time `json:"started_at"`
		Language     string    `json:"language"`
		ThumbnailURL string    `json:"thumbnail_url"`
		TagIds       []any     `json:"tag_ids"`
		Tags         []string  `json:"tags"`
		IsMature     bool      `json:"is_mature"`
	} `json:"data"`
	Pagination struct {
	} `json:"pagination"`
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("can't load env file")
	}

	http.HandleFunc("/hell", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {

		client := http.Client{}
		url := "https://api.twitch.tv/helix/streams?user_login=frslushh&user_login=clix&user_login=pewdiepie"
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			//Handle Error
			log.Println("here")
			log.Println(err)
		}

		req.Header = http.Header{
			"Authorization": {os.Getenv("BEARER_TOKEN")},
			"Client-Id":     {os.Getenv("CLIENT_ID")},
		}

		res, err := client.Do(req)
		if err != nil {
			//Handle Error
			log.Println(err)
		}

		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		log.Println(bodyString)
		//fmt.Fprintf(w, bodyString)
		fmt.Println("hi")

		/*Parsing out twitch streamer username*/
		twitchUser := TwitchUsers{}
		err = json.Unmarshal(bodyBytes, &twitchUser)
		if err != nil {
			log.Println()
			log.Println("can't parse json")
			log.Fatal(err)
		}

		// for index := range twitchUser.Data {
		// 	fmt.Fprintf(w, twitchUser.Data[index].UserLogin+"\n")
		// 	fmt.Fprintln(w, "length: "+strconv.Itoa(len(twitchUser.Data)))
		// }

		length := len(twitchUser.Data)

		if length > 0 {
			userName := twitchUser.Data[0].UserLogin
			url := "https://www.twitch.tv/" + userName

			http.Redirect(w, r, url, http.StatusTemporaryRedirect)
		} else {
			url := "http://127.0.0.1:5500/hell"

			http.Redirect(w, r, url, http.StatusTemporaryRedirect)

		}

	})

	fmt.Printf("Starting server at port 5500\n")
	if err := http.ListenAndServe(":5500", nil); err != nil {
		log.Fatal(err)
	}
}
