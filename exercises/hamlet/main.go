package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type keyval struct {
	key   string
	value int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s <filename>\n", os.Args[0])
		return
	}

	// First call counts words
	wordcount := countwords(os.Args[1])
	// Second call sorts counts
	sortedwordcount := sortwordcounts(wordcount)

	// Only displays top 30 most used words
	for _, entry := range sortedwordcount[:30] {
		fmt.Println(entry.key, entry.value)
	}
}

func countwords(filename string) map[string]int {
	wordcount := make(map[string]int)

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reg, err := regexp.Compile("[^a-z ]+")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		// Uses regular expression to remove all characters except for a-z and space
		line = reg.ReplaceAllString(line, "")
		words := strings.Split(line, " ")

		for _, word := range words {
			if len(word) > 0 {
				wordcount[word]++
			}
		}
	}

	return wordcount
}

func sortwordcounts(wordcount map[string]int) []keyval {
	var sortedwordcount []keyval

	for word, count := range wordcount {
		sortedwordcount = append(sortedwordcount, keyval{word, count})
	}

	sort.Slice(sortedwordcount,
		func(i, j int) bool {
			return sortedwordcount[i].value > sortedwordcount[j].value
		},
	)

	return sortedwordcount
}
