package main

import (
	"bufio"
	"day2/utils"
	"fmt"
	"os"
	"path/filepath"
)

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		panic(err)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	utils.Check(err)
	defer closeFile(file)
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	// read lines from input.txt
	curdir, err := os.Getwd()
	utils.Check(err)

	path := filepath.Join(curdir, "day2", "input.txt")
	// path := filepath.Join(curdir, "day2", "sample.txt")
	lines, err := readLines(path)
	utils.Check(err)

	part1 := Part1(lines)
	part2 := Part2(lines)

	fmt.Println(part1)
	fmt.Println(part2)
}
