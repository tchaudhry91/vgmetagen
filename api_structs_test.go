package vgmetagen

import (
	"testing"
)

var sampleGame1, sampleGame2, sampleGame3 GameResponse
var directory GamesDirectory

func init() {
	sampleGame1 = GameResponse{
		Name: "Awesome Game",
		ID:   1,
	}
	sampleGame2 = GameResponse{
		Name: "Awesome Game 2",
		ID:   2,
	}
	sampleGame2 = GameResponse{
		Name: "Awesome Game 3",
		ID:   3,
	}
	directory = GamesDirectory{}
	directory.initMap()
}

func TestGameAddition(t *testing.T) {
	directory.addGame(sampleGame1)
	directory.addGame(sampleGame2)
	if name := directory.Games[sampleGame1.ID]; name != "Awesome Game" {
		t.Errorf("Wrong Value for game addition: %s", name)
	}
}

func TestRandomnessOfReturnedGame(t *testing.T) {
	directory.addGame(sampleGame1)
	directory.addGame(sampleGame2)
	directory.addGame(sampleGame3)

	testMap := make(map[int]int)

	for i := 0; i < 100; i++ {
		game := directory.RandomGame()
		testMap[game.ID]++
	}
	for key, value := range testMap {
		t.Logf("%d:%d", key, value)
		if value < 10 {
			t.Errorf("Too low value for %d", key)
		}
	}
}
