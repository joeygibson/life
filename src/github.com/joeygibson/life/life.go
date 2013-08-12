package main

import (
	"flag"
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/joeygibson/life/entities"
	"os"
)

var size = flag.Int("s", 20, "Row/Column size")
var iterations = flag.Int("i", -1, "# of iterations")
var help = flag.Bool("h", false, "display help")
var hackerSeed = flag.Bool("H", false, "seed with the hacker emblem")
var sleepTime = flag.Int("w", 500, "milliseconds to sleep between iterations")

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

	board := entities.NewBoard(*size, *sleepTime)
	
	if *hackerSeed {
		board.HackerEmblemSeed()
	} else {
		board.Seed()
	}

	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)
	termbox.HideCursor()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	go func() {
		for {
			entities.EventChan <- termbox.PollEvent()
		}
	}()
	
	board.Play(*iterations)
}
