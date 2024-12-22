package model

import (
	"github.com/hippocampa/obeobe/cpl"
)

type Model struct {
	sheetName string
	beginCol  int
	lastCol   int
	cpl       []cpl.CPL
}

func NewModel(sheetName string, beginCol int) *Model {
	return &Model{
		sheetName: sheetName,
		beginCol:  beginCol,
		lastCol:   beginCol,
		cpl:       make([]cpl.CPL, 0),
	}
}

func (s *Model) AddCPL(cpl cpl.CPL) {
	s.cpl = append(s.cpl, cpl)
}

func (s *Model) GetSheetName() string {
	return s.sheetName
}

// Getters
func (s Model) SheetName() string {
	return s.sheetName
}

func (s Model) BeginCol() int {
	return s.beginCol
}

func (s Model) LastCol() int {
	return s.lastCol
}

func (s Model) CPL() []cpl.CPL {
	return s.cpl
}

// Setters
func (s *Model) SetSheetName(name string) {
	s.sheetName = name
}

func (s *Model) SetBeginCol(col int) {
	s.beginCol = col
}

func (s *Model) SetLastCol(col int) {
	s.lastCol = col
}

func (s *Model) SetCPL(cpls []cpl.CPL) {
	s.cpl = cpls
}
