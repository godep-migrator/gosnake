package main

import (
	//"github.com/davecgh/go-spew/spew"
	"github.com/nsf/termbox-go"
	"math"
)

const (
	SNAKE_DIRECTION_UP = iota
	SNAKE_DIRECTION_RIGHT
	SNAKE_DIRECTION_DOWN
	SNAKE_DIRECTION_LEFT
)

type Direction byte

func (direction Direction) angle() int {
	return int(direction) * 90
}

type Snake struct {
	direction Direction
	body      Body
}

type Body []Node

type Node struct {
	x int
	y int
}

func (snake *Snake) Move() {
	snake.MoveInScreenSize(ScreenSize{0, 0})
}

func (snake *Snake) MoveInScreenSize(screenSize ScreenSize) {
	// TODO:if eat fruit, don't kill tail here.
	snake.body = snake.body[1:]

	head := snake.head()
	var newHead Node

	// TODO: deal over edge case here.
	switch snake.direction {
	case SNAKE_DIRECTION_RIGHT:
		newHead = Node{x: head.x + 1, y: head.y}
	case SNAKE_DIRECTION_DOWN:
		newHead = Node{x: head.x, y: head.y + 1}
	case SNAKE_DIRECTION_LEFT:
		newHead = Node{x: head.x - 1, y: head.y}
	case SNAKE_DIRECTION_UP:
		newHead = Node{x: head.x, y: head.y - 1}
	}

	if screenSize.width > 0 && newHead.x >= screenSize.width {
		newHead.x = screenSize.width - newHead.x
	}

	snake.body = append(snake.body, newHead)
}

func (snake *Snake) Len() int {
	return len(snake.body)
}

func (snake *Snake) head() Node {
	return snake.body[snake.Len()-1]
}

// TODO: see if there's a way to restrict parameter direction to const
func (snake *Snake) Turn(direction Direction) {
	// You can't turn to opposite direction
	angle := float64(direction.angle() - snake.direction.angle())
	if math.Abs(angle) == 180.0 {
		return
	}

	snake.direction = direction
}

func (snake *Snake) Draw() {
	for _, node := range snake.body {
		termbox.SetCell(node.x, node.y, ' ', termbox.ColorDefault, termbox.ColorRed)
	}
}

func NewSnake() *Snake {
	snake := Snake{}

	// give default
	snake.direction = SNAKE_DIRECTION_RIGHT
	snake.body = Body{Node{x: 0, y: 0}, Node{x: 1, y: 0}}
	return &snake
}
