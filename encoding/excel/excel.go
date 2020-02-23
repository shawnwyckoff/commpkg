package excel

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/extrame/xls"
	"github.com/tealeg/xlsx"
)

type ExcelDoc struct {
	jsonstr string
	xlsx    *xlsx.File
	xls     *xls.WorkBook
}

type Sheet struct {
	xlsxSheet *xlsx.Sheet
	xlsSheet  *xls.WorkSheet
}

type Col struct {
	xlsxCol *xlsx.Col
	xlsCol  *xls.Col
}

type Row struct {
	xlsxCol *xlsx.Row
	xlsCol  *xls.Row
}

type Cell struct {
	xlsxCol *xlsx.Cell
	xlsCol  *xls.CellRange
}

func OpenPath(filename string) (*ExcelDoc, error) {
	return nil, nil
}

func OpenBytes(b []byte) (*ExcelDoc, error) {
	dx, err := xlsx.OpenBinary(b)
	if err == nil {
		return &ExcelDoc{xlsx: dx}, nil

	}

	reader, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		return nil, err
	}
	dx, err = xlsx.ReadZipReader(reader)
	if err != nil {
		return nil, err
	}
	return &ExcelDoc{xlsx: dx}, nil

}

func (d *ExcelDoc) Sheets(idx int) *Sheet {
	r := Sheet{}
	if d.xlsx != nil {
		r.xlsxSheet = d.xlsx.Sheets[idx]
	} else {
		r.xlsSheet = d.xls.GetSheet(idx)
	}
	return &r
}

func (d *ExcelDoc) SheetCount() int {
	if d.xlsx != nil {
		return len(d.xlsx.Sheets)
	} else {
		return d.xls.NumSheets()
	}
}

func (d *ExcelDoc) RowCount(sheetIdx int) int {
	if sheetIdx >= d.SheetCount() {
		return 0
	}

	if d.xlsx != nil {
		return len(d.xlsx.Sheets[sheetIdx].Rows)
	} else {
		return int(d.xls.GetSheet(sheetIdx).MaxRow) + 1
	}
}

func (d *ExcelDoc) CellCount(sheetIdx, rowIdx int) int {
	if sheetIdx >= d.SheetCount() {
		return 0
	}
	if rowIdx >= d.RowCount(rowIdx) {
		return 0
	}

	if d.xlsx != nil {
		return len(d.xlsx.Sheets[sheetIdx].Row(rowIdx).Cells)
	} else {
		return int(d.xls.GetSheet(sheetIdx).Row(rowIdx).LastCol()) + 1
	}
}

func (d *ExcelDoc) GetCell(sheetIdx, rowIdx, colIdx int) (string, bool) {
	if sheetIdx < 0 || rowIdx < 0 || colIdx < 0 {
		return "", false
	}
	if colIdx >= d.CellCount(sheetIdx, rowIdx) {
		fmt.Println(":", d.CellCount(sheetIdx, rowIdx))
		return "", false
	}
	if d.xlsx != nil {
		return d.xlsx.Sheets[sheetIdx].Rows[rowIdx].Cells[colIdx].String(), true
	} else {
		return d.xls.GetSheet(sheetIdx).Row(rowIdx).Col(colIdx), true
	}
}
