package days

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func dayOnePartOne(f io.Reader) {
	fmt.Println("STARTING DAY 1 PART 1")
	scanner := bufio.NewScanner(f)
	total := 0

	for scanner.Scan() {
		// do something with a line
		text := scanner.Text()
		holder := []string{}
		for _, char := range text {
			if _, err := strconv.ParseInt(string(char), 10, 64); err == nil {
				holder = append(holder, string(char))
			}
		}
		final := holder[0] + holder[len(holder)-1]
		end, _ := strconv.ParseInt(string(final), 10, 64)
		total += int(end)
	}
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("ENDING DAY 1 PART 1")
}

func dayOnePartTwo(f io.Reader) {
	fmt.Println("STARTING DAY 1 PART 2")
	scanner := bufio.NewScanner(f)
	words := [10]string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
	trigger := [7]string{
		"z",
		"o",
		"t",
		"f",
		"s",
		"e",
		"n",
	}
	total := 0

	for scanner.Scan() {
		// do something with a line
		text := scanner.Text()
		holder := []string{}
		for idx, char := range text {
			if _, err := strconv.ParseInt(string(char), 10, 64); err == nil {
				holder = append(holder, string(char))
			} else {
				for _, t := range trigger {
					if string(char) == t {
						for index, word := range words {
							if idx+len(word) <= len(text) {
								check := text[idx : idx+len(word)]
								if check == word {
									holder = append(holder, fmt.Sprint(index))
								}
							}
						}
					}
				}
			}
		}
		final := holder[0] + holder[len(holder)-1]
		end, _ := strconv.ParseInt(string(final), 10, 64)
		total += int(end)
	}
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("ENDING DAY 1 PART 2")
}

func DayOne() {
	fmt.Println("STARTING DAY 1")
	f, _ := os.Open("data/day1")
	f2, _ := os.Open("data/day1")
	defer f.Close()
	defer f2.Close()
	dayOnePartOne(f)
	dayOnePartTwo(f2)
	fmt.Println("ENDING DAY 1")
}
