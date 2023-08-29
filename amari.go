package amarigo

const version = "0.1.0"

func New(token string) *AmariBot {

	bot := &AmariBot{
		apiKey:  token,
		baseURL: "https://amaribot.com/api/v1",
		debug:   false,
		version: version,
	}

	return bot
}

func (bot *AmariBot) getBulkUserLevel(guildID string, userID string) (string, error) {
	return "", nil
}

func (bot *AmariBot) getCombinedLeaderboard(guildID string, userID string) (string, error) {
	return "", nil
}

func (bot *AmariBot) getGuildRewards(guildID string, options PaginationOptions) (string, error) {
	return "", nil
}
