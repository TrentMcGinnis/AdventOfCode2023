package days

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

func partOne(f io.Reader) {
	fmt.Println("STARTING DAY 4 PART 1")
	scanner := bufio.NewScanner(f)
	total := 0

	for scanner.Scan() {
		// do something with a line
		matches := []string{}
		text := scanner.Text()
		gameData := strings.Split(text, ":")[1]
		split := strings.Split(gameData, "|")
		winners := strings.Split(split[0], " ")
		games := strings.Split(split[1], " ")
		for _, game := range games {
			for _, win := range winners {
				if game == win && game != "" && win != "" {
					matches = append(matches, game)
				}
			}
		}
		total += int(math.Pow(2, float64(len(matches))-1))
	}
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("ENDING DAY 4 PART 1")
}

type game struct {
	index int
	wins  int
	loops int
}

func partTwo(f io.Reader) {
	fmt.Println("STARTING DAY 4 PART 2")
	scanner := bufio.NewScanner(f)
	total := 0
	gamesList := []game{}

	index := 0
	for scanner.Scan() {
		// do something with a line
		matches := []string{}
		text := scanner.Text()
		gameData := strings.Split(text, ":")[1]
		split := strings.Split(gameData, "|")
		winners := strings.Split(split[0], " ")
		games := strings.Split(split[1], " ")
		for _, game := range games {
			for _, win := range winners {
				if game == win && game != "" && win != "" {
					matches = append(matches, game)
				}
			}
		}
		game := game{
			index: index,
			wins:  len(matches),
			loops: 1,
		}
		gamesList = append(gamesList, game)
		index += 1
	}
	for index, game := range gamesList {
		total += game.loops
		for j := 0; j < game.loops; j++ {
			for i := 1; i <= game.wins; i++ {
				gamesList[index+i].loops += 1
			}
		}
	}
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("ENDING DAY 4 PART 2")
}

func DayFour() {
	fmt.Println("STARTING DAY 4")
	f, _ := os.Open("data/day4")
	f2, _ := os.Open("data/day4")
	defer f.Close()
	defer f2.Close()
	partOne(f)
	partTwo(f2)
	fmt.Println("ENDING DAY 4")
}
