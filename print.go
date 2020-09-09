package prompt

import (
	"fmt"
	"strings"
)

func printPrompt(message string, suffix ...string) {
	if len(suffix) == 0 {
		fmt.Printf("%s: ", message)
	}

	s := strings.Join(suffix, " ")
	s = strings.TrimSpace(s)
	fmt.Printf("%s %s: ", message, s)
}

func printlnPrompt(message string, suffix ...string) {
	if len(suffix) == 0 {
		fmt.Printf("%s:\n", message)
	}

	s := strings.Join(suffix, " ")
	s = strings.TrimSpace(s)
	fmt.Printf("%s %s:\n", message, s)
}
