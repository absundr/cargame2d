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
	OperationQueue Queue
}

const (
	OneHz = 16
)

func (game *Game) Update() {
	// Add incoming cars 
	go func () {
		for {
			var car Car
			car.InitIncomingCar(game.Lanes, game.Boundaries.BoundaryYStart)
			game.IncomingCars = append(game.IncomingCars, car)
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		}
	}()

	// Update incoming car pos
	go func () {
		for {
			// Push the update operation into the queue
			game.OperationQueue.Enqueue(
				func () {UpdateIncomingCars(game.IncomingCars)}, "Update")
			time.Sleep(OneHz*3*time.Millisecond)
		}
	} ()

	// Process and execute operations in the queue 
	go func () {
		for {
			if len(game.OperationQueue.q) > 3 {
				m := make(map[string]func())
				for i := 0; i < 4; i++ {
					n, e := game.OperationQueue.Dequeue() 
					if e == nil {
						m[n.t] = n.fn
					}
				}
				// Maintains the order in which the operations are executed
				if m["Draw"] != nil {
					m["Draw"]()
				}				
				if m["Show"] != nil {
					m["Show"]()
				}				
				if m["Clear"] != nil && m["Update"] != nil {
					m["Clear"]()
					m["Update"]()
				}				
			}
		}
	} ()

	// Handle keyboard input
	for {
		switch ev := game.Screen.PollEvent().(type) {
		case *tcell.EventResize:
			game.Sync()
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
		// Push draw operation into the queue
		game.OperationQueue.Enqueue(
			func () {DrawIncomingCars(game.Screen, game.IncomingCars, game.Styles.Foreground)}, "Draw")

		// Push the show operation into the queue
		game.OperationQueue.Enqueue(
			func () {game.Screen.Show()}, "Show")

		// Push the clear operation into the queue
		game.OperationQueue.Enqueue(
			func () { ClearIncomingCars(game.Screen, game.IncomingCars, game.Styles.Background)}, "Clear")

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
	game.OperationQueue = *NewQueue()

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

func (game *Game) Sync() {
	game.Screen.Sync()
	game.Screen.Clear()
	game.OperationQueue = *NewQueue()
	game.DrawMap()
	game.Car.UpdateCarPos(game.Lanes, game.Boundaries.BoundaryYEnd, game.Car.Lane)
	game.Car.DrawCar(game.Screen, game.Styles.Foreground)
	for i := range game.IncomingCars {
		game.IncomingCars[i].UpdateCarPos(game.Lanes, game.IncomingCars[i].Body[len(game.IncomingCars[i].Body)-1].PosY, game.IncomingCars[i].Lane)
	}
}
