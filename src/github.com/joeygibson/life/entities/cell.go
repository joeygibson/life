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
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
package entities

type Cell struct {
	alive bool
}

func (c *Cell) Rune() rune {
	if c.Alive() {
		return '*'
	} else {
		return ' '
	}
}	

func (c *Cell) String() string {
	if c.Alive() {
		return "*"
	} else {
		return " "
	}
}

func (c *Cell) Alive() bool {
	return c.alive
}

func (c *Cell) SetAlive(al bool) {
	c.alive = al
}

func (c Cell) Copy() Cell {
	var newCell Cell

	newCell.SetAlive(c.Alive())

	return newCell
}

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
	
