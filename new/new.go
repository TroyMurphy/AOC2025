package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please provide a day number as an argument.")
		return
	}
	dayString := os.Args[1]

	_, err := strconv.Atoi(dayString)
	if err != nil {
		panic(err)
	}
	writeDay(dayString, []byte(""))

	baseURL := "https://adventofcode.com/2025"
	req, _ := http.NewRequest("GET", baseURL+"/day/"+dayString+"/input", nil)
	sessionToken, found := os.LookupEnv("AOC_TOKEN")
	if !found {
		fmt.Println("Set AOC_TOKEN in environment variables before runninng")
		return
	}
	sessionCookie := &http.Cookie{Name: "session", Value: sessionToken}
	req.AddCookie(sessionCookie)

	if response, err := http.DefaultClient.Do(req); err == nil {
		defer response.Body.Close()

		if body, err := io.ReadAll(response.Body); err == nil {
			writeDay(dayString, body)
		}
	}
}

func writeDay(day string, content []byte) {
	thisDir, _ := os.Getwd()
	path := filepath.Join(thisDir, "day"+day, "input.txt")
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	check(err)
	fo, err := os.Create(path)
	check(err)

	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	w := bufio.NewWriter(fo)

	check(err)
	if _, writeErr := w.Write(content); writeErr != nil {
		panic(writeErr)
	}
}
