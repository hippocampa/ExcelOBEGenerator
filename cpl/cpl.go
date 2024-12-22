package cpl

type CPL struct {
	name string
	cpmk []CPMK
}

func NewCPL(name string, cpmk []CPMK) *CPL {
	return &CPL{name, cpmk}
}

// Getters
func (c CPL) Name() string {
	return c.name
}

func (c CPL) CPMK() []CPMK {
	return c.cpmk
}

// Setters
func (c *CPL) SetName(name string) {
	c.name = name
}

func (c *CPL) SetCPMK(cpmk []CPMK) {
	c.cpmk = cpmk
}
