package s9

import (
	"bytes"
	"os"

	"github.com/baranovskis/bmd-converter/bmd"
)

type ItemTooltipText struct {
	bmd.BaseBmdAdapter // inherits methods
}

func (it *ItemTooltipText) GetStruct() interface{} {
	type bmdStruct struct {
		ID   uint16 `csv:"id"`
		Text string `csv:"text" size:"256"` // set fixed string size
		Unk1 int16  `csv:"unk_1"`
	}

	return &bmdStruct{}
}

func (it *ItemTooltipText) GetVersion(data *[]byte) (res int16, err error) {
	return -1, nil
}

func (it *ItemTooltipText) GetRowsCount(data *[]byte) (res int32, err error) {
	return -1, nil
}

func (it *ItemTooltipText) GetCrcValue(data *[]byte) (res int32, err error) {
	return bmd.GetCrcValue(data)
}

func (it *ItemTooltipText) SetVersion(file *os.File) error {
	return nil
}

func (it *ItemTooltipText) SetRowsCount(file *os.File, count int) error {
	return nil
}

func (it *ItemTooltipText) SetCrcValue(file *os.File, buff bytes.Buffer) error {
	return bmd.SetCrcValue(file, buff, 0xE2F1)
}
