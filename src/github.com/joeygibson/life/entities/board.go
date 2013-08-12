/*
The MIT License (MIT)

Copyright (c) [year] [fullname]

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package entities

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type Board struct {
	size int
	sleepTime int
	cells [][]Cell
}

var EventChan chan termbox.Event

func (board Board) DisplayBoard() {
	termbox.SetCell(0, 0, '\u250c', termbox.ColorDefault, termbox.ColorDefault)
	for i := 0; i < board.size; i++ {
		termbox.SetCell(i + 1, 0, '\u2500', termbox.ColorDefault, termbox.ColorDefault)
	}
	termbox.SetCell(board.size + 1, 0, '\u2510', termbox.ColorDefault, termbox.ColorDefault)
	
	for i := 0; i < board.size; i++ {
		termbox.SetCell(0, i + 1, '\u2502', termbox.ColorDefault, termbox.ColorDefault)
		for j := 0; j < board.size; j++ {
			termbox.SetCell(i + 1, j + 1, board.cells[i][j].Rune(), termbox.ColorDefault, termbox.ColorDefault)
		}
		termbox.SetCell(board.size + 1, i + 1, '\u2502', termbox.ColorDefault, termbox.ColorDefault)
	}

	termbox.SetCell(0, board.size + 1, '\u2514', termbox.ColorDefault, termbox.ColorDefault)
	for i := 0; i < board.size; i++ {
		termbox.SetCell(i + 1, board.size + 1, '\u2500', termbox.ColorDefault, termbox.ColorDefault)
	}
	termbox.SetCell(board.size + 1, board.size + 1, '\u2518', termbox.ColorDefault, termbox.ColorDefault)
	
	termbox.Flush()
}

func (board Board) GetNeighbors(x, y int) []Cell {
	var cells []Cell

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			newX, newY := x + i, y + j

			if newX >= 0 && newY >= 0 &&
				newX < board.size && newY < board.size &&
				!(newX == x && newY == y) {
				cells = append(cells, board.cells[newX][newY])
			}
		}
	}
	
	return cells
}

func (board *Board) Step() {
	var newCells = createCells(board.size)

	for i := 0; i < board.size; i++ {
		for j := 0; j < board.size; j++ {
			newCells[i][j] = board.cells[i][j].Step(board.GetNeighbors(i, j))
		}
	}

	board.cells = newCells

	board.DisplayBoard()
}

func (board Board) Play(iterations int) {
	if iterations < 0 {
		for {
			board.Step()
			
			select {
				case ev := <- EventChan:
					if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
						return
					}
				default:
					// nothing
			}
			
			board.Sleep()
		}
	} else {
		for i := 0; i < iterations; i++ {
			board.Step()

			select {
				case ev := <- EventChan:
					if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
						return
					}
				default:
					// nothing
			}
			
			board.Sleep()
		}
	}
}

func (board Board) Sleep() {
	time.Sleep(time.Duration(board.sleepTime) * time.Millisecond)
}

func (board *Board) HackerEmblemSeed() {
	board.cells[0][2].SetAlive(true)
	board.cells[1][0].SetAlive(true)
	board.cells[1][2].SetAlive(true)
	board.cells[2][1].SetAlive(true)
	board.cells[2][2].SetAlive(true)
}

func (board *Board) TwoIslandPseudoStillLifeSeed() {
	board.cells[1][0].SetAlive(true)
	board.cells[2][0].SetAlive(true)
	board.cells[4][0].SetAlive(true)
	board.cells[5][0].SetAlive(true)
    board.cells[2][1].SetAlive(true)
    board.cells[4][1].SetAlive(true)
    board.cells[2][2].SetAlive(true)
    board.cells[4][2].SetAlive(true)
    board.cells[1][3].SetAlive(true)
    board.cells[2][3].SetAlive(true)
    board.cells[4][3].SetAlive(true)
    board.cells[5][3].SetAlive(true)
}

// Pseudo-random seeding of the board
func (board *Board) Seed() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	times := 0
	
	for times == 0 {
		times = r.Intn(board.size * board.size)
	}
		
	for t := 0; t < times; t++ {
		i, j := rand.Intn(board.size), rand.Intn(board.size)
		
		board.cells[i][j].SetAlive(true)
	}
}

func createCells(size int) [][]Cell {
	cells := make([][]Cell, size)
	
	for i := 0; i < size; i++ {
		row := make([]Cell, size)
		cells[i] = row
	}

	return cells
}

func NewBoard(size int, sleepTime int) Board {
	EventChan = make(chan termbox.Event)
	var board Board

	board.size = size
	board.sleepTime = sleepTime
	board.cells = createCells(size)

	return board
}

func ClearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
