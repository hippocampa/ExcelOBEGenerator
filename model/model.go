package model

import (
	"github.com/hippocampa/obeobe/cpl"
)

type Model struct {
	sheetName string
	cpl       []*cpl.CPL
}

func NewModel(sheetName string, beginCol int) *Model {
	return &Model{
		sheetName: sheetName,
		cpl:       make([]*cpl.CPL, 0),
	}
}

func (s *Model) AddCPL(c *cpl.CPL) {
	// if cpmk is empty then set the beginCol and endCol to 1
	if len(s.cpl) == 0 {
		c.SetBeginCol(1)
		c.SetEndCol(1)
	} else {
		// set begin col to the previous end col +1
		c.SetBeginCol(s.cpl[len(s.cpl)-1].EndCol() + 1)
	}

	s.cpl = append(s.cpl, c)
}

func (s *Model) GetSheetName() string {
	return s.sheetName
}

// Getters
func (s Model) SheetName() string {
	return s.sheetName
}

func (s Model) CPL() []*cpl.CPL {
	return s.cpl
}

// Setters
func (s *Model) SetSheetName(name string) {
	s.sheetName = name
}

func (s *Model) SetCPL(cpls []*cpl.CPL) {
	s.cpl = cpls
}
