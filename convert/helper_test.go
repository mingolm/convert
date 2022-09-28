package convert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntToXlsxAxis(t *testing.T) {
	for idx, axis := range map[int]string{
		0:  "A",
		1:  "B",
		25: "Z",
		26: "AA",
		27: "AB",
	} {
		assert.Equal(t, axis, intToXlsxAxis(idx), idx)
	}
}
