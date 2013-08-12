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
	
