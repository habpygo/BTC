package utils

import (
	"bufio"
	"fmt"
	"os"
)

var PokerArray []string

// ReadLine reads all the lines in a game and returns a string slice
func ReadLine(file string) ([]string, int, error) {

	count := 0
	infile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer infile.Close()
	var lines []string
	scanner := bufio.NewScanner(infile)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		count++
	}
	return lines, count, scanner.Err()
}
