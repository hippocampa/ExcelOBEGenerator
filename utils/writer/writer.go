package writer

import (
	"fmt"

	"github.com/hippocampa/obeobe/cpl"
	"github.com/hippocampa/obeobe/model"
	"github.com/xuri/excelize/v2"
)

func writeCPMK(f *excelize.File, m *model.Model, cpl *cpl.CPL) error {
	// TODO! CPMK BeginCol and EndCol
	colCounter := cpl.BeginCol()
	fmt.Println("Begin CPL Column is", colCounter)
	for _, cpmk := range cpl.CPMK() {
		for k, v := range cpmk.Values {
			if v != 0 {
				cpl.SetEndCol(colCounter + 1)
				colName, _ := excelize.ColumnNumberToName(cpl.EndCol())
				keyLoc := fmt.Sprintf("%s%d", colName, 2)
				f.SetCellValue(m.GetSheetName(), keyLoc, k)
				valueCell := fmt.Sprintf("%s%d", colName, 3)
				f.SetCellValue(m.GetSheetName(), valueCell, v/100.0)
				colCounter++

			}
		}
	}
	cpl.SetEndCol(colCounter - 1)
	return nil
}

func writeCPL(f *excelize.File, m *model.Model) error {
	var lastEnd int
	for i, cplItem := range m.CPL() {
		if i == 0 {
			cplItem.SetBeginCol(1)
			cplItem.SetEndCol(0)
		} else {
			cplItem.SetBeginCol(lastEnd + 1)
			cplItem.SetEndCol(lastEnd + 1)
		}

		colName, _ := excelize.ColumnNumberToName(cplItem.BeginCol())
		cell := fmt.Sprintf("%s%d", colName, cplItem.Row())
		f.SetCellValue(m.GetSheetName(), cell, cplItem.Name())

		writeCPMK(f, m, cplItem)

		lastEnd = cplItem.EndCol()
	}
	return nil
}

func WriteSheet(f *excelize.File, m *model.Model) error {
	_, err := f.NewSheet(m.GetSheetName())
	if err != nil {
		return err
	}
	writeCPL(f, m)
	return nil
}

// func writeCPMK(f *excelize.File, m *model.Model, cpl cpl.CPL, styles map[StyleType]int) error {
// 	for _, cpmk := range cpl.CPMK() {
// 		cpKeyRow := 2
// 		cpValueRow := 3

// 		for k, v := range cpmk.Values {
// 			if v != 0 {
// 				colName, _ := excelize.ColumnNumberToName(m.LastCol())

// 				// Write key cell
// 				keyCell := fmt.Sprintf("%s%d", colName, cpKeyRow)
// 				f.SetCellValue(m.GetSheetName(), keyCell, k)
// 				ApplyStyle(f, styles, StyleBorder, m.GetSheetName(), keyCell)

// 				// Write value cell
// 				valueCell := fmt.Sprintf("%s%d", colName, cpValueRow)
// 				f.SetCellValue(m.GetSheetName(), valueCell, v/100.0)
// 				ApplyStyle(f, styles, StylePercentage, m.GetSheetName(), valueCell)

// 				m.SetLastCol(m.LastCol() + 1)
// 			}
// 		}
// 	}
// 	return nil
// }

// func writeCPL(f *excelize.File, m *model.Model) error {
// 	styles, err := InitStyles(f)
// 	if err != nil {
// 		return err
// 	}
// 	for _, cpl := range m.CPL() {
// 		colName, _ := excelize.ColumnNumberToName(m.LastCol())
// 		row := 1
// 		cell := fmt.Sprintf("%s%d", colName, row)
// 		// Write CPL name
// 		f.SetCellValue(m.GetSheetName(), cell, cpl.Name())
// 		// Write CPMK
// 		if err := writeCPMK(f, m, cpl, styles); err != nil {
// 			return err
// 		}

// 	}

// 	return nil
// }

// func WriteSheet(f *excelize.File, m *model.Model) error {
// 	// Set sheet name
// 	_, err := f.NewSheet(m.GetSheetName())
// 	if err != nil {
// 		return err
// 	}
// 	if err := writeCPL(f, m); err != nil {
// 		return err
// 	}
// 	nextColForCPL, _ := excelize.ColumnNumberToName(m.LastCol() + 1)
// 	fmt.Println("Next CPL Column is", nextColForCPL)
// 	return nil

// }

func SaveToExcel(f *excelize.File) error {
	if err := f.SaveAs("test.xlsx"); err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
