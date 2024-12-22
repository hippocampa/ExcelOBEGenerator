package cpl

type CPL struct {
	name     string
	cpmk     []CPMK
	row      int
	beginCol int
	endCol   int
}

func NewCPL(name string) *CPL {
	return &CPL{
		name:     name,
		cpmk:     make([]CPMK, 0),
		beginCol: 0,
		endCol:   0,
		row:      1,
	}
}

func (c *CPL) AddCPMK(cpmk CPMK) {
	// if cpmk is empty then set the beginCol and endCol to 1
	if len(c.cpmk) == 0 {
		c.beginCol = 1
		c.endCol = 1
	} else {
		// set cpmk beginCol and endCol
		cpmk.SetBeginCol(c.endCol + 1)

	}

	c.cpmk = append(c.cpmk, cpmk)
}

// Getters
func (c CPL) Name() string {
	return c.name
}

func (c CPL) CPMK() []CPMK {
	return c.cpmk
}

func (c CPL) Row() int {
	return c.row
}

func (c CPL) BeginCol() int {
	return c.beginCol
}

func (c CPL) EndCol() int {
	return c.endCol
}

// Setters
func (c *CPL) SetName(name string) {
	c.name = name
}

func (c *CPL) SetCPMK(cpmk []CPMK) {
	c.cpmk = cpmk
}

func (c *CPL) SetRow(row int) {
	c.row = row
}

func (c *CPL) SetBeginCol(col int) {
	c.beginCol = col
}

func (c *CPL) SetEndCol(col int) {
	c.endCol = col
}
