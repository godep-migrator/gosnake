package main

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

type (
	Color           termbox.Attribute
	CharacterStatus byte
)

const (
	COLOR_CHARACTER Color = Color(termbox.ColorRed)
	COLOR_FRUIT     Color = Color(termbox.ColorGreen)
)

const (
	CHARACTER_STATUS_MOVE CharacterStatus = iota
	CHARACTER_STATUS_GROW
	CHARACTER_STATUS_DEAD
)

type Character interface {
	Move(screenSize ScreenSize, fruit Node) CharacterStatus
	Turn(direction Direction)
	Draw()
	Body() Body
}

type ScreenSize struct {
	width  int
	height int
}

type Scene struct {
	character Character
	size      ScreenSize
	fruit     Node
}

func (scene *Scene) SetSize(width int, height int) {
	scene.size = ScreenSize{width: width, height: height}
}

func (scene *Scene) Draw() (stop bool) {
	switch scene.character.Move(scene.size, scene.fruit) {
	case CHARACTER_STATUS_GROW:
		scene.generateFruit()
	case CHARACTER_STATUS_DEAD:
		stop = true
	}

	ClearScene()
	scene.drawFruit()
	scene.character.Draw()

	return stop
}

func (scene *Scene) availableNodes() (availableNodes []Node) {
	for x := 0; x < scene.size.width; x++ {
		for y := 0; y < scene.size.height; y++ {
			node := Node{x: x, y: y}
			if !scene.character.Body().Contains(node) {
				availableNodes = append(availableNodes, node)
			}
		}
	}

	return availableNodes
}

func (scene *Scene) randomAvailableNode() Node {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	nodes := scene.availableNodes()
	return nodes[r.Intn(len(nodes))]
}

func (scene *Scene) generateFruit() {
	node := scene.randomAvailableNode()
	scene.fruit = node
}

func (scene *Scene) drawFruit() {
	DrawPoint(scene.fruit.x, scene.fruit.y, COLOR_FRUIT)
}

func NewScene(character Character, screenSize ScreenSize) *Scene {
	scene := Scene{size: screenSize, character: character}
	scene.generateFruit()
	return &scene
}
