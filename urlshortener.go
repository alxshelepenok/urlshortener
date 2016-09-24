package urlshortener

import (
	"encoding/json"
	"github.com/parnurzeal/gorequest"
	"fmt"
)

var key string

type ShortMsg struct {
	Kind    string `json:"kind"`
	Id      string `json:"id"`
	LongUrl string `json:"longUrl"`
}

func Setup(k string) error {
	key = k

	return nil
}

func Shorten(url string) (string, error) {
	var ShortMsg *ShortMsg

	request := gorequest.New()

	gUrl := "https://www.googleapis.com/urlshortener/v1/url?key=" + key

	res, _, err := request.Post(gUrl).
	Set("Accept", "application/json").
	Set("Content-Type", "application/json").
	Send(`{"longUrl":"` + url + `"}`).End()

	if err != nil {
		return "", fmt.Errorf("Request error: %v", err)
	}

	if res.Status == "200 OK" {
		if err := json.NewDecoder(res.Body).Decode(&ShortMsg); err != nil {
			return "", fmt.Errorf("Decode error: %v", err)
		}

		return ShortMsg.Id, nil
	}

	return "", fmt.Errorf("Request error: %v", res.Status)
}