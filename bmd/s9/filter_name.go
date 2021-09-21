package s9

import (
	"bytes"
	"os"

	"github.com/baranovskis/bmd-converter/bmd"
)

type FilterName struct {
	bmd.BaseBmdAdapter // inherits methods
}

func (fn *FilterName) GetStruct() interface{} {
	type bmdStruct struct {
		Text string `csv:"text" size:"20"` // set fixed string size
	}

	return &bmdStruct{}
}

func (fn *FilterName) GetVersion(data *[]byte) (res int16, err error) {
	return -1, nil
}

func (fn *FilterName) GetRowsCount(data *[]byte) (res int32, err error) {
	return -1, nil
}

func (fn *FilterName) GetCrcValue(data *[]byte) (res int32, err error) {
	return bmd.GetCrcValue(data)
}

func (fn *FilterName) SetVersion(file *os.File) error {
	return nil
}

func (fn *FilterName) SetRowsCount(file *os.File, count int) error {
	return nil
}

func (fn *FilterName) SetCrcValue(file *os.File, buff bytes.Buffer) error {
	return bmd.SetCrcValue(file, buff, 0x2BC1)
}
