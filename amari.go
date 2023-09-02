package amarigo

import (
    "bytes"
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
        fmt.Print("DEBUG: ")
        fmt.Println(data)
    }
}

func (bot *AmariBot) fetch(url string) (body []byte, resp *http.Response, err error) {
    url = fmt.Sprintf("%s%s", bot.BaseURL, url)

    bot.printDebug(url)

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

    bot.printDebug(string(body))

    return
}

func (bot *AmariBot) fetchBody(url string, inputBody []byte) (body []byte, resp *http.Response, err error) {
    url = fmt.Sprintf("%s%s", bot.BaseURL, url)

    bot.printDebug(url)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return
    }

    req.Header.Set("Authorization", bot.ApiKey)
    req.Header.Set("User-Agent", fmt.Sprintf("amarigo/%s", bot.Version))

    req.Body = io.NopCloser(bytes.NewReader(inputBody))

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

    bot.printDebug(string(body))

    return
}

func (bot *AmariBot) GetLeaderboardPage(guildID string, page int) (leaderboard *Leaderboard, err error) {

    url := fmt.Sprintf("/guild/leaderboard/%s?page=%d", guildID, page)

    body, _, err := bot.fetch(url)
    if err != nil {
        return nil, err
    }

    bot.printDebug(body)
    err = json.Unmarshal(body, &leaderboard)
    bot.printDebug(leaderboard)

    return
}

func (bot *AmariBot) GetLeaderboardRaw(guildID string) (leaderboard *LeaderboardRaw, err error) {

    url := fmt.Sprintf("/guild/raw/leaderboard/%s", guildID)

    body, _, err := bot.fetch(url)
    if err != nil {
        return nil, err
    }

    bot.printDebug(body)
    err = json.Unmarshal(body, &leaderboard)
    bot.printDebug(leaderboard)

    return
}

func (bot *AmariBot) GetWeekly(guildID string) (weekly *Weekly, err error) {

    url := fmt.Sprintf("/guild/weekly/%s", guildID)

    body, _, err := bot.fetch(url)
    if err != nil {
        return nil, err
    }

    bot.printDebug(body)
    err = json.Unmarshal(body, &weekly)
    bot.printDebug(weekly)

    return
}

func (bot *AmariBot) GetWeeklyRaw(guildID string) (weekly *WeeklyRaw, err error) {

    url := fmt.Sprintf("/guild/raw/weekly/%s", guildID)

    body, _, err := bot.fetch(url)
    if err != nil {
        return nil, err
    }

    bot.printDebug(body)
    err = json.Unmarshal(body, &weekly)
    bot.printDebug(weekly)

    return
}

func (bot *AmariBot) GetGuildMember(guildID, userID string) (member *User, err error) {

    url := fmt.Sprintf("/guild/%s/member/%s", guildID, userID)

    body, _, err := bot.fetch(url)
    if err != nil {
        return nil, err
    }

    bot.printDebug(body)
    err = json.Unmarshal(body, &member)
    bot.printDebug(member)

    return
}

func (bot *AmariBot) GetGuildMembers(guildID string, members GetGuildMembers) (users []*User, err error) {

    url := fmt.Sprintf("/guild/%s/members", guildID)

    memberBody, err := json.Marshal(members.Members)

    body, _, err := bot.fetchBody(url, memberBody)
    if err != nil {
        return nil, err
    }

    bot.printDebug(body)
    err = json.Unmarshal(body, &members)
    bot.printDebug(members)

    return

}

func (bot *AmariBot) GetGuildRewards(guildID string) (rewards []*Reward, err error) {

    url := fmt.Sprintf("/guild/rewards/%s", guildID)

    body, _, err := bot.fetch(url)
    if err != nil {
        return nil, err
    }

    bot.printDebug(body)
    err = json.Unmarshal(body, &rewards)
    bot.printDebug(rewards)

    return
}
