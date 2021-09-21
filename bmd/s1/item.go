package s1

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
		Name            string `csv:"name" size:"30"`
		TwoHands        uint8  `csv:"two_hands"`
		ItemLevel       uint8  `csv:"item_level"`
		ItemSlot        uint8  `csv:"item_slot"`
		ItemSkill       uint8  `csv:"item_skill"`
		X               uint8  `csv:"x"`
		Y               uint8  `csv:"y"`
		DamageMin       uint8  `csv:"damage_min"`
		DamageMax       uint8  `csv:"damage_max"`
		DefenceRate     uint8  `csv:"defence_rate"`
		Defence         uint8  `csv:"defence"`
		Unk1            uint8  `csv:"unk_1"`
		AttackSpeed     uint8  `csv:"attack_speed"`
		WalkSpeed       uint8  `csv:"walk_speed"`
		Durability      uint8  `csv:"durability"`
		MagicDurability uint8  `csv:"magic_durability"`
		MagicPower      uint8  `csv:"magic_power"`
		ReqStr          uint16 `csv:"req_str"`
		ReqDex          uint16 `csv:"req_dex"`
		ReqEne          uint16 `csv:"req_ene"`
		ReqVit          uint16 `csv:"req_vit"`
		ReqLea          uint16 `csv:"req_lea"`
		ReqLevel        uint16 `csv:"req_level"`
		Unk2            uint16 `csv:"unk_2"`
		Value           uint16 `csv:"value"`
		Money           uint16 `csv:"money"`
		SetAttrib       uint8  `csv:"set_attrib"`
		ClassDW         uint8  `csv:"class_dw"`
		ClassDK         uint8  `csv:"class_dk"`
		ClassELF        uint8  `csv:"class_elf"`
		ClassMG         uint8  `csv:"class_mg"`
		ClassDL         uint8  `csv:"class_dl"`
		IceRes          uint8  `csv:"ice_res"`
		PoisonRes       uint8  `csv:"poison_res"`
		LightRes        uint8  `csv:"light_res"`
		FireRes         uint8  `csv:"fire_res"`
		EarthRes        uint8  `csv:"earth_res"`
		WindRes         uint8  `csv:"wind_res"`
		WaterRes        uint8  `csv:"water_res"`
		Unk3            uint8  `csv:"unk_3"`
		Unk4            uint8  `csv:"unk_4"`
		Unk5            uint8  `csv:"unk_5"`
	}

	return &bmdStruct{}
}

func (i *Item) GetVersion(data *[]byte) (res int16, err error) {
	return -1, nil
}

func (i *Item) GetRowsCount(data *[]byte) (res int32, err error) {
	return -1, nil
}

func (i *Item) GetCrcValue(data *[]byte) (res int32, err error) {
	return bmd.GetCrcValue(data)
}

func (i *Item) SetVersion(file *os.File) error {
	return nil
}

func (i *Item) SetRowsCount(file *os.File, count int) error {
	return nil
}

func (i *Item) SetCrcValue(file *os.File, buff bytes.Buffer) error {
	return bmd.SetCrcValue(file, buff, 0xE2F1)
}
