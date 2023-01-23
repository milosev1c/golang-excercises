package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func wordFreq(fp string) map[string]int {
	output := make(map[string]int)
	input, err := os.Open(fp)
	if err != nil {
		os.Exit(1)
	}
	pattern := regexp.MustCompile(`[\p{P}\p{S}]`)
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		word = pattern.ReplaceAllString(word, "")
		word = strings.ToLower(word)
		output[word] += 1
	}
	return output
}

func main() {
	counter := wordFreq("ch4/wordfreq/input.txt")
	keys := make([]string, 0, len(counter))
	for k := range counter {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return counter[keys[i]] > counter[keys[j]]
	})
	for _, key := range keys {
		fmt.Printf("%s: %d\n", key, counter[key])
	}
}
