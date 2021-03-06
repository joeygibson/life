# Conway's Game of Life in Go

[![Build Status](https://travis-ci.org/joeygibson/life.svg?branch=master)](https://travis-ci.org/joeygibson/life)

This is a simple, and probably naïve, implementation of [Conway's Game of Life](http://en.wikipedia.org/wiki/Conway%27s_game_of_life), written in [Go](http://golang.org/).

## Dependencies

The version on `master` uses NCurses through the [termbox-go](https://github.com/nsf/termbox-go) 
library. The `go get` command below will fetch it, if you don't already have it.

## Installation

Next, `get` this code and install it

    go get github.com/joeygibson/life

and you will have an executable called `life` in the `$GOPATH/bin` directory.

## Running

Just typing `life` will run a game with all parameters set to their defaults. You can run `life -h` to get a list of options and what their defaults are.
