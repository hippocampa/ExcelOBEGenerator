package main

import (
	"fmt"

	"github.com/hippocampa/obeobe/cpl"
	model "github.com/hippocampa/obeobe/model"
	"github.com/hippocampa/obeobe/utils/writer"
	"github.com/xuri/excelize/v2"
)

func main() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Create new sheet
	model := model.NewModel("Sheet1", 1)
	// create CPL
	cpl1 := cpl.NewCPL("CPL1", []cpl.CPMK{
		*cpl.New(100, 100, 100, 0, 100)})
	cpl2 := cpl.NewCPL("CPL2", []cpl.CPMK{
		*cpl.New(100, 100, 100, 0, 100)})
	model.AddCPL(*cpl1)
	model.AddCPL(*cpl2)

	// Write sheet
	if err := writer.WriteSheet(f, model); err != nil {
		fmt.Println(err)
	}
	if err := writer.SaveToExcel(f); err != nil {
		fmt.Println(err)
	}

}
