package s9

import (
	"bytes"
	"os"

	"github.com/baranovskis/bmd-converter/bmd"
)

type QuestProgress struct {
	bmd.BaseBmdAdapter // inherits methods
}

func (qp *QuestProgress) GetStruct() interface{} {
	type bmdStruct struct {
		ID1  uint16 `csv:"id_1"`
		ID2  uint16 `csv:"id_2"`
		ID3  byte   `csv:"id_3"`
		ID4  uint16 `csv:"id_4"`
		ID5  uint16 `csv:"id_5"`
		ID6  uint16 `csv:"id_6"`
		ID7  uint16 `csv:"id_7"`
		ID8  uint16 `csv:"id_8"`
		ID9  uint16 `csv:"id_9"`
		ID10 uint16 `csv:"id_10"`
		ID11 uint16 `csv:"id_11"`
		ID12 uint16 `csv:"id_12"`
		ID13 uint16 `csv:"id_13"`
		ID14 uint16 `csv:"id_14"`
		ID15 uint16 `csv:"id_15"`
		ID16 uint16 `csv:"id_16"`
		ID17 uint16 `csv:"id_17"`
		ID18 uint16 `csv:"id_18"`
		ID19 uint16 `csv:"id_19"`
		ID20 uint16 `csv:"id_20"`
		ID21 uint16 `csv:"id_21"`
	}

	return &bmdStruct{}
}

func (qp *QuestProgress) GetVersion(data *[]byte) (res int16, err error) {
	return -1, nil
}

func (qp *QuestProgress) GetRowsCount(data *[]byte) (res int32, err error) {
	return -1, nil
}

func (qp *QuestProgress) GetCrcValue(data *[]byte) (res int32, err error) {
	return -1, nil
}

func (qp *QuestProgress) SetVersion(file *os.File) error {
	return nil
}

func (qp *QuestProgress) SetRowsCount(file *os.File, count int) error {
	return nil
}

func (qp *QuestProgress) SetCrcValue(file *os.File, buff bytes.Buffer) error {
	return nil
}
