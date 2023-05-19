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
