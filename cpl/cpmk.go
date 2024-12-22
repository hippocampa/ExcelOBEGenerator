package cpl

import "fmt"

// Column types as constants
const (
	Partisipasi = "Partisipasi"
	Kuis        = "Kuis"
	UTS         = "UTS"
	UAS         = "UAS"
	Project     = "Project"
	Tugas       = "Tugas"
)

var Columns = []string{Partisipasi, Kuis, UTS, UAS, Project, Tugas}

type CPMK struct {
	Name     string
	Values   map[string]float32
	BeginCol int
	EndCol   int
	Row      int
}

// Setter for BeginCol
func (c *CPMK) SetBeginCol(beginCol int) {
	c.BeginCol = beginCol
}

// Getter for BeginCol
func (c *CPMK) GetBeginCol() int {
	return c.BeginCol
}

// Setter for EndCol
func (c *CPMK) SetEndCol(endCol int) {
	c.EndCol = endCol
}

// Getter for EndCol
func (c *CPMK) GetEndCol() int {
	return c.EndCol
}

// Setter for Row
func (c *CPMK) SetRow(row int) {
	c.Row = row
}

// Getter for Row
func (c *CPMK) GetRow() int {
	return c.Row
}

func New(name string, partisipasi float32, kuis float32, uts float32, uas float32, project float32, tugas float32) *CPMK {
	return &CPMK{
		Name: name,
		Values: map[string]float32{
			Partisipasi: partisipasi,
			Kuis:        kuis,
			UTS:         uts,
			UAS:         uas,
			Project:     project,
			Tugas:       tugas,
		},
		BeginCol: 1,
		EndCol:   1,
		Row:      2,
	}
}

func (c *CPMK) SetValue(column string, value float32) error {
	if _, exists := c.Values[column]; !exists {
		return fmt.Errorf("invalid column: %s", column)
	}
	c.Values[column] = value
	return nil
}

func (c *CPMK) GetValue(column string) (float32, error) {
	if value, exists := c.Values[column]; exists {
		return value, nil
	}
	return 0, fmt.Errorf("invalid column: %s", column)
}
