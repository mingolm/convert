package typ

import (
	"bufio"
	"context"
	"fmt"
	"strings"
)

type Converter interface {
	Run(context.Context) error
}

type Convert struct {
	*Config
}

type Config struct {
	SourceExt       Ext
	SourceBufReader *bufio.Reader
	TargetExt       Ext
	PrintProcessing bool
}

func New(conf *Config) *Convert {
	return &Convert{
		conf,
	}
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
