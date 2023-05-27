package main

import "github.com/gdamore/tcell"

func DrawIncomingCars(screen tcell.Screen, cars []Car, style tcell.Style) {
	for _, car := range cars {
		car.DrawCar(screen, style)
	}
}

func ClearIncomingCars(screen tcell.Screen, cars []Car, style tcell.Style) {
	for _, car := range cars {
		car.DrawCar(screen, style)
	}
}

func UpdateIncomingCars(cars []Car) {
	for i := range cars {
		for j := range cars[i].Body {
			cars[i].Body[j].PosY++ 
		}
	}
}

