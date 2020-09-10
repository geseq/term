package prompt

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// ErrInvalidInput implies the imput entered was invalid
var ErrInvalidInput = errors.New("invalid input")

var terminal = Terminal

// YesNo prompts the user for a yes/no response.
// In case of an invalid input, it returns ErrInvalidInput.
func YesNo(message string, defaultValue bool) (bool, error) {
	suffix := "(y/N)"
	if defaultValue {
		suffix = "(Y/n)"
	}
	terminal.PrintPrompt(message, suffix)
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		return defaultValue, err
	}

	switch char {
	case 'y', 'Y':
		return true, nil
	case 'n', 'N':
		return false, nil
	case '\n', '\r':
		return defaultValue, nil
	default:
		return false, ErrInvalidInput
	}
}

// YesNoWithRetry prompts the user for a yes/no response.
// In case of an invalid input, it prompts again until a valid response is entered.
func YesNoWithRetry(message string, defaultValue bool) (bool, error) {
	suffix := "(y/N)"
	if defaultValue {
		suffix = "(Y/n)"
	}
	terminal.PrintPrompt(message, suffix)
	for {
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
		if err != nil {
			return defaultValue, err
		}

		switch char {
		case 'y', 'Y':
			return true, nil
		case 'n', 'N':
			return false, nil
		case '\n', '\r':
			return defaultValue, nil
		default:
			terminal.PrintPrompt("Invalid value entered. Please try again", suffix)
		}
	}
}

// Choice prompts the user for a choice from a list of options returning the index and the option string.
// In case of an invalid input, it returns ErrInvalidInput.
func Choice(message string, options []string) (int, string, error) {
	terminal.Println(message)
	for i, item := range options {
		terminal.Println(i+1, item)
	}

	choice, err := strconv.Atoi(Input("Enter your choice"))
	if err != nil {
		return -1, "", ErrInvalidInput
	}

	choice -= 1
	if choice < 0 || choice >= len(options) {
		return -1, "", ErrInvalidInput
	}

	return choice, options[choice], nil
}

// ChoiceWithRetry prompts the user for a choice from a list of options returning the index and the option string.
// In case of an invalid input, it prompts again until a valid response is entered.
func ChoiceWithRetry(message string, options []string) (int, string, error) {
	terminal.Println(message)
	for i, item := range options {
		terminal.Println(i+1, item)
	}

	choice, err := strconv.Atoi(Input("Enter your choice"))
	if err != nil {
		return -1, "", ErrInvalidInput
	}

	choice -= 1
	for choice < 0 || choice >= len(options) {
		choice, err := strconv.Atoi(Input("Invalid value entered. Please try again"))
		if err != nil {
			return -1, "", ErrInvalidInput
		}
		choice -= 1
	}

	return choice, options[choice], nil
}

// Input prompt the user for any user input
func Input(message string) string {
	terminal.PrintPrompt(message)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	return strings.TrimRight(input, "\n\r")
}
