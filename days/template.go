package days

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func _partOne(f io.Reader) {
	fmt.Println("STARTING DAY _DAY_ PART 1")
	scanner := bufio.NewScanner(f)
	total := 0

	for scanner.Scan() {
		// do something with a line
		text := scanner.Text()
		fmt.Println(text)
	}
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("ENDING DAY _DAY_ PART 1")
}

func _partTwo(f io.Reader) {
	fmt.Println("STARTING DAY _DAY_ PART 2")
	scanner := bufio.NewScanner(f)
	total := 0

	for scanner.Scan() {
		// do something with a line
		text := scanner.Text()
		fmt.Println(text)
	}
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("ENDING DAY _DAY_ PART 2")
}

func Day_DAY_() {
	fmt.Println("STARTING DAY _DAY_")
	f, _ := os.Open("data/day_DAY_")
	f2, _ := os.Open("data/day_DAY_")
	defer f.Close()
	defer f2.Close()
	_partOne(f)
	_partTwo(f2)
	fmt.Println("ENDING DAY _DAY_")
}
