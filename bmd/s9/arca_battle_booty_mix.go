package s9

import (
	"bytes"
	"os"

	"github.com/baranovskis/bmd-converter/bmd"
)

type ArcaBattleBootyMix struct {
	bmd.BaseBmdAdapter // inherits methods
}

func (ab *ArcaBattleBootyMix) GetStruct() interface{} {
	type bmdStruct struct {
		Unk1 uint32 `csv:"unk_1"`
		Unk2 uint32 `csv:"unk_2"`
		Unk3 uint32 `csv:"unk_3"`
		Unk4 uint32 `csv:"unk_4"`
		Unk5 uint32 `csv:"unk_5"`
	}

	return &bmdStruct{}
}

func (ab *ArcaBattleBootyMix) GetVersion(data *[]byte) (res int16, err error) {
	return -1, nil
}

func (ab *ArcaBattleBootyMix) GetRowsCount(data *[]byte) (res int32, err error) {
	return bmd.GetRowsCount(data)
}

func (ab *ArcaBattleBootyMix) GetCrcValue(data *[]byte) (res int32, err error) {
	return bmd.GetCrcValue(data)
}

func (ab *ArcaBattleBootyMix) SetVersion(file *os.File) error {
	return nil
}

func (ab *ArcaBattleBootyMix) SetRowsCount(file *os.File, count int) error {
	return bmd.SetRowsCount(file, count)
}

func (ab *ArcaBattleBootyMix) SetCrcValue(file *os.File, buff bytes.Buffer) error {
	return bmd.SetCrcValue(file, buff, 0xE2F1)
}
