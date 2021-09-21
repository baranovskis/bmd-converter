package s9

import (
	"bytes"
	"os"

	"github.com/baranovskis/bmd-converter/bmd"
)

type NpcName struct {
	bmd.BaseBmdAdapter // inherits methods
}

func (nn *NpcName) GetStruct() interface{} {
	type bmdStruct struct {
		ID   uint16 `csv:"id"`
		Type uint16 `csv:"type"`
		Name string `csv:"text" size:"256"` // set fixed string size
	}

	return &bmdStruct{}
}

func (nn *NpcName) GetVersion(data *[]byte) (res int16, err error) {
	return -1, nil
}

func (nn *NpcName) GetRowsCount(data *[]byte) (res int32, err error) {
	return -1, nil
}

func (nn *NpcName) GetCrcValue(data *[]byte) (res int32, err error) {
	return -1, nil
}

func (nn *NpcName) SetVersion(file *os.File) error {
	return nil
}

func (nn *NpcName) SetRowsCount(file *os.File, count int) error {
	return nil
}

func (nn *NpcName) SetCrcValue(file *os.File, buff bytes.Buffer) error {
	return nil
}
