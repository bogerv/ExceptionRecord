package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/xuri/excelize"
)

func main() {
	xlsx, err := excelize.OpenFile("./Workbook.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cell := xlsx.GetCellValue("Sheet2", "A2")
	fmt.Println(cell)
	// get sheet index
	index := xlsx.GetSheetIndex("Sheet2")
	// get all the rows in a sheet
	rows := xlsx.GetRows("sheet" + strconv.Itoa(index))
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Println(colCell, "\t")
		}
		fmt.Println()
	}
}
