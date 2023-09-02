package amarigo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var client = &http.Client{}

const version = "0.1.0"

func New(token string) *AmariBot {

	bot := &AmariBot{
		ApiKey:  token,
		BaseURL: "https://amaribot.com/api/v1",
		Debug:   false,
		Version: version,
	}

	_, resp, err := bot.fetch("/guilds/1/leaderboard")
	if resp.StatusCode == 403 {
		panic("Invalid API key")
	} else if err != nil {
		panic(err)
	}

	return bot
}

func (bot *AmariBot) printDebug(data interface{}) {
	if bot.Debug {
		fmt.Println(fmt.Sprintf("%v", data))
	}
}

func (bot *AmariBot) fetch(url string) (body []byte, resp *http.Response, err error) {
	url = fmt.Sprintf("%s/%s", bot.BaseURL, url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.Header.Set("Authorization", bot.ApiKey)
	req.Header.Set("User-Agent", fmt.Sprintf("amarigo/%s", bot.Version))

	resp, err = client.Do(req)
	if err != nil {
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	body, err = io.ReadAll(resp.Body)
	return
}

func (bot *AmariBot) GetLeaderboardPage(guildID string, page int) (leaderboard *Leaderboard, err error) {

	url := fmt.Sprintf("/guilds/%s/leaderboard?page=%d", guildID, page)

	body, _, err := bot.fetch(url)
	if err != nil {
		return nil, err
	}

	bot.printDebug(body)

	err = json.Unmarshal(body, &leaderboard)
	bot.printDebug(leaderboard)

	return
}
