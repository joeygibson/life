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
package main

import (
	"flag"
	"fmt"
	"github.com/joeygibson/life/entities"
	"github.com/nsf/termbox-go"
	"os"
	"time"
)

var columns = flag.Int("c", -1, "number of columns of board")
var rows = flag.Int("r", -1, "number of rows of board")
var iterations = flag.Int("i", -1, "# of iterations")
var help = flag.Bool("h", false, "display help")
var hackerSeed = flag.Bool("H", false, "seed with the hacker emblem")
var sleepTime = flag.Int("w", 500, "milliseconds to sleep between iterations")
var stopping = false

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage: life [options]\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Usage = Usage
	flag.Parse()

	if *help {
		Usage()
	}
	
	TermboxMain()
}

func TermboxMain() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	// Without this, the termainl screen will be all wonky
	defer termbox.Close()

	// Get screen dimensions from terminal
	termColumns, termRows := termbox.Size()

	// If the user didn't specify a # of rows, use the terminal height
	if *rows < 0 {
		*rows = termRows - 2
	}

	// If the user didn't specify a # of columns, use the terminal width
	if *columns < 0 {
		*columns = termColumns - 2
	}

	board := entities.NewBoard(*rows, *columns)

	// If the user requested the Gosper's Glider seed, use that. Otherwise, go random.
	if *hackerSeed {
		board.HackerEmblemSeed()
	} else {
		board.Seed()
	}

	termbox.SetInputMode(termbox.InputEsc)
	termbox.HideCursor()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// This goroutine polls termbox for keypresses and signals the game
	// to stop when the user presses Esc.
	go func() {
		for {
			ev := termbox.PollEvent()

			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
				stopping = true
				break
			}
		}
	}()

	Play(board, *iterations)
}

// Plays the game. If the user didn't give an iteration count, we will
// run it in an infinite loop
func Play(board entities.Board, iterations int) {
	if iterations < 0 {
		for {
			if stopping {
				return
			}

			board = board.Step()
			displayBoard(board)

			Sleep()
		}
	} else {
		for i := 0; i < iterations; i++ {
			if stopping {
				return
			}

			board = board.Step()
			displayBoard(board)

			Sleep()
		}
	}
}

// This displays the board, using the termbox library. Each time the game's generation
// is advanced, this function will redraw every cell, and the border.
func displayBoard(board entities.Board) {
	termbox.SetCell(0, 0, '\u250c', termbox.ColorDefault, termbox.ColorDefault)
	for j := 0; j < board.Columns; j++ {
		termbox.SetCell(j+1, 0, '\u2500', termbox.ColorDefault, termbox.ColorDefault)
	}
	termbox.SetCell(board.Columns+1, 0, '\u2510', termbox.ColorDefault, termbox.ColorDefault)

	for i := 0; i < board.Rows; i++ {
		termbox.SetCell(0, i+1, '\u2502', termbox.ColorDefault, termbox.ColorDefault)
		for j := 0; j < board.Columns; j++ {
			termbox.SetCell(j+1, i+1, board.Cells[i][j].Rune(), termbox.ColorDefault, termbox.ColorDefault)
		}
		termbox.SetCell(board.Columns+1, i+1, '\u2502', termbox.ColorDefault, termbox.ColorDefault)
	}

	termbox.SetCell(0, board.Rows+1, '\u2514', termbox.ColorDefault, termbox.ColorDefault)
	for i := 0; i < board.Columns; i++ {
		termbox.SetCell(i+1, board.Rows+1, '\u2500', termbox.ColorDefault, termbox.ColorDefault)
	}
	termbox.SetCell(board.Columns+1, board.Rows+1, '\u2518', termbox.ColorDefault, termbox.ColorDefault)

	termbox.Flush()
}

// Time to wait between generations. The default is 500 milliseconds.
func Sleep() {
	time.Sleep(time.Duration(*sleepTime) * time.Millisecond)
}
