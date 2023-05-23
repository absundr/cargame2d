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

	game.Boundaries.BoundaryXStart = posX1 - 1
	game.Boundaries.BoundaryXEnd = endX1
	game.Boundaries.BoundaryYStart = 0
	game.Boundaries.BoundaryYEnd = h - 1

	laneW := (game.Boundaries.BoundaryXEnd-game.Boundaries.BoundaryXStart)/4
	game.Lanes = [LaneCount]Lane{
		{
			StartX: game.Boundaries.BoundaryXStart+2,
			EndX: game.Boundaries.BoundaryXStart+laneW+1,
		},
		{
			StartX: game.Boundaries.BoundaryXStart+laneW+2,
			EndX: game.Boundaries.BoundaryXStart+laneW*2+1,
		},
		{
			StartX: game.Boundaries.BoundaryXStart+laneW*2+2,
			EndX: game.Boundaries.BoundaryXStart+laneW*3+1,
		},
		{
			StartX: game.Boundaries.BoundaryXStart+laneW*3+2,
			EndX: game.Boundaries.BoundaryXStart+laneW*4+1,
		},
	}
}