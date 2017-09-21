package vgmetagen

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	urlRawBase = "https://giantbomb.com/api"
)

// InitGamesList initializes a games list with Names and Giantbomb IDs of top N games.
// The list is sorted in ascending order based on the number of user reviews to guage popularity
func InitGamesList(apiKey string, num int) (GamesDirectory, error) {
	var games GamesDirectory
	var errorReturn error
	resultsChan := make(chan GameResponse, 1000)
	numGoRoutines := (num / 100) + 1
	offsetChan := make(chan int, 1)
	log.Infof("Will need %d go routines", numGoRoutines)
	for i := 0; i <= num; i += 100 {
		offsetChan <- i
		go func() {
			urlGames, err := url.Parse(urlRawBase)
			if err != nil {
				log.Panicf("Could Not Parse raw urlGames %s", urlRawBase)
			}
			urlGames.Path += "/games"
			offset := <-offsetChan
			params := url.Values{}
			params.Add("api_key", apiKey)
			params.Add("offset", strconv.Itoa(offset))
			params.Add("sort", "number_of_user_reviews:desc")
			params.Add("format", "json")
			params.Add("field_list", "name,id")
			urlGames.RawQuery = params.Encode()

			log.Infof("Making call to for offset %d\n", offset)
			response, err := http.Get(urlGames.String())
			if err != nil {
				log.Error("Could not get Games List because:", err)
				errorReturn = err
			}
			defer response.Body.Close()
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Error("Could not read body:", err)
				errorReturn = err
			}
			var gamesData GamesResponse
			err = json.Unmarshal(body, &gamesData)
			if err != nil {
				log.Error("Error UnMarshalling JSON:", err)
				errorReturn = err
			}
			for _, gameResp := range gamesData.Results {
				resultsChan <- gameResp
			}
			numGoRoutines--
			if numGoRoutines == 0 {
				close(resultsChan)
			}
		}()
	}
	games.initMap()
	for game := range resultsChan {
		games.addGame(game)
	}
	return games, errorReturn
}
