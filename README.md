Conway's Game of Life in Go
===========================

This is a simple, and probably naïve, implementation of [Conway's Game of Life](http://en.wikipedia.org/wiki/Conway%27s_game_of_life), written in [Go](http://golang.org/). 

The version on `master ` uses NCurses through the [termbox-go](https://github.com/nsf/termbox-go) library. You'll need to fetch it with

# go get github.com/nsf/termbox-go

After that, do a 

# go install github.com/joeygibson/life

and you will have an executable called `life` in the `bin` directory.

