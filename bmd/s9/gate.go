package s9

import (
	"bytes"
	"os"

	"github.com/baranovskis/bmd-converter/bmd"
)

type Gate struct {
	bmd.BaseBmdAdapter // inherits methods
}

func (g *Gate) GetStruct() interface{} {
	type bmdStruct struct {
		Flag      uint8  `csv:"flag"`
		MapNumber uint8  `csv:"map_number"`
		StartX    uint8  `csv:"start_x"`
		StartY    uint8  `csv:"start_y"`
		EndX      uint8  `csv:"end_x"`
		EndY      uint8  `csv:"end_y"`
		Target    uint16 `csv:"target"`
		Direction uint16 `csv:"direction"`
		MinLevel  uint16 `csv:"min_level"`
		MaxLevel  uint16 `csv:"max_level"`
	}

	return &bmdStruct{}
}

func (g *Gate) GetVersion(data *[]byte) (res int16, err error) {
	return -1, nil
}

func (g *Gate) GetRowsCount(data *[]byte) (res int32, err error) {
	return -1, nil
}

func (g *Gate) GetCrcValue(data *[]byte) (res int32, err error) {
	return -1, nil
}

func (g *Gate) SetVersion(file *os.File) error {
	return nil
}

func (g *Gate) SetRowsCount(file *os.File, count int) error {
	return nil
}

func (g *Gate) SetCrcValue(file *os.File, buff bytes.Buffer) error {
	return nil
}
