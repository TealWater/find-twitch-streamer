package utils

import (
	"strings"
	"time"
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

/*
Takes a string full of comma separated twitch user names and formats them to be sent to the twitch api
*/
func ProcessNames(names *string) (newNames []string) {

	stringSlice := strings.Split(*names, ",")
	var i int = 0
	for _, v := range stringSlice {
		stringSlice[i] = strings.ReplaceAll(strings.ToLower(v), " ", "")
		i += 1
	}

	return stringSlice
}

/*
Adds '&' to the twitch api url
*/
func FormatUrl(stringSlice *[]string, url *string) {
	var i int = 0
	var length int = len(*stringSlice)
	for _, v := range *stringSlice {
		user := "user_login=" + v

		if i < length-1 {
			user += "&"
		}

		(*stringSlice)[i] = user
		*url += (*stringSlice)[i]
		i += 1
	}
}
