package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		panic(err)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	check(err)
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
	check(err)

	path := filepath.Join(curdir, "day6", "input.txt")
	// path := filepath.Join(curdir, "day6", "sample.txt")
	lines, err := readLines(path)
	check(err)

	// part1 := Part1(lines)
	part2 := Part2(lines)

	// fmt.Println(part1)
	// fmt.Println("-------------")
	fmt.Println(part2)
}
