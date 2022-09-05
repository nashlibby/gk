package gk

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"net/http"
	"time"
)

type ExcelExport struct {
	FileName string      `json:"file_name"`
	File     *xlsx.File  `json:"file"`
	Sheet    *xlsx.Sheet `json:"sheet"`
}

func NewExport(fileName, sheetName string) *ExcelExport {
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet(sheetName)
	return &ExcelExport{FileName: fileName, File: file, Sheet: sheet}
}

func (e *ExcelExport) SetCol(f func(*xlsx.Sheet)) *ExcelExport {
	f(e.Sheet)
	return e
}

func (e *ExcelExport) Header(head []string) *ExcelExport {
	row := e.Sheet.AddRow()
	for _, v := range head {
		cell := row.AddCell()
		cell.Value = v
	}
	return e
}

func (e *ExcelExport) Data(data [][]string) *ExcelExport {
	if len(data) > 0 {
		for _, d := range data {
			row := e.Sheet.AddRow()
			for _, v := range d {
				cell := row.AddCell()
				cell.Value = v
			}
		}
	}
	return e
}

func (e *ExcelExport) DownloadWithGin(c *gin.Context) {
	var buffer bytes.Buffer
	_ = e.File.Write(&buffer)
	content := bytes.NewReader(buffer.Bytes())
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, e.FileName))
	c.Writer.Header().Add("Content-Type", "application/x-excel")
	http.ServeContent(c.Writer, c.Request, e.FileName, time.Now(), content)
	return
}
