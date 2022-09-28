package main

import (
	"path"
	"strings"
)

func GetExt(file string) string {
	return strings.TrimLeft(path.Ext(path.Base(file)), ".")
}
