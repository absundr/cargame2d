## Objective
A console-based simple car game writen in Go. The goal is to avoid oncoming traffic and try to survive on the road for as long as possible.

## Prerequisites
To run the game, make sure you have the following installed on your system:

Golang (version 1.16 or later)
Git (optional, for cloning the repository)

## Getting Started
To get started, follow these steps:

1. Clone the repository or download the source code:
```
git clone https://github.com/absundr/cargame2d.git
```
2. Navigate to the project directory:
```
cd cargame2d
```
3. Build the game:
```
go build
```
4. Run the game:
```
./cargame2d
```

## Gameplay Instructions
Use the arrow keys to control the car.
Avoid colliding with obstacles and other cars.

## Acknowledgements
This project was made possible by using `https://github.com/gdamore/tcell`, a package to draw cells on the terminal. 

## Final thoughts
This project was made for fun and also as way to learn golang. It makes use of go routines which have proven to be invaluable. It doesn't put huge emphasis on the design. It mainly focuses on how you can build a terminal game without using a game engine and handle most of the operations at a low level.


