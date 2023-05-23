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

type Lane struct {
	StartX int
	EndX int
}

const (
	LaneCount int = 4
)
type Game struct {
	Screen tcell.Screen
	Styles Style
	Props Props
	Car Car
	Lanes [LaneCount]Lane
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
			} else if ev.Key() == tcell.KeyRight && game.Car.Lane < LaneCount-1 {
				game.ClearCarPos()
				game.UpdateCarPos(game.Car.Lane+1)
				game.DrawCar(game.Styles.Foreground)
			} else if ev.Key() == tcell.KeyLeft && game.Car.Lane > 0 {
				game.ClearCarPos()
				game.UpdateCarPos(game.Car.Lane-1)
				game.DrawCar(game.Styles.Foreground)
			}		
		}
	}
}

func (game *Game) Draw() {
	game.Screen.Clear()
	game.DrawMap()
	game.InitCar()
	
	carStyle := game.Styles.Foreground
	game.DrawCar(carStyle)
	
	for {
		game.Screen.Show()
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