package s9

import (
	"bytes"
	"os"

	"github.com/baranovskis/bmd-converter/bmd"
)

type EvolutionMonsterBox struct {
	bmd.BaseBmdAdapter // inherits methods
}

// Unfinished!
func (em *EvolutionMonsterBox) GetStruct() interface{} {
	type bmdStruct struct {
		Unk1  uint8 `csv:"unk_1"`
		Unk2  uint8 `csv:"unk_2"`
		Unk3  uint8 `csv:"unk_3"`
		Unk4  uint8 `csv:"unk_4"`
		Unk5  uint8 `csv:"unk_5"`
		Unk6  uint8 `csv:"unk_6"`
		Unk7  uint8 `csv:"unk_7"`
		Unk8  uint8 `csv:"unk_8"`
		Unk9  uint8 `csv:"unk_9"`
		Unk10 uint8 `csv:"unk_10"`
		Unk11 uint8 `csv:"unk_11"`
		Unk12 uint8 `csv:"unk_12"`
	}

	return &bmdStruct{}
}

func (em *EvolutionMonsterBox) GetVersion(data *[]byte) (res int16, err error) {
	return -1, nil
}

func (em *EvolutionMonsterBox) GetRowsCount(data *[]byte) (res int32, err error) {
	return bmd.GetRowsCount(data)
}

func (em *EvolutionMonsterBox) GetCrcValue(data *[]byte) (res int32, err error) {
	return bmd.GetCrcValue(data)
}

func (em *EvolutionMonsterBox) SetVersion(file *os.File) error {
	return nil
}

func (em *EvolutionMonsterBox) SetRowsCount(file *os.File, count int) error {
	return bmd.SetRowsCount(file, count)
}

func (em *EvolutionMonsterBox) SetCrcValue(file *os.File, buff bytes.Buffer) error {
	return bmd.SetCrcValue(file, buff, 0xE2F1)
}
