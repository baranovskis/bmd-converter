package s9

import (
	"bytes"
	"os"

	"github.com/baranovskis/bmd-converter/bmd"
)

type ItemLevelTooltip struct {
	bmd.BaseBmdAdapter // inherits methods
}

func (il *ItemLevelTooltip) GetStruct() interface{} {
	type bmdStruct struct {
		ID    uint16 `csv:"id"`
		Text  string `csv:"text" size:"64"` // set fixed string size
		Unk1  int16  `csv:"unk_1"`
		Unk2  int16  `csv:"unk_2"`
		Unk3  int16  `csv:"unk_3"`
		Unk4  int16  `csv:"unk_4"`
		Unk5  int16  `csv:"unk_5"`
		Unk6  int16  `csv:"unk_6"`
		Unk7  int16  `csv:"unk_7"`
		Unk8  int16  `csv:"unk_8"`
		Unk9  int16  `csv:"unk_9"`
		Unk10 int16  `csv:"unk_10"`
		Unk11 int16  `csv:"unk_11"`
		Unk12 int16  `csv:"unk_12"`
		Unk13 int16  `csv:"unk_13"`
		Unk14 int16  `csv:"unk_14"`
		Unk15 int16  `csv:"unk_15"`
		Unk16 int16  `csv:"unk_16"`
		Unk17 int16  `csv:"unk_17"`
		Unk18 int16  `csv:"unk_18"`
		Unk19 int16  `csv:"unk_19"`
		Unk20 int16  `csv:"unk_20"`
		Unk21 int16  `csv:"unk_21"`
		Unk22 int16  `csv:"unk_22"`
		Unk23 int16  `csv:"unk_23"`
		Unk24 int16  `csv:"unk_24"`
	}

	return &bmdStruct{}
}

func (il *ItemLevelTooltip) GetVersion(data *[]byte) (res int16, err error) {
	return -1, nil
}

func (il *ItemLevelTooltip) GetRowsCount(data *[]byte) (res int32, err error) {
	return -1, nil
}

func (il *ItemLevelTooltip) GetCrcValue(data *[]byte) (res int32, err error) {
	return bmd.GetCrcValue(data)
}

func (il *ItemLevelTooltip) SetVersion(file *os.File) error {
	return nil
}

func (il *ItemLevelTooltip) SetRowsCount(file *os.File, count int) error {
	return nil
}

func (il *ItemLevelTooltip) SetCrcValue(file *os.File, buff bytes.Buffer) error {
	return bmd.SetCrcValue(file, buff, 0xE2F1)
}
