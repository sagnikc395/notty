package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

//model that holds our data structure

// define the states of our models as enums
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
	// ... (some other models)
	store *Store
	notes []Note
	//current note that we edit
	// to click and see current note
	currentNote Note
	listIndex   int
}

func NewModel(store *Store) model {
	notes, err := store.GetNotes()
	if err != nil {
		log.Fatalf("unable to get notes: %v", err)
	}
	return model{
		state: listView,
		store: store,
		notes: notes,
	}
}

//bubbletea needs to implement 3 methods -> init, view and update

func (m model) Init() tea.Cmd {
	return nil
}

// receives a message and returns an updated model and the command
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//switch on key press by checking the type of message
	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String() // up,down,ctrlc,q etc.
		switch m.state {
		case listView:
			switch key {
			case "q":
				return m, tea.Quit
			case "n":
				//keybinding for creating new update
				//model update
				//state change
				m.state = titleView
				//... show the input
			case "up", "k":
				if m.listIndex > 0 {
					m.listIndex--
				}
			case "down", "j":
				if m.listIndex < len(m.notes)-1 {
					m.listIndex++
				}
			case "enter":
				m.currentNote = m.notes[m.listIndex]
				m.state = bodyView
				//show the textarea

			}

		}
	}
	return m, nil
}
