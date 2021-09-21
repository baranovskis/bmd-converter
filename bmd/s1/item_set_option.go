package s1

import (
	"bytes"
	"os"

	"github.com/baranovskis/bmd-converter/bmd"
)

type ItemSetOption struct {
	bmd.BaseBmdAdapter // inherits methods
}

func (is *ItemSetOption) GetStruct() interface{} {
	type bmdStruct struct {
		Name           string `csv:"name" size:"32"`
		OptIdx11       int8   `csv:"opt_idx_1_1"`
		OptIdx21       int8   `csv:"opt_idx_2_1"`
		OptIdx12       int8   `csv:"opt_idx_1_2"`
		OptIdx22       int8   `csv:"opt_idx_2_2"`
		OptIdx13       int8   `csv:"opt_idx_1_3"`
		OptIdx23       int8   `csv:"opt_idx_2_3"`
		OptIdx14       int8   `csv:"opt_idx_1_4"`
		OptIdx24       int8   `csv:"opt_idx_2_4"`
		OptIdx15       int8   `csv:"opt_idx_1_5"`
		OptIdx25       int8   `csv:"opt_idx_2_5"`
		OptIdx16       int8   `csv:"opt_idx_1_6"`
		OptIdx26       int8   `csv:"opt_idx_2_6"`
		OptVal11       int8   `csv:"opt_val_1_1"`
		OptVal21       int8   `csv:"opt_val_2_1"`
		OptVal12       int8   `csv:"opt_val_1_2"`
		OptVal22       int8   `csv:"opt_val_2_2"`
		OptVal13       int8   `csv:"opt_val_1_3"`
		OptVal23       int8   `csv:"opt_val_2_3"`
		OptVal14       int8   `csv:"opt_val_1_4"`
		OptVal24       int8   `csv:"opt_val_2_4"`
		OptVal15       int8   `csv:"opt_val_1_5"`
		OptVal25       int8   `csv:"opt_val_2_5"`
		OptVal16       int8   `csv:"opt_val_1_6"`
		OptVal26       int8   `csv:"opt_val_2_6"`
		SpecialOptIdx1 int8   `csv:"special_opt_idx_1"`
		SpecialOptIdx2 int8   `csv:"special_opt_idx_2"`
		SpecialOptVal1 int8   `csv:"special_opt_val_1"`
		SpecialOptVal2 int8   `csv:"special_opt_val_2"`
		Unk            int8   `csv:"unk_1"`
		FullOptIdx1    int8   `csv:"full_opt_idx_1"`
		FullOptIdx2    int8   `csv:"full_opt_idx_2"`
		FullOptIdx3    int8   `csv:"full_opt_idx_3"`
		FullOptIdx4    int8   `csv:"full_opt_idx_4"`
		FullOptIdx5    int8   `csv:"full_opt_idx_5"`
		FullOptVal1    int8   `csv:"full_opt_val_1"`
		FullOptVal2    int8   `csv:"full_opt_val_2"`
		FullOptVal3    int8   `csv:"full_opt_val_3"`
		FullOptVal4    int8   `csv:"full_opt_val_4"`
		FullOptVal5    int8   `csv:"full_opt_val_5"`
		ClassDW        int8   `csv:"class_dw"`
		ClassDK        int8   `csv:"class_dk"`
		ClassELF       int8   `csv:"class_elf"`
		ClassMG        int8   `csv:"class_mg"`
		ClassDL        int8   `csv:"class_dl"`
	}

	return &bmdStruct{}
}

func (is *ItemSetOption) GetVersion(data *[]byte) (res int16, err error) {
	return -1, nil
}

func (is *ItemSetOption) GetRowsCount(data *[]byte) (res int32, err error) {
	return -1, nil
}

func (is *ItemSetOption) GetCrcValue(data *[]byte) (res int32, err error) {
	return bmd.GetCrcValue(data)
}

func (is *ItemSetOption) SetVersion(file *os.File) error {
	return nil
}

func (is *ItemSetOption) SetRowsCount(file *os.File, count int) error {
	return nil
}

func (is *ItemSetOption) SetCrcValue(file *os.File, buff bytes.Buffer) error {
	return bmd.SetCrcValue(file, buff, 0xA2F1)
}
