package s9

import (
	"bytes"
	"os"

	"github.com/baranovskis/bmd-converter/bmd"
)

type Filter struct {
	bmd.BaseBmdAdapter // inherits methods
}

func (f *Filter) GetStruct() interface{} {
	type bmdStruct struct {
		Text string `csv:"text" size:"20"` // set fixed string size
	}

	return &bmdStruct{}
}

func (f *Filter) GetVersion(data *[]byte) (res int16, err error) {
	return -1, nil
}

func (f *Filter) GetRowsCount(data *[]byte) (res int32, err error) {
	return -1, nil
}

func (f *Filter) GetCrcValue(data *[]byte) (res int32, err error) {
	return bmd.GetCrcValue(data)
}

func (f *Filter) SetVersion(file *os.File) error {
	return nil
}

func (f *Filter) SetRowsCount(file *os.File, count int) error {
	return nil
}

func (f *Filter) SetCrcValue(file *os.File, buff bytes.Buffer) error {
	return bmd.SetCrcValue(file, buff, 0x3E7D)
}
