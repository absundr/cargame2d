package main

import (
	"os"
	"time"

	"github.com/gdamore/tcell"
)

type Game struct {
	Screen tcell.Screen
	Styles Style
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
	render := func () {
		w, h := game.Screen.Size()

		for i := 0; i < w; i++ {
			for j := 0; j < h; j++ {
				game.Screen.Clear()
				game.Screen.SetContent(i, j, 'F', nil, game.Styles.Foreground)
				game.Screen.Show()
				time.Sleep(OneHz * time.Millisecond)
			}
		}
	}

	for {
		render()
		time.Sleep(OneHz * time.Millisecond)	
	}
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