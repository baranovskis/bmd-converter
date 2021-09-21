package s9

import (
	"bytes"
	"os"

	"github.com/baranovskis/bmd-converter/bmd"
)

type QuestWords struct {
	bmd.BaseBmdAdapter // inherits methods
}

func (qw *QuestWords) GetStruct() interface{} {
	type bmdStruct struct {
		ID   int32  `csv:"id"`
		Size uint16 `csv:"-"`
		Text string `csv:"text" size:"auto"`
	}

	return &bmdStruct{}
}

func (qw *QuestWords) GetVersion(data *[]byte) (res int16, err error) {
	return -1, nil
}

func (qw *QuestWords) GetRowsCount(data *[]byte) (res int32, err error) {
	return -1, nil
}

func (qw *QuestWords) GetCrcValue(data *[]byte) (res int32, err error) {
	return -1, nil
}

func (qw *QuestWords) SetVersion(file *os.File) error {
	return nil
}

func (qw *QuestWords) SetRowsCount(file *os.File, count int) error {
	return nil
}

func (qw *QuestWords) SetCrcValue(file *os.File, buff bytes.Buffer) error {
	return nil
}
