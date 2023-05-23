package main

import "github.com/gdamore/tcell"

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

func (game *Game) InitCar() {
	// Make this random
	lane := 2

	game.UpdateCarPos(lane)
}

func (game *Game) UpdateCarPos(lane int) {
	endY := game.Props.BoundaryYEnd

	carStartX, carStartY := game.Lanes[lane].StartX+laneAdjustmentX, endY-carHeight
	carEndX, carEndY := game.Lanes[lane].EndX-laneAdjustmentX, endY

	body := make([]Cell, 0)
	counter := 0
	for i := carStartX; i < carEndX; i++ {
		for j := carStartY; j < carEndY; j++ {
			body = append(body, Cell{PosX: i, PosY: j})
			counter++
		}
	}

	game.Car = Car{
		Body: body,
		Lane: lane,
	}
}

func (game *Game) ClearCarPos() {
	game.DrawCar(game.Styles.Background)
}

func (game *Game) DrawCar(style tcell.Style) {
	for _, cell := range game.Car.Body {
		game.Screen.SetContent(cell.PosX, cell.PosY, tcell.RuneBlock, nil, style)
	}
}