package amarigo

type AmariBot struct {
	apiKey      string
	baseURL     string
	customFetch string
	debug       bool
	version     string
}

type AmariError struct {
	message string
	name    string
	status  int
}

type RatelimitError struct {
	message   string
	name      string
	status    int
	remaining int
}

type RequestHandler struct {
	_client *AmariBot
}

type APIBulkUsers struct {
	members         []APIUser
	queried_members int
	total_members   int
}

type APIError struct {
	error string
}

type APILeaderboard struct {
	count       int
	data        []APILeaderboardUser
	total_count int
}

type APILeaderboardUser struct {
	exp      int
	id       string
	level    int
	username string
}

type APIRawLeaderboard struct {
	count int
	data  []APILeaderboardUser
}

type APIRawWeeklyLeaderboard struct {
	count int
	data  []APIWeeklyUser
}

type APIRewards struct {
	count int
	data  []APIReward
}

type APIReward struct {
	level  int
	roleID string
}

type APIUser struct {
	exp       int
	id        string
	level     int
	username  string
	weeklyExp int
}

type APIWeeklyLeaderboard struct {
	count       int
	data        []APIWeeklyUser
	total_count int
}

type APIWeeklyUser struct {
	exp      int
	id       string
	username string
}

type CustomFetch struct {
	url     string
	options map[string]string
}

type CustomResponse struct {
	body string
}

type PaginationOptions struct {
	limit int
	page  int
}

type RawPaginationOptions struct {
	limit int
}
