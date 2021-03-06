package xlsx

import (
	"testing"
)

// Test that we can successfully read an XLSX file generated by
// Wps on windows. you can download it freely from http://www.wps.cn/
func TestWpsBlankLine(t *testing.T) {
	xlsxFile, error := OpenFile("wpsBlankLineTest.xlsx")
	if error != nil {
		t.Error(error.Error())
		return
	}
	if xlsxFile == nil {
		t.Error("OpenFile returned nil FileInterface without generating an os.Error")
		return
	}
	sheet := xlsxFile.Sheets[0]
	row := sheet.Rows[0]
	cell := row.Cells[0]
	s := cell.String()
	expected := "编号"
	if s != expected {
		t.Errorf("[TestWpsBlankLine] expected cell A1 = '%s', but got :'%s'", expected, s)
		return
	}
	row = sheet.Rows[2]
	cell = row.Cells[0]
	s = cell.String()
	if s != expected {
		t.Errorf("[TestWpsBlankLine] expected cell A3 = '%s', but got :'%s'", expected, s)
		return
	}

	row = sheet.Rows[4]
	cell = row.Cells[1]
	s = cell.String()
	if s != "" {
		t.Errorf("[TestWpsBlankLine] expected cell B5 = '%s', but got :'%s'", "", s)
		return
	}
	s = sheet.Rows[4].Cells[2].String()
	if s != expected {
		t.Errorf("[TestWpsBlankLine] expected cell C5 = '%s', but got :'%s'", expected, s)
		return
	}
}
