package s1

import (
	"bytes"
	"os"

	"github.com/baranovskis/bmd-converter/bmd"
)

type MoveReq struct {
	bmd.BaseBmdAdapter // inherits methods
}

func (m *MoveReq) GetStruct() interface{} {
	type bmdStruct struct {
		ID         uint32 `csv:"id"`
		ServerName string `csv:"server_name" size:"32"`
		ClientName string `csv:"client_name" size:"32"`
		Level      uint32 `csv:"level"`
		Money      uint32 `csv:"money"`
		Gate       uint32 `csv:"gate"`
	}

	return &bmdStruct{}
}

func (m *MoveReq) GetVersion(data *[]byte) (res int16, err error) {
	return -1, nil
}

func (m *MoveReq) GetRowsCount(data *[]byte) (res int32, err error) {
	return bmd.GetRowsCount(data)
}

func (m *MoveReq) GetCrcValue(data *[]byte) (res int32, err error) {
	return -1, nil
}

func (m *MoveReq) SetVersion(file *os.File) error {
	return nil
}

func (m *MoveReq) SetRowsCount(file *os.File, count int) error {
	return bmd.SetRowsCount(file, count)
}

func (m *MoveReq) SetCrcValue(file *os.File, buff bytes.Buffer) error {
	return nil
}
