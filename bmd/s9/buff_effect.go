package s9

import (
	"bytes"
	"os"

	"github.com/baranovskis/bmd-converter/bmd"
)

type BuffEffect struct {
	bmd.BaseBmdAdapter // inherits methods
}

func (be *BuffEffect) GetStruct() interface{} {
	type bmdStruct struct {
		ID          uint16 `csv:"id"`
		Group       uint8  `csv:"group"`
		ItemGroup   uint8  `csv:"item_group"`
		ItemIndex   uint8  `csv:"item_index"`
		Name        string `csv:"name" size:"50"` // set fixed string size
		State1      uint8  `csv:"state_1"`
		State2      uint8  `csv:"state_2"`
		State3      uint8  `csv:"state_3"`
		Description string `csv:"description" size:"100"` // set fixed string size
	}

	return &bmdStruct{}
}

func (be *BuffEffect) GetVersion(data *[]byte) (res int16, err error) {
	return -1, nil
}

func (be *BuffEffect) GetRowsCount(data *[]byte) (res int32, err error) {
	return bmd.GetRowsCount(data)
}

func (be *BuffEffect) GetCrcValue(data *[]byte) (res int32, err error) {
	return bmd.GetCrcValue(data)
}

func (be *BuffEffect) SetVersion(file *os.File) error {
	return nil
}

func (be *BuffEffect) SetRowsCount(file *os.File, count int) error {
	return bmd.SetRowsCount(file, count)
}

func (be *BuffEffect) SetCrcValue(file *os.File, buff bytes.Buffer) error {
	return bmd.SetCrcValue(file, buff, 0xE2F1)
}
