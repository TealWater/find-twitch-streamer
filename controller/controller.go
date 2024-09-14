package controller

import (
	"encoding/json"
	twitch "find-twitch-streamer/utils"
	"html/template"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"

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
	tpl, _ = tpl.ParseGlob("views/*.html")
	tpl.ExecuteTemplate(w, "home.html", nil)
}

func GetNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	tpl, _ = tpl.ParseGlob("views/*.html")
	tpl.ExecuteTemplate(w, "notFound.html", nil)
}

func GetRandomStreamer(w http.ResponseWriter, r *http.Request) (finalUrl string) {
	url := "https://api.twitch.tv/helix/streams?"

	bodyBytes := hitTwitchApi(w, *r, url, nil)

	/*Parsing out twitch streamer username*/
	twitchUser := new(twitch.TwitchUsers)
	err := json.Unmarshal(bodyBytes, &twitchUser)
	if err != nil {
		log.Println()
		log.Println("can't parse json")
		log.Fatal(err)
	}

	//get the length of the names and choose one at random
	length := len(twitchUser.Data)
	val := rand.Intn(length)
	url = "https://www.twitch.tv/" + twitchUser.Data[val].UserLogin
	return url
}

func NotFoundRedirectHandler(w http.ResponseWriter, r *http.Request) {
	var err error = nil
	if err = r.ParseForm(); err != nil {
		log.Println("can't parse form")
		log.Fatal(err)
	}

	tpl, _ = tpl.ParseGlob("views/*.html")
	for key := range r.Form {
		/*can't use <http.Redirect()> here beacuse <http.ResponseWriter> was already in use from <func FindStreamHandler()>*/
		if strings.Compare(key, "backToHome") == 0 {
			tpl.ExecuteTemplate(w, "home.html", nil)
		} else if strings.Compare(key, "randomStream") == 0 {
			tpl.ExecuteTemplate(w, "randomStreamer.html", nil)
		}

	}

}

func FindStreamHandler(w http.ResponseWriter, r *http.Request) {

	var err error = nil
	if err = r.ParseForm(); err != nil {
		log.Println("can't parse form")
		log.Fatal(err)
	}
	names := r.FormValue("streamerNames")

	// redirect to /home if text input is empty
	if len(names) < 1 && r.FormValue("randomStream") != "" {
		randdomStreamer := GetRandomStreamer(w, r)
		http.Redirect(w, r, randdomStreamer, http.StatusTemporaryRedirect)
		return
	} else if len(names) < 1 {
		http.Redirect(w, r, "localhost:8080/home", http.StatusTemporaryRedirect)
		return
	}

	streamNames := twitch.ProcessNames(&names)
	log.Println(streamNames)

	url := "https://api.twitch.tv/helix/streams?"
	bodyBytes := hitTwitchApi(w, *r, url, streamNames)

	/*Parsing out twitch streamer username*/
	twitchUser := new(twitch.TwitchUsers)
	err = json.Unmarshal(bodyBytes, &twitchUser)
	if err != nil {
		log.Println()
		log.Println("can't parse json")
		log.Fatal(err)
	}

	length := len(twitchUser.Data)
	//TODO: allow users to choose which streamers to watch if 2 or more that the users entered are online
	if length > 0 {
		userName := twitchUser.Data[0].UserLogin
		url := "https://www.twitch.tv/" + userName
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	} else {
		url := "http://localhost:5500/notFound"
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}

func hitTwitchApi(w http.ResponseWriter, r http.Request, url string, streamNames []string) (data []byte) {
	client := http.Client{}

	if streamNames != nil {
		twitch.FormatUrl(&streamNames, &url)
	}

	log.Println(url)

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

	log.Println("*****Auth: ", req.Header.Values("Authorization"))

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
	return bodyBytes
}
