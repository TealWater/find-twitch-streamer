package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	twitch "twitch_auth/webServ/utils"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("can't load env file")
	}
}

var tpl *template.Template

func GetHomePageHandler(w http.ResponseWriter, r *http.Request) {
	tpl, _ = tpl.ParseFiles("views/home.html")
	tpl.ExecuteTemplate(w, "home.html", nil)
}

func FindStreamHandler(w http.ResponseWriter, r *http.Request) {
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
	twitchUser := new(twitch.TwitchUsers)
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
}
