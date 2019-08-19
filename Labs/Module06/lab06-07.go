// Lab 06  Part 7
package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

// Read a file into a []byte slice
// and return a string
func myread(filename string) (text string) {
	rawtext, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	// convert the rawtext into a string
	text = string(rawtext)
	return
}

// Process the string into a list of lowercase words
// with no punctuation
func strproc(instr, punctuation string) []string {
	// take a slice to use as a work area
	str := instr[:]
	// iterate over the ";,." string
	for _, char := range punctuation {
		// convert the rune to a string of length 1
		str = strings.Replace(str, string(char), "", -1)
	}
	// split the string into words and lower case letters
	words := strings.Split(strings.ToLower(str), " ")
	return words
}

// convert the list of words into a frequency of occurence map
func makemap(words []string) (m map[string]int) {
	// create the map
	m = make(map[string]int, 20)
	for _, word := range words {
		// for each word in the text, increment the
		// count for that word
		m[word] += 1
	}
	return
}

// converts one map into another.
func transmap(words map[string]int) (count map[int][]string) {
	// the new map
	count = make(map[int][]string, len(words))
	// for each word, find the count and add to the list
	for word, num := range words {
		count[num] = append(count[num], word)
	}
	return

}

// calculate the final result
func results(count map[int][]string) {
	// create a slice of ints for holding the frequencies
	keys := make([]int, 0)
	// extract all the keys
	for key, _ := range count {
		keys = append(keys, key)
	}
	fmt.Println(keys) // just for info sake
	// sort the keys into descending order
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	// iterate through the map in key order and add up
	// how many words are in the slice in each element
	for _, key := range keys {
		fmt.Println("There is/are", len(count[key]), "word(s) that occur", key, "times")
	}
}

func main() {

	wordlist := strproc(myread("churchill.txt"), ".;,")
	results(transmap(makemap(wordlist)))

}
