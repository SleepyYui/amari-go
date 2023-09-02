package amarigo

type AmariBot struct {
	ApiKey      string
	BaseURL     string
	CustomFetch string
	Debug       bool
	Version     string
}

type AmariError struct {
	Message string
	Name    string
	Status  int
}

type RatelimitError struct {
	Message   string
	Name      string
	Status    int
	Remaining int
}

type RequestHandler struct {
	_client *AmariBot
}

type Leaderboard struct {
	Data  []LeaderboardUser `json:"current_page"`
	Count int               `json:"count"`
	Total int               `json:"total_count"`
}

type LeaderboardUser struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Exp       int    `json:"exp"`
	Level     int    `json:"level"`
	WeeklyExp int    `json:"weeklyExp"`
}
