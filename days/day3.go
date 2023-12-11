package days

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type cog struct {
	x     int
	y     int
	parts []int
}

func dayThreePartOne(f io.Reader) {
	fmt.Println("STARTING DAY 3 PART 1")
	data := ""
	scanner := bufio.NewScanner(f)
	coords := [][3]int{}
	y := 0
	height := 0
	width := 0

	for scanner.Scan() {
		// do something with a line
		text := scanner.Text()
		width = len(string(text))
		data += string(text)
		str := ""
		// x, y, value
		coord := [3]int{-1, -1, -1}
		for x, char := range text {
			_, err := strconv.ParseInt(string(char), 10, 64)
			if err == nil {
				if coord[0] == -1 {
					coord[0] = x
					coord[1] = y
				}
				str += string(char)
			} else if str != "" {
				conv, _ := strconv.ParseInt(string(str), 10, 64)
				coord[2] = int(conv)
				coords = append(coords, coord)
				coord = [3]int{-1, -1, -1}
				str = ""
			}
		}
		if str != "" {
			conv, _ := strconv.ParseInt(string(str), 10, 64)
			coord[2] = int(conv)
			coords = append(coords, coord)
			coord = [3]int{-1, -1, -1}
			str = ""
		}
		y++
	}
	height = y
	total := 0

	for _, coord := range coords {
		x, y, val := coord[0], coord[1], coord[2]
		above := ""
		left := ""
		right := ""
		below := ""

		textLength := len(fmt.Sprint(val))

		if y > 0 {
			index := x + (y-1)*width
			if x > 0 {
				above = data[index-1 : index+textLength+1]
			} else {
				above = data[index : index+textLength+1]
			}
		}
		if x > 0 {
			index := x + y*width
			left = string(data[index-1 : index])
		}
		if y < (height - 1) {
			index := x + (y+1)*width
			if x > 0 {
				below = data[index-1 : index+textLength+1]
			} else {
				below = data[index : index+textLength+1]
			}
		}
		if x < width {
			index := x + y*width
			right = string(data[index+textLength])
		}
		check := above + below + left + right
		good := false
		for _, v := range check {
			if string(v) != "." {
				good = true
			}
		}
		if good {
			total += val
		}
	}
	fmt.Println(total)
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("ENDING DAY 3 PART 1")
}

func checkForCog(val string, x int, y int, cogHashes map[string]cog, part int) {
	idx := strings.Index(val, "*")
	if idx >= 0 {
		cogHash := fmt.Sprintf("X:%dY:%d", x+idx, y)
		cogFound, exists := cogHashes[cogHash]
		if exists {
			cogFound.parts = append(cogFound.parts, part)
			cogHashes[cogHash] = cogFound
		}
	}
}

func dayThreePartTwo(f io.Reader) {
	fmt.Println("STARTING DAY 3 PART 2")
	data := ""
	scanner := bufio.NewScanner(f)
	coords := [][3]int{}
	cogHashes := make(map[string]cog)

	y := 0
	height := 0
	width := 0

	for scanner.Scan() {
		// do something with a line
		text := scanner.Text()
		width = len(string(text))
		data += string(text)
		str := ""
		// x, y, value
		coord := [3]int{-1, -1, -1}
		for x, char := range text {
			_, err := strconv.ParseInt(string(char), 10, 64)
			if err == nil {
				if coord[0] == -1 {
					coord[0] = x
					coord[1] = y
				}
				str += string(char)
			} else {
				if str != "" {
					conv, _ := strconv.ParseInt(string(str), 10, 64)
					coord[2] = int(conv)
					coords = append(coords, coord)
					coord = [3]int{-1, -1, -1}
					str = ""
				}
				if string(char) == "*" {
					cog := cog{
						x,
						y,
						[]int{},
					}
					cogHash := fmt.Sprintf("X:%dY:%d", x, y)
					cogHashes[cogHash] = cog
				}
			}
		}
		if str != "" {
			conv, _ := strconv.ParseInt(string(str), 10, 64)
			coord[2] = int(conv)
			coords = append(coords, coord)
			coord = [3]int{-1, -1, -1}
			str = ""
		}
		y++
	}
	height = y

	for _, coord := range coords {
		// fmt.Println(coord)
		x, y, val := coord[0], coord[1], coord[2]
		above := ""
		left := ""
		right := ""
		below := ""

		textLength := len(fmt.Sprint(val))

		if y > 0 {
			index := x + (y-1)*width
			if x > 0 {
				above = data[index-1 : index+textLength+1]
				checkForCog(above, x-1, y-1, cogHashes, val)
			} else {
				above = data[index : index+textLength+1]
				checkForCog(above, x, y-1, cogHashes, val)
			}
		}
		if x > 0 {
			index := x + y*width
			left = string(data[index-1 : index])
			checkForCog(left, x-1, y, cogHashes, val)
		}
		if y < (height - 1) {
			index := x + (y+1)*width
			if x > 0 {
				below = data[index-1 : index+textLength+1]
				checkForCog(below, x-1, y+1, cogHashes, val)
			} else {
				below = data[index : index+textLength+1]
				checkForCog(below, x, y+1, cogHashes, val)
			}
		}
		if x < width {
			index := x + y*width
			right = string(data[index+textLength])
			checkForCog(right, x+textLength, y, cogHashes, val)
		}
	}

	total := 0
	for _, cog := range cogHashes {
		if len(cog.parts) == 2 {
			total += cog.parts[0] * cog.parts[1]
		}
	}

	fmt.Println(total)
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("ENDING DAY 3 PART 2")
}

func DayThree() {
	fmt.Println("STARTING DAY 3")
	f, _ := os.Open("data/day3")
	f2, _ := os.Open("data/day3")
	defer f.Close()
	defer f2.Close()
	dayThreePartOne(f)
	dayThreePartTwo(f2)
	fmt.Println("ENDING DAY 3")
}
