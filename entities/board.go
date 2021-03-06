/*
The MIT License (MIT)

Copyright (c) 2013 Joey Gibson

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
	"math/rand"
	"time"
)

// The board that the game is played on. Sizes set by user or terminal
type Board struct {
	Rows      int
	Columns   int
	SleepTime int
	Cells     [][]Cell
}

// Given a cell, this function will compute the 8 neighbors
// surrounding the cell. If a neighbord is located off the board
// it will wrap around to the other side.
func (board Board) GetNeighbors(r, c int) []Cell {
	var cells []Cell

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			newR, newC := r+i, c+j

			if newR >= board.Rows {
				newR = 0
			} else if newR < 0 {
				newR = board.Rows - 1
			}

			if newC >= board.Columns {
				newC = 0
			} else if newC < 0 {
				newC = board.Columns - 1
			}

			if !(newR == r && newC == c) {
				cells = append(cells, board.Cells[newR][newC])
			}
		}
	}

	return cells
}

// Advance the board by one generation. A copy of the original board is
// updated and returned.
func (board Board) Step() Board {
	var newCells = createCells(board.Rows, board.Columns)

	for i := 0; i < board.Rows; i++ {
		for j := 0; j < board.Columns; j++ {
			newCells[i][j] = board.Cells[i][j].Step(board.GetNeighbors(i, j))
		}
	}

	board.Cells = newCells

	return board
}

// A seed that places a single Gosper's Glider on the board
func (board *Board) HackerEmblemSeed() {
	board.Cells[0][2].SetAlive(true)
	board.Cells[1][0].SetAlive(true)
	board.Cells[1][2].SetAlive(true)
	board.Cells[2][1].SetAlive(true)
	board.Cells[2][2].SetAlive(true)
}

// Pseudo-random seeding of the board. This may overseed it a bit...
func (board *Board) Seed() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	times := 0

	for times == 0 {
		times = r.Intn(board.Rows * board.Columns)
	}

	for t := 0; t < times; t++ {
		i, j := rand.Intn(board.Rows), rand.Intn(board.Columns)

		board.Cells[i][j].SetAlive(true)
	}
}

// Create the matrix that is the guts of the board.
func createCells(Rows, Columns int) [][]Cell {
	cells := make([][]Cell, Rows)

	for i := 0; i < Rows; i++ {
		row := make([]Cell, Columns)
		cells[i] = row
	}

	return cells
}

// Factory function to create a new board of the given size. The board
// that is returned has not yet been seeded.
func NewBoard(Rows, Columns int) Board {
	var board Board

	board.Rows = Rows
	board.Columns = Columns
	board.Cells = createCells(Rows, Columns)

	return board
}
