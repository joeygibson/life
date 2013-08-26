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
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
package entities

// A single cell that is either alive or dead.
type Cell struct {
	alive bool
}

// Returns a character indicating whether it's alive or dead.
// It's a rune, because termbox needed a rune, and not a string.
func (c *Cell) Rune() rune {
	if c.Alive() {
		return '*'
	} else {
		return ' '
	}
}	

// Is it alive?
func (c *Cell) Alive() bool {
	return c.alive
}

// Kill it or aliven it.
func (c *Cell) SetAlive(al bool) {
	c.alive = al
}

// This clones the cell, so that during the Step, the existing
// board won't change
func (c Cell) Copy() Cell {
	var newCell Cell

	newCell.SetAlive(c.Alive())

	return newCell
}

// Creates a copy of itself, and then sets it's alive/dead status
// based on the alive/dead status of its neighbors
func (c Cell) Step(neighbors []Cell) Cell {
	var liveCount int
	
	for _, n := range neighbors {
		if n.Alive() {
			liveCount = liveCount + 1
		}
	}

	newCell := c.Copy()
	
	if c.Alive() {
		if liveCount < 2 || liveCount > 3 {
			newCell.SetAlive(false)
		}
	} else {
		if !c.Alive() && liveCount == 3 {
			newCell.SetAlive(true)
		}
	}

	return newCell
}
	
