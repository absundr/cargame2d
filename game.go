package main

import (
	"os"

	"github.com/gdamore/tcell"
)

type Boundaries struct {
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
	Boundaries Boundaries
	Car Car
	Lanes [LaneCount]Lane
	IncomingCars []Car
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
				game.Car.ClearCarPos(game.Screen, game.Styles.Background)
				game.Car.UpdateCarPos(game.Lanes, game.Boundaries.BoundaryYEnd, game.Car.Lane+1)
				game.Car.DrawCar(game.Screen, game.Styles.Foreground)
			} else if ev.Key() == tcell.KeyLeft && game.Car.Lane > 0 {
				game.Car.ClearCarPos(game.Screen, game.Styles.Background)
				game.Car.UpdateCarPos(game.Lanes, game.Boundaries.BoundaryYEnd, game.Car.Lane-1)
				game.Car.DrawCar(game.Screen, game.Styles.Foreground)
			}		
		}
	}
}

func (game *Game) Draw() {
	game.Screen.Clear()
	game.DrawMap()
	
	var car Car
	car.InitCar(game.Lanes, game.Boundaries.BoundaryYEnd)
	carStyle := game.Styles.Foreground
	car.DrawCar(game.Screen, carStyle)
	game.Car = car
	
	for {
		game.Screen.Show()
	}
}

func (game *Game) New() *Game {
	game.Screen = InitScreen()
	game.Styles = InitStyles()
	game.Screen.SetStyle(game.Styles.Background)
	return game
}

func (game *Game) Run() {
	go game.Draw()
	game.Update()
}

func (game *Game) End() {
	game.Screen.Fini()
	os.Exit(0)
}