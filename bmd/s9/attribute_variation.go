package s9

import (
	"bytes"
	"os"

	"github.com/baranovskis/bmd-converter/bmd"
)

type AttributeVariation struct {
	bmd.BaseBmdAdapter // inherits methods
}

func (av *AttributeVariation) GetStruct() interface{} {
	type bmdStruct struct {
		Type      uint32 `csv:"type"`
		ID        uint32 `csv:"id"`
		ItemType  uint32 `csv:"item_type"`
		ItemIndex uint32 `csv:"item_index"`
		Unk1      uint32 `csv:"unk_1"`
		Unk2      uint32 `csv:"unk_2"`
		Unk3      uint32 `csv:"unk_3"`
		Unk4      uint32 `csv:"unk_4"`
		Unk5      uint32 `csv:"unk_5"`
		Unk6      uint32 `csv:"unk_6"`
		Unk7      uint32 `csv:"unk_7"`
		Folder1   string `csv:"folder_1" size:"256"`
		Unk8      uint32 `csv:"unk_8"`
		Folder2   string `csv:"folder_2" size:"256"`
		Unk9      uint32 `csv:"unk_9"`
		Folder3   string `csv:"folder_3" size:"256"`
		Unk10     uint32 `csv:"unk_10"`
		Folder4   string `csv:"folder_4" size:"256"`
		Unk11     uint32 `csv:"unk_11"`
		Folder5   string `csv:"folder_5" size:"256"`
		Unk12     uint32 `csv:"unk_12"`
		Folder6   string `csv:"folder_6" size:"256"`
		Unk13     uint32 `csv:"unk_13"`
		Texture1  string `csv:"texture_1" size:"256"`
		Texture2  string `csv:"texture_2" size:"256"`
		Texture3  string `csv:"texture_3" size:"256"`
		Texture4  string `csv:"texture_4" size:"256"`
		Texture5  string `csv:"texture_5" size:"256"`
		Texture6  string `csv:"texture_6" size:"256"`
	}

	return &bmdStruct{}
}

func (av *AttributeVariation) GetVersion(data *[]byte) (res int16, err error) {
	return -1, nil
}

func (av *AttributeVariation) GetRowsCount(data *[]byte) (res int32, err error) {
	return bmd.GetRowsCount(data)
}

func (av *AttributeVariation) GetCrcValue(data *[]byte) (res int32, err error) {
	return bmd.GetCrcValue(data)
}

func (av *AttributeVariation) SetVersion(file *os.File) error {
	return nil
}

func (av *AttributeVariation) SetRowsCount(file *os.File, count int) error {
	return bmd.SetRowsCount(file, count)
}

func (av *AttributeVariation) SetCrcValue(file *os.File, buff bytes.Buffer) error {
	return bmd.SetCrcValue(file, buff, 0xE2F1)
}
