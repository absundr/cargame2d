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
	roadWidth := endX1 - posX1

	for i := posX1+roadWidth/4; i < endX1-roadWidth/4; i += roadWidth/4 {
		for j := 0; j < h; j++ {
			game.Screen.SetContent(i, j, tcell.RuneVLine, nil, game.Styles.Foreground);
		}
	} 
}
