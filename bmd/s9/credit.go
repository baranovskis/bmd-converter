package s9

import (
	"bytes"
	"os"

	"github.com/baranovskis/bmd-converter/bmd"
)

type Credit struct {
	bmd.BaseBmdAdapter // inherits methods
}

func (c *Credit) GetStruct() interface{} {
	type bmdStruct struct {
		Type uint8  `csv:"type"`
		Text string `csv:"text" size:"32"` // set fixed string size
	}

	return &bmdStruct{}
}

func (c *Credit) GetVersion(data *[]byte) (res int16, err error) {
	return -1, nil
}

func (c *Credit) GetRowsCount(data *[]byte) (res int32, err error) {
	return -1, nil
}

func (c *Credit) GetCrcValue(data *[]byte) (res int32, err error) {
	return -1, nil
}

func (c *Credit) SetVersion(file *os.File) error {
	return nil
}

func (c *Credit) SetRowsCount(file *os.File, count int) error {
	return nil
}

func (c *Credit) SetCrcValue(file *os.File, buff bytes.Buffer) error {
	return nil
}
