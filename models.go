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

// GameResponse holds information for a game's Name and it's GiantBomb ID
type GameResponse struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// GamesDirectory holds information for Games in a map
type GamesDirectory struct {
	Games map[int]string `json:"games"`
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
