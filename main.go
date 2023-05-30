package main

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if _, err := tea.Program(struct{}{}); err != nil {
		os.Exit(1)
	}
}
