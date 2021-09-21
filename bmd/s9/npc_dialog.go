package s9

import (
	"bytes"
	"os"

	"github.com/baranovskis/bmd-converter/bmd"
)

type NpcDialog struct {
	bmd.BaseBmdAdapter // inherits methods
}

func (nd *NpcDialog) GetStruct() interface{} {
	type bmdStruct struct {
		GroupID     uint16 `csv:"group_id"`
		NpcID       uint16 `csv:"npc_id"`
		TextID      uint32 `csv:"text_id"`
		BtnTextID1  uint32 `csv:"bnt_text_id_1"`
		BtnID1      uint32 `csv:"btn_id_1"`
		BtnTextID2  uint32 `csv:"bnt_text_id_2"`
		BtnID2      uint32 `csv:"btn_id_2"`
		BtnTextID3  uint32 `csv:"bnt_text_id_3"`
		BtnID3      uint32 `csv:"btn_id_3"`
		BtnTextID4  uint32 `csv:"bnt_text_id_4"`
		BtnID4      uint32 `csv:"btn_id_4"`
		BtnTextID5  uint32 `csv:"bnt_text_id_5"`
		BtnID5      uint32 `csv:"btn_id_5"`
		BtnTextID6  uint32 `csv:"bnt_text_id_6"`
		BtnID6      uint32 `csv:"btn_id_6"`
		BtnTextID7  uint32 `csv:"bnt_text_id_7"`
		BtnID7      uint32 `csv:"btn_id_7"`
		BtnTextID8  uint32 `csv:"bnt_text_id_8"`
		BtnID8      uint32 `csv:"btn_id_8"`
		BtnTextID9  uint32 `csv:"bnt_text_id_9"`
		BtnID9      uint32 `csv:"btn_id_9"`
		BtnTextID10 uint32 `csv:"bnt_text_id_10"`
		BtnID10     uint32 `csv:"btn_id_10"`
	}

	return &bmdStruct{}
}

func (nd *NpcDialog) GetVersion(data *[]byte) (res int16, err error) {
	return -1, nil
}

func (nd *NpcDialog) GetRowsCount(data *[]byte) (res int32, err error) {
	return -1, nil
}

func (nd *NpcDialog) GetCrcValue(data *[]byte) (res int32, err error) {
	return -1, nil
}

func (nd *NpcDialog) SetVersion(file *os.File) error {
	return nil
}

func (nd *NpcDialog) SetRowsCount(file *os.File, count int) error {
	return nil
}

func (nd *NpcDialog) SetCrcValue(file *os.File, buff bytes.Buffer) error {
	return nil
}
