package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dictionary := map[string]string{
		"a": "Alfa",
		"b": "Bravo",
		"c": "Charlie",
		"d": "Delta",
		"e": "Echo",
		"f": "Foxtrot",
		"g": "Golf",
		"h": "Hotel",
		"i": "India",
		"j": "Juliett",
		"k": "Kilo",
		"l": "Lima",
		"m": "Mike",
		"n": "November",
		"o": "Oscar",
		"p": "Papa",
		"q": "Quebec",
		"r": "Romeo",
		"s": "Sierra",
		"t": "Tango",
		"u": "Uniform",
		"v": "Victor",
		"w": "Whiskey",
		"x": "X-ray",
		"y": "Yankee",
		"z": "Zulu",
		"1": "One",
		"2": "Two",
		"3": "Three",
		"4": "Four",
		"5": "Five",
		"6": "Six",
		"7": "Seven",
		"8": "Eight",
		"9": "Nine",
		"0": "Zero",
	}

	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		fmt.Println("Usage: nato <sentence>")
		return
	}
	// Join all arguments together, ignoring whitespace between them
	sentence := strings.Join(argsWithoutProg, "")

	result := []string{}

	for _, char := range strings.Split(sentence, "") {
		// Convert to lowercase for case-insensitive matching
		charLower := strings.ToLower(char)

		// Check if the character exists as a key in the dictionary
		if word, exists := dictionary[charLower]; exists {
			// fmt.Printf("%s: %s\n", char, word)
			result = append(result, word)
		} else if char == " " {
			// Handle spaces
			fmt.Println()
		} else {
			// Character not found in dictionary (e.g., punctuation)
			fmt.Printf("%s: (not found)\n", char)
		}
	}

	fmt.Println(strings.Join(result, " "))
}
