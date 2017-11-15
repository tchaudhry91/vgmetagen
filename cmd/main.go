package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/tchaudhry91/vgmetagen"
	"log"
	"os"
)

func prettyPrintGame(game vgmetagen.Game) {
	fmt.Println("Release Date:", game.OriginalReleaseDate)
	fmt.Println("Platforms:")
	for _, item := range game.Platforms {
		fmt.Println(item.Name)
	}
	fmt.Println("Developers:")
	for _, item := range game.Developers {
		fmt.Println(item.Name)
	}
	fmt.Println("Publishers:")
	for _, item := range game.Publishers {
		fmt.Println(item.Name)
	}
	fmt.Println("SimilarGames:")
	for _, item := range game.SimilarGames {
		fmt.Println(item.Name)
	}
}

func main() {
	var apiKey = flag.String("apiKey", "", "Input your GiantBomb Api Key")
	var directorySize = flag.Int("directorySize", 100, "Select the size of the Games Directory to create")
	var directoryStep = flag.Int("directoryStep", 20, "Select the size of games per each request while building directory")
	flag.Parse()

	if *apiKey == "" {
		log.Fatal("Empty Api Key, please run with '-h' for details")
	}

	directory, err := vgmetagen.InitGamesList(*apiKey, *directorySize, *directoryStep)
	if err != nil {
		log.Fatal("Error building Games list")
	}

	for {
		game, err := directory.RandomGame()
		if err != nil {
			log.Fatal("Improper game directory")
		}
		gameData, err := vgmetagen.GetGameData(*apiKey, game.ID)
		if err != nil {
			log.Print("Error while getting game data:", err)
			continue
		}
		prettyPrintGame(gameData)
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("GUESS:")
		scanner.Scan()
		guess := scanner.Text()
		fmt.Printf("Your Guess: %s\nActual Game: %s", guess, game.Name)
	}
}
