package s9

import (
	"bytes"
	"os"

	"github.com/baranovskis/bmd-converter/bmd"
)

type Item struct {
	bmd.BaseBmdAdapter // inherits methods
}

func (i *Item) GetStruct() interface{} {
	type bmdStruct struct {
		ID              uint32 `csv:"id"`
		ItemType        uint16 `csv:"item_type"`
		ItemIndex       uint16 `csv:"item_index"`
		TexturePath     string `csv:"texture_path" size:"260"`
		ModelName       string `csv:"model_name" size:"260"`
		Name            string `csv:"name" size:"30"`
		Type1           uint8  `csv:"type_1"`
		Type2           uint8  `csv:"type_2"`
		Type3           uint8  `csv:"type_3"`
		TwoHands        uint8  `csv:"two_hands"`
		ItemLevel       uint16 `csv:"item_level"`
		ItemSlot        uint16 `csv:"item_slot"`
		ItemSkill       uint16 `csv:"item_skill"`
		X               uint8  `csv:"x"`
		Y               uint8  `csv:"y"`
		DamageMin       uint16 `csv:"damage_min"`
		DamageMax       uint16 `csv:"damage_max"`
		DefenceRate     uint16 `csv:"defence_rate"`
		Defence         uint16 `csv:"defence"`
		Unk1            uint16 `csv:"unk_1"` // Check me
		AttackSpeed     uint8  `csv:"attack_speed"`
		WalkSpeed       uint8  `csv:"walk_speed"`
		Durability      uint8  `csv:"durability"`
		MagicDurability uint8  `csv:"magic_durability"`
		MagicPower      uint8  `csv:"magic_power"`
		Unk2            uint8  `csv:"unk_2"` // Check me
		ReqStr          uint16 `csv:"req_str"`
		ReqDex          uint16 `csv:"req_dex"`
		ReqEne          uint16 `csv:"req_ene"`
		ReqVit          uint16 `csv:"req_vit"`
		ReqLea          uint16 `csv:"req_lea"`
		ReqLevel        uint16 `csv:"req_level"`
		ItemValue       uint16 `csv:"item_value"`
		ItemPrice       uint32 `csv:"item_price"`
		SetAttrib       uint8  `csv:"set_attrib"`
		ClassDW         uint8  `csv:"class_dw"`
		ClassDK         uint8  `csv:"class_dk"`
		ClassELF        uint8  `csv:"class_elf"`
		ClassMG         uint8  `csv:"class_mg"`
		ClassDL         uint8  `csv:"class_dl"`
		ClassSUM        uint8  `csv:"class_sum"`
		ClassRF         uint8  `csv:"class_rf"`
		Resistance1     uint8  `csv:"resistance_1"`
		Resistance2     uint8  `csv:"resistance_2"`
		Resistance3     uint8  `csv:"resistance_3"`
		Resistance4     uint8  `csv:"resistance_4"`
		Resistance5     uint8  `csv:"resistance_5"`
		Resistance6     uint8  `csv:"resistance_6"`
		Resistance7     uint8  `csv:"resistance_7"`
		IsApplyToDrop   uint8  `csv:"allow_drop"`
		Unk3            uint8  `csv:"unk_3"` // Check me
		Unk4            uint8  `csv:"unk_4"` // Check me
		Unk5            uint8  `csv:"unk_5"` // Check me
		Unk6            uint8  `csv:"unk_6"` // Check me
		IsExpensive     uint8  `csv:"is_expensive"`
		Unk7            uint8  `csv:"unk_7"`     // Check me
		StackMax        uint8  `csv:"stack_max"` // Check me
		Pad             uint8  `csv:"pad"`
	}

	//+0 = Is drop apply item
	//+5 = IsExpensive (sell notice, drop block)
	//+7 = Max durability (for potions, rena and etc.)
	return &bmdStruct{}
}

func (i *Item) GetVersion(data *[]byte) (res int16, err error) {
	return -1, nil
}

func (i *Item) GetRowsCount(data *[]byte) (res int32, err error) {
	return bmd.GetRowsCount(data)
}

// GetCrcValue is return crc value
func (i *Item) GetCrcValue(data *[]byte) (res int32, err error) {
	return bmd.GetCrcValue(data)
}

func (i *Item) SetVersion(file *os.File) error {
	return nil
}

func (i *Item) SetRowsCount(file *os.File, count int) error {
	return bmd.SetRowsCount(file, count)
}

func (i *Item) SetCrcValue(file *os.File, buff bytes.Buffer) error {
	return bmd.SetCrcValue(file, buff, 0xE2F1)
}
