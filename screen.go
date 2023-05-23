package main

import (
	"log"

	"github.com/gdamore/tcell"
)

func InitScreen() tcell.Screen {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%v", err)
	}
	
	if err = screen.Init(); err != nil {
		log.Fatalf("%v", err)
	}

	return screen
}

func (game *Game) DrawMap() {
	w, h := game.Screen.Size()
	
	// Draw boundaries
	endX1, endX2 := w/4, w
	posX1, posX2 := 0, w-endX1

	for i := posX1; i < endX1; i++ {
		for j := 0; j < h; j++ {
			game.Screen.SetContent(i, j, tcell.RuneBlock, nil, game.Styles.Foreground);
		}
	}
	
	for i := posX2; i < endX2; i++ {
		for j := 0; j < h; j++ {
			game.Screen.SetContent(i, j, tcell.RuneBlock, nil, game.Styles.Foreground);
		}
	}

	// Draw road
	endX1, posX1 = posX2, endX1
	for i := posX1; i < endX1; i++ {
		for j := 0; j < h; j++ {
			game.Screen.SetContent(i, j, tcell.RuneBlock, nil, game.Styles.Background);
		}
	}
	
	// Draw lanes
	roadWidth := endX1-posX1
	laneWidth := roadWidth/4
	for i := posX1+laneWidth; i < endX1; i += laneWidth {
		for j := 0; j < h; j++ {
			game.Screen.SetContent(i, j, tcell.RuneVLine, nil, game.Styles.Foreground);
		}
	}

	game.Props.BoundaryXStart = posX1 - 1
	game.Props.BoundaryXEnd = endX1
	game.Props.BoundaryYStart = 0
	game.Props.BoundaryYEnd = h - 1

	laneW := (game.Props.BoundaryXEnd-game.Props.BoundaryXStart)/4
	game.Lanes = [LaneCount]Lane{
		{
			StartX: game.Props.BoundaryXStart+2,
			EndX: game.Props.BoundaryXStart+laneW+1,
		},
		{
			StartX: game.Props.BoundaryXStart+laneW+2,
			EndX: game.Props.BoundaryXStart+laneW*2+1,
		},
		{
			StartX: game.Props.BoundaryXStart+laneW*2+2,
			EndX: game.Props.BoundaryXStart+laneW*3+1,
		},
		{
			StartX: game.Props.BoundaryXStart+laneW*3+2,
			EndX: game.Props.BoundaryXStart+laneW*4+1,
		},
	}
}