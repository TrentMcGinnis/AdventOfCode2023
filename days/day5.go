package days

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type data struct {
	key         int
	startSource int
	endSource   int
	startDest   int
	endDest     int
}

type seed struct {
	seed       int
	soil       int
	fertilizer int
	water      int
	light      int
	temp       int
	humidity   int
	location   int
}

func _5partOne(f io.Reader) {
	fmt.Println("STARTING DAY 5 PART 1")
	lines := []string{}
	seeds := []seed{}
	dataPoints := []data{}
	scanner := bufio.NewScanner(f)

	index := 0
	marker := false
	concat := []string{}
	for scanner.Scan() {
		// do something with a line
		text := scanner.Text()
		if index == 0 {
			for _, seed_v := range strings.Split(strings.Split(text, ":")[1], " ") {
				seed_int, err := strconv.ParseInt(seed_v, 10, 64)
				if err == nil {
					seeds = append(seeds, seed{
						seed:       int(seed_int),
						soil:       -1,
						fertilizer: -1,
						water:      -1,
						light:      -1,
						temp:       -1,
						humidity:   -1,
						location:   -1,
					})
				}
			}
		} else if strings.Contains(text, ":") {
			marker = true
		} else if text == "" {
			marker = false
			if len(concat) >= 1 {
				lines = append(lines, strings.Join(concat, "::"))
				concat = []string{}
			}
		} else if marker {
			concat = append(concat, text)
		}
		index++
	}

	for idx, line := range lines {
		// fmt.Println(line)
		for _, part := range strings.Split(line, "::") {
			spl := strings.Split(strings.Trim(part, "\n"), " ")
			// fmt.Println(spl)
			// Dest, Source, Range Length
			destStart, _ := strconv.ParseInt(spl[0], 10, 64)
			sourceStart, _ := strconv.ParseInt(spl[1], 10, 64)
			rangeLength, _ := strconv.ParseInt(spl[2], 10, 64)

			// fmt.Println(idx)
			// fmt.Printf("Start S: %d\n", sourceStart)
			// fmt.Printf("End S: %d\n", sourceStart+rangeLength)
			// fmt.Printf("Start D: %d\n", destStart)
			// fmt.Printf("End S: %d\n", destStart+rangeLength)
			// fmt.Println("=========")

			dataPoints = append(dataPoints, data{
				key:         idx,
				startSource: int(sourceStart),
				endSource:   int(sourceStart) + int(rangeLength) - 1,
				startDest:   int(destStart),
				endDest:     int(destStart) + int(rangeLength) - 1,
			})
		}
	}
	lowestLoc := math.MaxInt

	// fmt.Println(dataPoints)

	for _, s := range seeds {
		s.soil = s.seed
		for _, d := range dataPoints {
			switch d.key {
			case 0:
				if s.seed >= d.startSource &&
					s.seed <= d.endSource {
					s.soil = d.startDest + (s.seed - d.startSource)
				}
				s.fertilizer = s.soil
			case 1:
				if s.soil >= d.startSource &&
					s.soil <= d.endSource {
					s.fertilizer = d.startDest + (s.soil - d.startSource)
				}
				s.water = s.fertilizer
			case 2:
				if s.fertilizer >= d.startSource &&
					s.fertilizer <= d.endSource {
					s.water = d.startDest + (s.fertilizer - d.startSource)
				}
				s.light = s.water
			case 3:
				if s.water >= d.startSource &&
					s.water <= d.endSource {
					s.light = d.startDest + (s.water - d.startSource)
				}
				s.temp = s.light
			case 4:
				if s.light >= d.startSource &&
					s.light <= d.endSource {
					s.temp = d.startDest + (s.light - d.startSource)
				}
				s.humidity = s.temp
			case 5:
				if s.temp >= d.startSource &&
					s.temp <= d.endSource {
					s.humidity = d.startDest + (s.temp - d.startSource)
				}
				s.location = s.humidity
			case 6:
				if s.humidity >= d.startSource &&
					s.humidity <= d.endSource {
					s.location = d.startDest + (s.humidity - d.startSource)
				}
				s.location = s.humidity
			}
		}
		if s.location < lowestLoc {
			lowestLoc = s.location
		}
		fmt.Println(s)
	}
	fmt.Println(lowestLoc)
	fmt.Println("ENDING DAY 5 PART 1")
}

func _5partTwo(f io.Reader) {
	fmt.Println("STARTING DAY 5 PART 2")
	scanner := bufio.NewScanner(f)
	// total := 0

	// for scanner.Scan() {
	// 	// do something with a line
	// 	text := scanner.Text()
	// 	fmt.Println(text)
	// }
	// fmt.Println(total)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("ENDING DAY 5 PART 2")
}

func Day5() {
	fmt.Println("STARTING DAY 5")
	f, _ := os.Open("data/day5")
	f2, _ := os.Open("data/day5")
	defer f.Close()
	defer f2.Close()
	_5partOne(f)
	_5partTwo(f2)
	fmt.Println("ENDING DAY 5")
}
