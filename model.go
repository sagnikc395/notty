package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

//model that holds ou data structure

const (
	listView uint = iota
	titleView
	bodyView
)

type model struct {
	state uint
	// maybe store in something like sqlite
	// store Store
	// textarea.Model
	// ...
}

func NewModel() model {
	return model{
		state: listView,
	}
}

//bubbletea needs to implement 3 methods -> init, view and update

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}


