package convert

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"strconv"
	"strings"
)

func New(conf *Config) *Convert {
	return &Convert{
		Config: conf,
	}
}

type Config struct {
	SourceExt       Ext
	SourceBufReader *bufio.Reader
	TargetExt       Ext
	Output          string
	PrintProcessing bool
}

type Convert struct {
	*Config
}

func (cv *Convert) Run(ctx context.Context) error {
	csvReader := csv.NewReader(cv.SourceBufReader)
	xlsxWriter := excelize.NewFile()
	activeIndex := xlsxWriter.NewSheet("sheet1")
	xlsxWriter.SetActiveSheet(activeIndex)
	var row int
	for {
		record, err := csvReader.Read()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		row++
		for idx, val := range record {
			axis := intToXlsxAxis(idx) + strconv.Itoa(row)
			if err = xlsxWriter.SetCellValue("sheet1", axis, val); err != nil {
				return err
			}
		}
	}
	if err := xlsxWriter.SaveAs(cv.Output); err != nil {
		return err
	}
	return nil
}

type Ext uint8

const (
	ExtXls Ext = iota
	ExtXlsx
	ExtCsv
	ExtXml
	ExtJson
)

var extStringMap = map[Ext]string{
	ExtXls:  "xls",
	ExtXlsx: "xlsx",
	ExtCsv:  "csv",
	ExtXml:  "xml",
	ExtJson: "json",
}

func (t *Ext) FormString(ext string) error {
	ext = strings.ToLower(strings.TrimLeft(ext, "."))
	for typ, str := range extStringMap {
		if str == ext {
			*t = typ
			return nil
		}
	}

	return fmt.Errorf("un-supported ext: %s", ext)
}

func (t *Ext) String() string {
	return extStringMap[*t]
}
