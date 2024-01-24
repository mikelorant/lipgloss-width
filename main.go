package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

type Text struct {
	str  string
	bold bool
}

var debug = true

func main() {
	texts := []Text{
		{str: "", bold: false},
		{str: "", bold: true},
		{str: "foo", bold: false},
		{str: "foo", bold: true},
	}

	Render(texts)
}

func Render(texts []Text) {
	for idx, text := range texts {
		fmt.Println("Test:", idx)
		fmt.Println("Lipgloss Width:", Width(text))

		if idx != len(texts)-1 {
			fmt.Println()
		}
	}
}

func Width(txt Text) int {
	style := lipgloss.NewStyle().
		Bold(txt.bold).
		MarginRight(1)

	render := style.Render(txt.str)

	if debug {
		fmt.Println("String:", txt.str)
		fmt.Println("Bold:", txt.bold)
		fmt.Println("Color:", !termenv.EnvNoColor())
		fmt.Println("Prompt Len:", len(txt.str))
		fmt.Println("Style Width:", len(render))
	}

	return lipgloss.Width(render)
}
