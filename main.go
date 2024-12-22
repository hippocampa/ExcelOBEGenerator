package main

import (
	"fmt"

	"github.com/hippocampa/obeobe/cpl"
	model "github.com/hippocampa/obeobe/model"
	"github.com/hippocampa/obeobe/utils/writer"
	"github.com/xuri/excelize/v2"
)

func main() {
	var (
		model1 = model.NewModel("Sheet1", 1)
	)
	var (
		CPMK1 = cpl.New("CPMK1", 100, 100, 100, 100, 100)
		CPMK2 = cpl.New("CPMK2", 100, 0, 100, 0, 100)
		CPMK3 = cpl.New("CPMK3", 100, 0, 100, 0, 100)
	)
	var (
		CPL1 = cpl.NewCPL("CPL1")
		CPL2 = cpl.NewCPL("CPL2")
	)
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	CPL1.AddCPMK(*CPMK1)
	CPL1.AddCPMK(*CPMK2)
	CPL2.AddCPMK(*CPMK3)
	// Create new sheet
	model1.AddCPL(CPL1)
	model1.AddCPL(CPL2)

	// // Write sheet
	if err := writer.WriteSheet(f, model1); err != nil {
		fmt.Println(err)
	}
	if err := writer.SaveToExcel(f); err != nil {
		fmt.Println(err)
	}

}
