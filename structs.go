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
	Users []*LeaderboardUser `json:"data"`
	Count int                `json:"count"`
	Total int                `json:"total_count"`
}

type LeaderboardRaw struct {
	Users []*LeaderboardUser `json:"data"`
	Count int                `json:"count"`
}

type LeaderboardUser struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Exp       int    `json:"exp"`
	Level     int    `json:"level"`
	WeeklyExp int    `json:"weeklyExp"`
}

type Weekly struct {
	Users []*WeeklyUser `json:"data"`
	Count int           `json:"count"`
	Total int           `json:"total_count"`
}

type WeeklyRaw struct {
	Users []*WeeklyUser `json:"data"`
	Count int           `json:"count"`
}

type WeeklyUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Exp      int    `json:"exp"`
}

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Exp       int    `json:"exp"`
	Level     int    `json:"level"`
	WeeklyExp int    `json:"weeklyExp"`
}

type Rewards struct {
	Rewards []*Reward `json:"data"`
	Count   int       `json:"count"`
}

type Reward struct {
	RoleID string `json:"role_id"`
	Level  int    `json:"level"`
}

type MemberIDList struct {
	Members []string `json:"members"`
}

type GetGuildMembers struct {
	GuildID        string  `json:"guild"`
	Members        []*User `json:"members"`
	TotalMembers   int     `json:"total_members"`
	QueriedMembers int     `json:"queried_members"`
}
