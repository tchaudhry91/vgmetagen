package vgmetagen

// GamesResponse is a data struct for unmarshalling json output from the GiantBomb api endpoint /games
type GamesResponse struct {
	ErrorStatus          string         `json:"error"`
	Limit                int            `json:"limit"`
	Offset               int            `json:"offset"`
	NumberOfPageResults  int            `json:"number_of_page_results"`
	NumberOfTotalResults int            `json:"number_of_total_results"`
	StatusCode           int            `json:"status_code"`
	Results              []GameResponse `json:"results"`
	Version              string         `json:"version"`
}

// GameGiantBombResponse is a data struct for umarshalling json output from the GiantBomb api endpoint /game
type GameGiantBombResponse struct {
	ErrorStatus          string `json:"error"`
	Limit                int    `json:"limit"`
	Offset               int    `json:"offset"`
	NumberOfPageResults  int    `json:"number_of_page_results"`
	NumberOfTotalResults int    `json:"number_of_total_results"`
	StatusCode           int    `json:"status_code"`
	Results              Game   `json:"results"`
	Version              string `json:"version"`
}

// GameResponse holds information for a game's Name and it's GiantBomb ID
type GameResponse struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// GamesDirectory holds information for Games in a map
type GamesDirectory struct {
	Games map[int]string `json:"games"`
}

// Game holds all the meta information for a Game
type Game struct {
	Aliases             string `json:"aliases"`
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	OriginalReleaseDate string `json:"original_release_date"`
	Platforms           []struct {
		APIDetailURL  string `json:"api_detail_url"`
		ID            int    `json:"id"`
		Name          string `json:"name"`
		SiteDetailURL string `json:"site_detail_url"`
		Abbreviation  string `json:"abbreviation"`
	} `json:"platforms"`
	Developers []struct {
		APIDetailURL  string `json:"api_detail_url"`
		ID            int    `json:"id"`
		Name          string `json:"name"`
		SiteDetailURL string `json:"site_detail_url"`
	} `json:"developers"`
	Publishers []struct {
		APIDetailURL  string `json:"api_detail_url"`
		ID            int    `json:"id"`
		Name          string `json:"name"`
		SiteDetailURL string `json:"site_detail_url"`
	} `json:"publishers"`
	Concepts []struct {
		APIDetailURL  string `json:"api_detail_url"`
		ID            int    `json:"id"`
		Name          string `json:"name"`
		SiteDetailURL string `json:"site_detail_url"`
	} `json:"concepts"`
	SimilarGames []struct {
		APIDetailURL  string `json:"api_detail_url"`
		ID            int    `json:"id"`
		Name          string `json:"name"`
		SiteDetailURL string `json:"site_detail_url"`
	} `json:"similar_games"`
}

// initMap initializes the map in the directory
func (directory *GamesDirectory) initMap() {
	directory.Games = make(map[int]string)
}

// addGame adds a game string as value to the directory with giantbomb id as key
func (directory *GamesDirectory) addGame(game GameResponse) {
	directory.Games[game.ID] = game.Name
}

// RandomGame returns a random game from the directory
func (directory *GamesDirectory) RandomGame() GameResponse {
	// Get a random element from the directory
	var game GameResponse
	for id, name := range directory.Games {
		game.ID = id
		game.Name = name
		break
	}
	return game
}
