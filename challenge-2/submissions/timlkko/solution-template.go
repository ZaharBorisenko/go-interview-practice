package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	// Read input from standard input
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()

		// Call the ReverseString function
		output := ReverseString(input)

		// Print the result
		fmt.Println(output)
	}
}

// ReverseString returns the reversed string of s.
func ReverseString(s string) string {
	string_letters := []string{}
	for i := 0; i < len(s); i++ {
	string_letters = append(string_letters, string(s[i]))
	}
	slices.Reverse(string_letters)
	reverse := strings.Join(string_letters, "")

	return reverse
}
