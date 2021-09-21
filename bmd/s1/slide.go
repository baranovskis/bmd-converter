package s1

import (
	"bytes"
	"os"

	"github.com/baranovskis/bmd-converter/bmd"
)

type Slide struct {
	bmd.BaseBmdAdapter // inherits methods
}

func (s *Slide) GetStruct() interface{} {
	type bmdStruct struct {
		Delay      uint32 `csv:"delay"`
		SlideCount uint32 `csv:"slide_count"`
		Text       string `csv:"text_1" size:"256"`
		Text2      string `csv:"text_2" size:"256"`
		Text3      string `csv:"text_3" size:"256"`
		Text4      string `csv:"text_4" size:"256"`
		Text5      string `csv:"text_5" size:"256"`
		Text6      string `csv:"text_6" size:"256"`
		Text7      string `csv:"text_7" size:"256"`
		Text8      string `csv:"text_8" size:"256"`
		Text9      string `csv:"text_9" size:"256"`
		Text10     string `csv:"text_10" size:"256"`
		Text11     string `csv:"text_11" size:"256"`
		Text12     string `csv:"text_12" size:"256"`
		Text13     string `csv:"text_13" size:"256"`
		Text14     string `csv:"text_14" size:"256"`
		Text15     string `csv:"text_15" size:"256"`
		Text16     string `csv:"text_16" size:"256"`
		Text17     string `csv:"text_17" size:"256"`
		Text18     string `csv:"text_18" size:"256"`
		Text19     string `csv:"text_19" size:"256"`
		Text20     string `csv:"text_20" size:"256"`
		Text21     string `csv:"text_21" size:"256"`
		Text22     string `csv:"text_22" size:"256"`
		Text23     string `csv:"text_23" size:"256"`
		Text24     string `csv:"text_24" size:"256"`
		Text25     string `csv:"text_25" size:"256"`
		Text26     string `csv:"text_26" size:"256"`
		Text27     string `csv:"text_27" size:"256"`
		Text28     string `csv:"text_28" size:"256"`
		Text29     string `csv:"text_29" size:"256"`
		Text30     string `csv:"text_30" size:"256"`
		Text31     string `csv:"text_31" size:"256"`
		Text32     string `csv:"text_32" size:"256"`
	}

	return &bmdStruct{}
}

func (s *Slide) GetVersion(data *[]byte) (res int16, err error) {
	return -1, nil
}

func (s *Slide) GetRowsCount(data *[]byte) (res int32, err error) {
	return -1, nil
}

func (s *Slide) GetCrcValue(data *[]byte) (res int32, err error) {
	return -1, nil
}

func (s *Slide) SetVersion(file *os.File) error {
	return nil
}

func (s *Slide) SetRowsCount(file *os.File, count int) error {
	return nil
}

func (s *Slide) SetCrcValue(file *os.File, buff bytes.Buffer) error {
	return nil
}
