package writer

import (
	"fmt"

	"github.com/hippocampa/obeobe/cpl"
	"github.com/hippocampa/obeobe/model"
	"github.com/xuri/excelize/v2"
)

func writeCPMK(f *excelize.File, m *model.Model, cpl *cpl.CPL, styles map[StyleType]int) error {
	colCounter := cpl.BeginCol()
	for i, cpmk := range cpl.CPMK() {
		cpl.CPMK()[i].SetBeginCol(colCounter)
		for k, v := range cpmk.Values {
			if v != 0 {
				cpl.SetEndCol(colCounter)
				colName, _ := excelize.ColumnNumberToName(cpl.EndCol())
				keyLoc := fmt.Sprintf("%s%d", colName, 2)
				f.SetCellValue(m.GetSheetName(), keyLoc, k)
				ApplyStyle(f, styles, StyleBorder, m.GetSheetName(), keyLoc)
				valueCell := fmt.Sprintf("%s%d", colName, 3)
				f.SetCellValue(m.GetSheetName(), valueCell, v/100.0)
				ApplyStyle(f, styles, StylePercentage, m.GetSheetName(), valueCell)
				colCounter++
			}
		}
		cpl.CPMK()[i].SetEndCol(colCounter - 1)

	}
	cpl.SetEndCol(colCounter - 1)
	return nil
}

func merge(f *excelize.File, sheetName string, beginCol, endCol int) error {
	colNameBegin, _ := excelize.ColumnNumberToName(beginCol)
	colNameEnd, _ := excelize.ColumnNumberToName(endCol)
	cellBegin := fmt.Sprintf("%s%d", colNameBegin, 1)
	cellEnd := fmt.Sprintf("%s%d", colNameEnd, 1)
	if err := f.MergeCell(sheetName, cellBegin, cellEnd); err != nil {
		return err

	}
	return nil
}

func writeCPL(f *excelize.File, m *model.Model, styles map[StyleType]int) error {
	var lastEnd int
	for i, cplItem := range m.CPL() {
		if i == 0 {
			cplItem.SetBeginCol(1)
			cplItem.SetEndCol(1)
		} else {
			cplItem.SetBeginCol(lastEnd + 1)
			cplItem.SetEndCol(lastEnd + 1)
		}

		colName, _ := excelize.ColumnNumberToName(cplItem.BeginCol())
		cell := fmt.Sprintf("%s%d", colName, cplItem.Row())
		f.SetCellValue(m.GetSheetName(), cell, cplItem.Name())

		writeCPMK(f, m, cplItem, styles)
		merge(f, m.GetSheetName(), cplItem.BeginCol(), cplItem.EndCol())
		ApplyStyle(f, styles, StyleBorder, m.GetSheetName(), cell)

		lastEnd = cplItem.EndCol()
	}
	return nil
}

func WriteSheet(f *excelize.File, m *model.Model) error {
	_, err := f.NewSheet(m.GetSheetName())
	if err != nil {
		return err
	}
	styles, err := InitStyles(f)
	if err != nil {
		return err
	}
	if err := writeCPL(f, m, styles); err != nil {
		return err
	}
	return nil
}

func SaveToExcel(f *excelize.File) error {
	if err := f.SaveAs("test.xlsx"); err != nil {
		return err
	}
	return nil

}
