package main

import (
	"os"

	"github.com/gdamore/tcell"
)

type Props struct {
	BoundaryXStart int
	BoundaryXEnd int
	BoundaryYStart int
	BoundaryYEnd int
}
type Game struct {
	Screen tcell.Screen
	Styles Style
	Props Props
}

const (
	OneHz = 16
)

func (game *Game) Update() {
	for {
		switch ev := game.Screen.PollEvent().(type) {
		case *tcell.EventResize:
			game.Screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				game.End()
			}		
		}
	}
}

func (game *Game) Draw() {
	game.Screen.Clear()
	game.DrawMap()
	game.Screen.Show()
}

func (game *Game) New() {
	game.Screen = InitScreen()
	game.Styles = InitStyles()
	game.Screen.SetStyle(game.Styles.Background)
}

func (game *Game) Run() {
	go game.Draw()
	game.Update()
}

func (game *Game) End() {
	game.Screen.Fini()
	os.Exit(0)
}