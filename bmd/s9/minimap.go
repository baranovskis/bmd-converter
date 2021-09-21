package s9

import (
	"bytes"
	"os"

	"github.com/baranovskis/bmd-converter/bmd"
)

type MiniMap struct {
	bmd.BaseBmdAdapter // inherits methods
}

func (m *MiniMap) GetStruct() interface{} {
	type bmdStruct struct {
		Flag uint32 `csv:"flag"`
		X    uint32 `csv:"x"`
		Y    uint32 `csv:"y"`
		Dir  uint32 `csv:"dir"`
		Text string `csv:"text" size:"32"`
		Idk  uint32 `csv:"idk"`
	}

	return &bmdStruct{}
}

func (m *MiniMap) GetVersion(data *[]byte) (res int16, err error) {
	return -1, nil
}

func (m *MiniMap) GetRowsCount(data *[]byte) (res int32, err error) {
	return -1, nil
}

func (m *MiniMap) GetCrcValue(data *[]byte) (res int32, err error) {
	return bmd.GetCrcValue(data)
}

func (m *MiniMap) SetVersion(file *os.File) error {
	return nil
}

func (m *MiniMap) SetRowsCount(file *os.File, count int) error {
	return nil
}

func (m *MiniMap) SetCrcValue(file *os.File, buff bytes.Buffer) error {
	return bmd.SetCrcValue(file, buff, 0xE2F1)
}
