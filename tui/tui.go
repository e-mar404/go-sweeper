package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)


func Run() {
	p := tea.NewProgram(model{}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("unexpected error while running program: %v\n", err)
	}
}


