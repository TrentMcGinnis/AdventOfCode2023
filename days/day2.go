package days

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func dayTwoPartOne(f io.Reader, cubes [3]int, triggers [3]string) {
	fmt.Println("STARTING DAY 2 PART 1")
	scanner := bufio.NewScanner(f)
	total := 0

	for scanner.Scan() {
		// do something with a line
		text := scanner.Text()
		splitted := strings.Split(text, ":")
		id := fmt.Sprint(strings.Split(splitted[0], " ")[1])
		cubeList := strings.Split(splitted[1], ";")
		good := true
		for _, elem := range cubeList {
			for colorIndex, trigger := range triggers {
				splitByComma := strings.Split(elem, ",")
				for _, val := range splitByComma {
					indexOfTrigger := strings.Index(val, trigger)
					if indexOfTrigger >= 0 {
						num := val[:indexOfTrigger]
						if conv, err := strconv.ParseInt(strings.TrimSpace(num), 10, 64); err == nil {
							if int(conv) > cubes[colorIndex] {
								good = false
							}
						}
					}
				}
			}
		}
		if good {
			if conv, err := strconv.ParseInt(strings.TrimSpace(id), 10, 64); err == nil {
				total += int(conv)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(total)
	fmt.Println("ENDING DAY 2 PART 1")
}

func dayTwoPartTwo(f io.Reader) {
	fmt.Println("STARTING DAY 2 PART 2")
	scanner := bufio.NewScanner(f)
	total := 0

	for scanner.Scan() {
		colors := [3][]int{}
		// do something with a line
		text := scanner.Text()
		splitted := strings.Split(text, ":")
		cubeList := strings.Split(splitted[1], ";")
		for _, elem := range cubeList {
			splitByComma := strings.Split(elem, ",")
			for _, val := range splitByComma {
				numberColor := strings.Split(val, " ")
				number := numberColor[1]
				color := numberColor[2]
				index := -1
				if color == "red" {
					index = 0
				} else if color == "green" {
					index = 1
				} else {
					index = 2
				}
				if index >= 0 {
					if conv, err := strconv.ParseInt(strings.TrimSpace(number), 10, 64); err == nil {
						colors[index] = append(colors[index], int(conv))
					}
				}
			}
		}
		maxRed := -1
		maxGreen := -1
		maxBlue := -1
		for idx, elem := range colors {
			max := -1
			for _, v := range elem {
				if v > max {
					max = v
				}
			}
			switch idx {
			case 0:
				maxRed = max
			case 1:
				maxGreen = max
			case 2:
				maxBlue = max
			}
		}
		total += maxRed * maxGreen * maxBlue
	}
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("ENDING DAY 2 PART 2")
}

func DayTwo() {
	cubes := [3]int{12, 13, 14}
	triggers := [3]string{"red", "green", "blue"}
	fmt.Println("STARTING DAY 2")
	f, _ := os.Open("data/day2")
	f2, _ := os.Open("data/day2")
	defer f.Close()
	defer f2.Close()
	dayTwoPartOne(f, cubes, triggers)
	dayTwoPartTwo(f2)
	fmt.Println("ENDING DAY 2")
}
