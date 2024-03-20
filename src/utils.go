package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"unicode/utf8"
)

/*
Get the color codes for terminal output

	:parameter
	:return
	*	colorMap: map that returns the terminal escape sequences to color strings
*/
func getColorMap() map[string]string {
	colorMap := make(map[string]string)
	colorMap["reset"] = "\033[0m"
	colorMap["bold"] = "\033[1m"
	colorMap["underline"] = "\033[4m"
	colorMap["strike"] = "\033[9m"
	colorMap["italic"] = "\033[3m"
	colorMap["red"] = "\033[31m"
	colorMap["green"] = "\033[32m"
	colorMap["yellow"] = "\033[33m"
	colorMap["blue"] = "\033[34m"
	colorMap["purple"] = "\033[35m"
	colorMap["cyan"] = "\033[36m"
	colorMap["white"] = "\033[37m"
	return colorMap
}

/*
Check if flag was used

	:parameter
	*	name: name of the flag
	:return
	*	found: true if flag was supplied
*/
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

/*
Check if file is hidden (more than one '.' in the path)

	:parameter
	*	s: string to test
	*	dot_max: max number of dots to be true
	:return
	*	dot_count: whether there were more than dot_max
*/
func isHidden(s *string, dotMax int) bool {
	var testRune = '.'
	dot_count := 0
	for _, i := range *s {
		if i == testRune {
			dot_count++
		}
	}
	return dot_count > dotMax
}

/*
Reverse a slice containing runes in place

	:parameter
	*	inSlice: slice that should be reversed
	:return
*/
func reverseRune(inSlice []rune) {
	n := len(inSlice)
	for i := 0; i < n/2; i++ {
		inSlice[i], inSlice[n-1-i] = inSlice[n-1-i], inSlice[i]
	}
}

/*
Check whether a file is a binary file by checking if the first 4096 byte can be converted to utf8

	:parameter
	*	filePath: path to the file that should be tested
	:return
	*	binTest: true if file can't be converted to utf8 (isBinary)
*/
func isBinary(filePath *string) *bool {
	fileTest, err := os.Open(*filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer fileTest.Close()
	breader := bufio.NewReader(fileTest)
	buf := make([]byte, 4096)
	chunck, err := breader.Read(buf)
	buf = buf[:chunck]
	binTest := !utf8.Valid(buf)
	return &binTest
}
