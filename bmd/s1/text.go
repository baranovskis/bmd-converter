package s1

import (
	"bytes"
	"os"

	"github.com/baranovskis/bmd-converter/bmd"
)

type Text struct {
	bmd.BaseBmdAdapter // inherits methods
}

func (t *Text) GetStruct() interface{} {
	type bmdStruct struct {
		Text string `csv:"text" size:"300"`
	}

	return &bmdStruct{}
}

func (t *Text) GetVersion(data *[]byte) (res int16, err error) {
	return -1, nil
}

func (t *Text) GetRowsCount(data *[]byte) (res int32, err error) {
	return -1, nil
}

func (t *Text) GetCrcValue(data *[]byte) (res int32, err error) {
	return -1, nil
}

func (t *Text) SetVersion(file *os.File) error {
	return nil
}

func (t *Text) SetRowsCount(file *os.File, count int) error {
	return nil
}

func (t *Text) SetCrcValue(file *os.File, buff bytes.Buffer) error {
	return nil
}
