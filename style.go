package main

import "github.com/gdamore/tcell"

type Style struct {
	Background tcell.Style
	Foreground tcell.Style
	Primary tcell.Style
	Secondary tcell.Style
}

func InitStyles() Style {
	style := Style {
		Background: tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorBlack),
		Foreground: tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite),
		Primary: tcell.StyleDefault.Background(tcell.ColorBlue),
		Secondary: tcell.StyleDefault.Background(tcell.ColorRed),
	}
	return style
}