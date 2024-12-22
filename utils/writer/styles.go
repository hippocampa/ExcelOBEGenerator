package writer

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type StyleType int

const (
	StyleBorder StyleType = iota
	StylePercentage
)

// Initialize styles once and return the map of style IDs
func InitStyles(f *excelize.File) (map[StyleType]int, error) {
	styles := make(map[StyleType]int)

	// Basic border style
	borderStyle, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "top", Color: "000000", Style: 1},
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
	})
	if err != nil {
		return nil, err
	}
	styles[StyleBorder] = borderStyle

	// Percentage with border style
	percentStyle, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "top", Color: "000000", Style: 1},
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
		NumFmt: 9,
	})
	if err != nil {
		return nil, err
	}
	styles[StylePercentage] = percentStyle

	return styles, nil
}

func ApplyStyle(f *excelize.File, styles map[StyleType]int, styleType StyleType, sheet, cell string) error {
	style, exists := styles[styleType]
	if !exists {
		return fmt.Errorf("style %v not found", styleType)
	}
	return f.SetCellStyle(sheet, cell, cell, style)
}
