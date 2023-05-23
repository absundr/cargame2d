package main

import (
	"math/rand"

	"github.com/gdamore/tcell"
)

type Cell struct {
	PosX int
	PosY int
}

type Car struct {
	Body []Cell
	Lane int
}

const (
	laneAdjustmentX = 2
	laneAdjustmentY = 1
	carHeight       = 7
)

func (car *Car) InitCar(lanes [LaneCount]Lane, endY int) {
	lane := 2

	car.UpdateCarPos(lanes, endY, lane)
}

func (car *Car) InitIncomingCar(lanes [LaneCount]Lane, endY int) {
	lane := rand.Intn(LaneCount-1)

	car.UpdateCarPos(lanes, endY, lane)
}

func (car *Car) UpdateCarPos(lanes [LaneCount]Lane, endY int, lane int) {
	carStartX, carStartY := lanes[lane].StartX+laneAdjustmentX, endY-carHeight
	carEndX, carEndY := lanes[lane].EndX-laneAdjustmentX, endY

	body := make([]Cell, 0)
	counter := 0
	for i := carStartX; i < carEndX; i++ {
		for j := carStartY; j < carEndY; j++ {
			body = append(body, Cell{PosX: i, PosY: j})
			counter++
		}
	}

	car.Body = body
	car.Lane = lane
}

func (car *Car) ClearCarPos(screen tcell.Screen, style tcell.Style) {
	car.DrawCar(screen, style)
}

func (car *Car) DrawCar(screen tcell.Screen, style tcell.Style) {
	for _, cell := range car.Body {
		screen.SetContent(cell.PosX, cell.PosY, tcell.RuneBlock, nil, style)
	}
}
