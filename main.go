package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 320
	screenHeight = 240
	tileSize     = 5
)

func main() {
	// Set starting values
	rand.Seed(time.Now().UnixNano())
	game := &Game{
		snake:    NewSnake(),
		food:     NewFood(),
		gameOver: false,
		ticks:    0,
		speed:    10,
	}
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Snaykish")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

// Point represents x and y coordinates
type Point struct {
	X int
	Y int
}

// Snake represents the snake body, direction, and growth counter
type Snake struct {
	Body          []Point
	Direction     Point
	GrowthCounter int
}
