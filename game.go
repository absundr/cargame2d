package main

import (
	"math/rand"
	"os"
	"time"

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
	// Add incoming cars 
	go func () {
		for {
			// for i, c := range game.IncomingCars {
			// 	Log("Car no: " + strconv.Itoa(i))
			// 	Log("xStart: " + strconv.Itoa(game.Boundaries.BoundaryXStart))				
			// 	Log("xEnd: " + strconv.Itoa(game.Boundaries.BoundaryXEnd))				
			// 	Log("yStart: " + strconv.Itoa(game.Boundaries.BoundaryYStart))				
			// 	Log("yEnd: " + strconv.Itoa(game.Boundaries.BoundaryYEnd))
			// 	Log("-----")
			// 	Log("xCarStart: " + strconv.Itoa(c.Body[0].PosX))				
			// 	Log("xCarEnd: " + strconv.Itoa(c.Body[len(c.Body)-1].PosX))				
			// 	Log("yCarStart: " + strconv.Itoa(c.Body[0].PosY))				
			// 	Log("xCarEnd: " + strconv.Itoa(c.Body[len(c.Body)-1].PosY))				
			// 	Log("***************************")
			// }
			var car Car
			car.InitIncomingCar(game.Lanes, game.Boundaries.BoundaryYStart)
			game.IncomingCars = append(game.IncomingCars, car)
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		}
	}()

	// Handle keyboard input
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
	for {
		// Draw incoming cars
		for _, car := range game.IncomingCars {
			car.DrawCar(game.Screen, game.Styles.Foreground)
		}

		// Display
		game.Screen.Show()

		// Clear incoming cars
		for i := range game.IncomingCars {
			game.IncomingCars[i].ClearCarPos(game.Screen, game.Styles.Background)
		}

		// Update incoming car pos
		// TODO: 
		// This is updating the state and technically not drawing anything
		// Ideally it should exist in the update func 
		// but running them in parallel isn't working as expected
		for i := range game.IncomingCars {
			for j := range game.IncomingCars[i].Body {
				game.IncomingCars[i].Body[j].PosY++ 
			}
		}

		time.Sleep(OneHz*3*time.Millisecond)
	}
}

func (game *Game) New() *Game {
	game.Screen = InitScreen()
	game.Styles = InitStyles()

	game.Screen.Clear()
	game.Screen.SetStyle(game.Styles.Background)
	game.DrawMap()
	
	var car Car
	car.InitCar(game.Lanes, game.Boundaries.BoundaryYEnd)
	car.DrawCar(game.Screen, game.Styles.Foreground)
	game.Car = car

	game.IncomingCars = make([]Car, 0)

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