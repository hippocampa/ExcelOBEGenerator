package cpl

import "fmt"

// Column types as constants
const (
	Partisipasi = "Partisipasi"
	Kuis        = "Kuis"
	UTS         = "UTS"
	UAS         = "UAS"
	Project     = "Project"
)

var Columns = []string{Partisipasi, Kuis, UTS, UAS, Project}

type CPMK struct {
	Values map[string]float32
}

func New(partisipasi float32, kuis float32, uts float32, uas float32, project float32) *CPMK {
	return &CPMK{
		Values: map[string]float32{
			Partisipasi: partisipasi,
			Kuis:        kuis,
			UTS:         uts,
			UAS:         uas,
			Project:     project,
		},
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
