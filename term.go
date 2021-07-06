package term

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Term interface {
	EnableAltScreenBuffer() Term
	DisableAltScreenBuffer() Term
	ClearScreen() Term
	ClearScreenAndScrollback() Term
	Println(message ...interface{}) Term
	PrintTemp(message ...interface{}) Term
	PrintPrompt(message string, suffix ...string) Term
}

type term struct{}

var terminal = &term{}

// Terminal returns an instance of Term
func Terminal() Term {
	return terminal
}

// Terminal returns a new instance of Term
func NewTerm() Term {
	return &term{}
}

// EnableAltScreenBuffer enables the alternative screen buffer
func (t *term) EnableAltScreenBuffer() Term {
	fmt.Println("\x1b[?1049h")
	return t
}

// DisableAltScreenBuffer disables the alternative screen buffer
func (t *term) DisableAltScreenBuffer() Term {
	fmt.Println("\x1b[?1049l")
	return t
}

// ClearScreen clears the entire screen
func (t *term) ClearScreen() Term {
	fmt.Println("\x1b[2J")
	return t
}

// ClearScreenAndScrollback clears the entire screen and deletes all
// lines saved in the scrollback buffer
func (t *term) ClearScreenAndScrollback() Term {
	fmt.Println("\x1b[3J")
	return t
}

// Println prints the message to the screen
func (t *term) Println(message ...interface{}) Term {
	fmt.Println(message...)
	return t
}

// PrintTemp prints a temporary message in the alternative screen buffer
// which will be cleared on hitting [ENTER]
func (t *term) PrintTemp(message ...interface{}) Term {
	t.EnableAltScreenBuffer().
		ClearScreen().
		ClearScreenAndScrollback().
		Println("").
		Println("The following message is temporary and will be cleared when you hit [ENTER]").
		Println("").
		Println(message...).
		Println("").
		PrintPrompt("Press [enter] to return to your shell")

	_, _ = bufio.NewReader(os.Stdin).ReadString('\n')

	t.ClearScreen().
		ClearScreenAndScrollback().
		DisableAltScreenBuffer()

	return t
}

// PrintPrompt prints a prompt message along with the provided suffixes
func (t *term) PrintPrompt(message string, suffix ...string) Term {
	if len(suffix) == 0 {
		fmt.Printf("%s: ", message)
		return t
	}

	s := strings.Join(suffix, " ")
	s = strings.TrimSpace(s)
	fmt.Printf("%s %s: ", message, s)

	return t
}
