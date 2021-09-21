package s9

import (
	"bytes"
	"os"

	"github.com/baranovskis/bmd-converter/bmd"
)

// Text is ..
type Text struct {
	bmd.BaseBmdAdapter // inherits methods
}

// GetStruct is ...
func (t *Text) GetStruct() interface{} {
	type bmdStruct struct {
		ID   uint32 `csv:"id" xor:"false"` // default xor:"true"
		Size uint32 `csv:"-" xor:"false"`  // default xor:"true"
		Text string `csv:"text" size:"auto"`
	}

	return &bmdStruct{}
}

// GetVersion is return bmd version
func (t *Text) GetVersion(data *[]byte) (res int16, err error) {
	return bmd.GetVersion(data)
}

// GetRowsCount is return rows count
func (t *Text) GetRowsCount(data *[]byte) (res int32, err error) {
	return bmd.GetRowsCount(data)
}

// GetCrcValue is return crc value
func (t *Text) GetCrcValue(data *[]byte) (res int32, err error) {
	return -1, nil
}

// SetVersion is write bmd version into file start
func (t *Text) SetVersion(file *os.File) error {
	return bmd.SetVersion(file, 21575)
}

// SetRowsCount is writec total rows count into file start
func (t *Text) SetRowsCount(file *os.File, count int) error {
	return bmd.SetRowsCount(file, count)
}

// SetCrcValue is write calculated crc value into file end
func (t *Text) SetCrcValue(file *os.File, buff bytes.Buffer) error {
	return nil
}
