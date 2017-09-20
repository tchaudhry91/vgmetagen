package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

var (
	gamesList []GameResponse
)

const (
	urlRawBase = "https://giantbomb.com/api"
)

type GamesResponse struct {
	ErrorStatus          string         `json:"error"`
	Limit                float64        `json:"limit"`
	Offset               float64        `json:"offset"`
	NumberOfPageResults  float64        `json:"number_of_page_results"`
	NumberOfTotalResults float64        `json:"number_of_total_results"`
	StatusCode           float64        `json:"status_code"`
	Results              []GameResponse `json:"results"`
	Version              string         `json:"version"`
}

type GameResponse struct {
	Name string  `json:"name"`
	Id   float64 `json:"id"`
}

func initGamesList(num int) {
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
			params.Add("api_key", os.Getenv("GB_API_KEY"))
			params.Add("offset", strconv.Itoa(offset))
			params.Add("sort", "number_of_user_reviews:desc")
			params.Add("format", "json")
			params.Add("field_list", "name,id")
			urlGames.RawQuery = params.Encode()

			log.Infof("Making call to %s\n", urlGames.String())
			response, err := http.Get(urlGames.String())
			if err != nil {
				log.Error("Could not get Games List because:", err)
			}
			defer response.Body.Close()
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Error("Could not read body:", err)
				return
			}
			var gamesData GamesResponse
			err = json.Unmarshal(body, &gamesData)
			if err != nil {
				log.Error("Error UnMarshalling JSON:", err)
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
	for game := range resultsChan {
		gamesList = append(gamesList, game)
	}
}

func getRandomGame() {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 100; i++ {
		index := rand.Intn(len(gamesList))
		log.Infof("Game:%s - Id:%f", gamesList[index].Name, gamesList[index].Id)
	}
}

func main() {
	initGamesList(150)
	getRandomGame()
}
