package bmd

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/gocarina/gocsv"
)

var (
	XorKey = []byte{0xFC, 0xCF, 0xAB}
)

type BaseBmdAdapter interface {
	GetStruct() interface{}

	GetVersion(data *[]byte) (res int16, err error)
	GetRowsCount(data *[]byte) (res int32, err error)
	GetCrcValue(data *[]byte) (res int32, err error)

	SetVersion(f *os.File) error
	SetRowsCount(f *os.File, count int) error
	SetCrcValue(f *os.File, b bytes.Buffer) error
}

type BaseBmdInfo struct {
	adapter BaseBmdAdapter
}

func NewBmdInfo(b BaseBmdAdapter) *BaseBmdInfo {
	return &BaseBmdInfo{adapter: b}
}

func (b BaseBmdInfo) Decrypt(data []byte, out *os.File) error {
	dataType := reflect.TypeOf(b.adapter.GetStruct())
	dataArray := reflect.MakeSlice(reflect.SliceOf(dataType), 0, 0)

	if version, err := b.adapter.GetVersion(&data); version != -1 {
		if err != nil {
			return err
		}
		fmt.Println("bmd version:", uint32(version))
	}

	if count, err := b.adapter.GetRowsCount(&data); count != -1 {
		if err != nil {
			return err
		}
		fmt.Println("rows count:", uint32(count))
	}

	if crc, err := b.adapter.GetCrcValue(&data); crc != -1 {
		if err != nil {
			return err
		}
		fmt.Println("file crc:", uint32(crc))
	}

	reader := bytes.NewReader(data)

	for reader.Len() > 0 {
		data := b.adapter.GetStruct()
		val := reflect.ValueOf(data).Elem()

		totalLen := 0

		for i := 0; i < val.NumField(); i++ {
			valueField := val.Field(i)
			valueType := val.Type().Field(i)

			len, err := getFieldSize(val, i)
			if err != nil {
				return err
			}

			tmp := make([]byte, len)

			err = binary.Read(reader, binary.LittleEndian, &tmp)
			if err != nil {
				return err
			}

			// Do slice XOR only when it was enabled
			if tagValue := valueType.Tag.Get("xor"); tagValue != "false" {
				doSliceXor(tmp, totalLen)
				totalLen += len
			}

			buff := bytes.NewReader(tmp)

			switch valueField.Interface().(type) {
			// Int
			case int:
				var val int
				err = binary.Read(buff, binary.LittleEndian, &val)
				valueField.SetInt(int64(val))
			case int8:
				var val int8
				err = binary.Read(buff, binary.LittleEndian, &val)
				valueField.SetInt(int64(val))
			case int16:
				var val int16
				err = binary.Read(buff, binary.LittleEndian, &val)
				valueField.SetInt(int64(val))
			case int32:
				var val int32
				err = binary.Read(buff, binary.LittleEndian, &val)
				valueField.SetInt(int64(val))
			case int64:
				var val int64
				err = binary.Read(buff, binary.LittleEndian, &val)
				valueField.SetInt(val)

			// Uint
			case uint:
				var val uint
				err = binary.Read(buff, binary.LittleEndian, &val)
				valueField.SetUint(uint64(val))
			case uint8:
				var val uint8
				err = binary.Read(buff, binary.LittleEndian, &val)
				valueField.SetUint(uint64(val))
			case uint16:
				var val uint16
				err = binary.Read(buff, binary.LittleEndian, &val)
				valueField.SetUint(uint64(val))
			case uint32:
				var val uint32
				err = binary.Read(buff, binary.LittleEndian, &val)
				valueField.SetUint(uint64(val))
			case uint64:
				var val uint64
				err = binary.Read(buff, binary.LittleEndian, &val)
				valueField.SetUint(val)

			// String
			case string:
				if strLen := clen(tmp); strLen > 0 {
					aaa := string(tmp[:strLen])
					valueField.SetString(aaa)
				} else {
					valueField.SetString(string('\x00'))
				}

			default:
				return errors.New("unknown structure type")
			}

			if err != nil {
				return err
			}
		}

		ptr := reflect.New(dataType)
		ptr.Elem().Set(reflect.ValueOf(data))
		value := reflect.ValueOf(ptr.Interface()).Elem()

		dataArray = reflect.Append(dataArray, value)
	}

	err := gocsv.MarshalFile(dataArray.Interface(), out)
	if err != nil {
		return err
	}

	fmt.Println("decrypt completed!")
	return nil
}

func (b BaseBmdInfo) Encrypt(data []byte, out *os.File) error {
	bmdStruct := b.adapter.GetStruct()
	sliceType := reflect.SliceOf(reflect.TypeOf(bmdStruct))
	slicePtr := reflect.New(sliceType)

	if err := gocsv.UnmarshalBytes(data, slicePtr.Interface()); err != nil {
		return err
	}

	sliceData := reflect.ValueOf(slicePtr.Interface()).Elem()
	rowsCount := sliceData.Len()

	err := b.adapter.SetVersion(out)
	if err != nil {
		return err
	}

	err = b.adapter.SetRowsCount(out, rowsCount)
	if err != nil {
		return err
	}

	var buff bytes.Buffer

	for i := 0; i < rowsCount; i++ {
		v := reflect.ValueOf(sliceData.Index(i).Interface()).Elem()

		// Find and update string length
		for e := 0; e < v.NumField(); e++ {
			currField := v.Field(e)
			currType := v.Type().Field(e)

			// Only for "auto" size mode
			if tagValue := currType.Tag.Get("size"); tagValue == "auto" {
				switch currField.Interface().(type) {
				case string:
					prevField := v.Field(e - 1)
					prevField.SetUint(uint64(len(currField.String()) + 1))
				}
			}
		}

		totalLen := 0

		for e := 0; e < v.NumField(); e++ {
			varValue := v.Field(e)
			varType := v.Type().Field(e)

			len, err := getFieldSize(v, e)
			if err != nil {
				return err
			}

			tmp := make([]byte, len)

			switch v := varValue.Interface().(type) {
			// Int
			case int8:
				tmp[0] = byte(v)
			case int16:
				binary.LittleEndian.PutUint16(tmp, uint16(v))
			case int32:
				binary.LittleEndian.PutUint32(tmp, uint32(v))
			case int64:
				binary.LittleEndian.PutUint64(tmp, uint64(v))

			// Uint
			case uint8:
				tmp[0] = v
			case uint16:
				binary.LittleEndian.PutUint16(tmp, v)
			case uint32:
				binary.LittleEndian.PutUint32(tmp, v)
			case uint64:
				binary.LittleEndian.PutUint64(tmp, v)

			// String
			case string:
				copy(tmp, varValue.Interface().(string))

			default:
				return errors.New("unknown structure type")
			}

			// Do slice XOR only when it was enabled
			if tagValue := varType.Tag.Get("xor"); tagValue != "false" {
				doSliceXor(tmp, totalLen)
				totalLen += len
			}

			err = binary.Write(out, binary.LittleEndian, tmp)
			if err != nil {
				return err
			}

			buff.Write(tmp)
		}
	}

	// Set crc32 value
	err = b.adapter.SetCrcValue(out, buff)
	if err != nil {
		return err
	}

	fmt.Println("encrypt completed!")
	return nil
}

func getFieldSize(val reflect.Value, pos int) (size int, err error) {
	varType := val.Type().Field(pos)

	// If type is string, then need check tag or previous field
	if varType.Type.Kind() == reflect.String {
		// There were 2 types of string size
		if tagValue := varType.Tag.Get("size"); tagValue != "auto" {
			// read from tag "size" fixed string size
			size, err = strconv.Atoi(tagValue)
		} else {
			// get from previous field string size
			size = int(val.Field(pos - 1).Uint())
		}
	} else {
		size = int(varType.Type.Size())
	}

	return
}

func doSliceXor(data []byte, totalLen int) {
	for i := 0; i < len(data); i++ {
		data[i] ^= XorKey[(i+totalLen)%len(XorKey)]
	}
}

func GetVersion(data *[]byte) (res int16, err error) {
	reader := bytes.NewReader((*data)[0:2]) // read first 2 bytes
	err = binary.Read(reader, binary.LittleEndian, &res)
	*data = (*data)[2:] // exclude first 2 bytes
	return
}

func GetRowsCount(data *[]byte) (res int32, err error) {
	reader := bytes.NewReader((*data)[0:4]) // read first 4 bytes
	err = binary.Read(reader, binary.LittleEndian, &res)
	*data = (*data)[4:] // exclude first 4 bytes
	return
}

func GetCrcValue(data *[]byte) (res int32, err error) {
	reader := bytes.NewReader((*data)[len(*data)-4:]) // read last 4 bytes
	err = binary.Read(reader, binary.LittleEndian, &res)
	*data = (*data)[:len(*data)-4] // exclude last 4 bytes
	return
}

func SetVersion(file *os.File, version int) error {
	fmt.Println("set version =", version)

	tmp := make([]byte, 2)
	binary.LittleEndian.PutUint16(tmp, uint16(version))
	return binary.Write(file, binary.LittleEndian, tmp)
}

func SetRowsCount(file *os.File, count int) error {
	fmt.Println("set rows count =", count)

	tmp := make([]byte, 4)
	binary.LittleEndian.PutUint32(tmp, uint32(count))
	return binary.Write(file, binary.LittleEndian, tmp)
}

func SetCrcValue(file *os.File, buff bytes.Buffer, keyVal uint32) error {
	reader := bytes.NewReader(buff.Bytes())
	crcVal := keyVal * 512

	varTemp := []uint32{
		0xFFFFFFFF,
		0x7FFFFFFF,
		0x3FFFFFFF,
		0x1FFFFFFF,
		0x0FFFFFFF,
		0x07FFFFFF,
		0x03FFFFFF,
		0x01FFFFFF,
		0x00FFFFFF,
	}

	var pos uint32

	for reader.Len() > 0 {
		var val uint32
		err := binary.Read(reader, binary.LittleEndian, &val)
		if err != nil {
			return err
		}

		if ((pos/4)+keyVal)%2 == 1 {
			crcVal += val
		} else {
			crcVal ^= val
		}

		if pos%16 == 0 {
			crcVal ^= (keyVal + crcVal) >> (((pos / 4) % 8) + 1) & varTemp[((pos/4)%8)+1]
		}

		pos += 4
	}

	fmt.Println("calculated crc =", crcVal)

	tmp := make([]byte, 4)
	binary.LittleEndian.PutUint32(tmp, crcVal)
	return binary.Write(file, binary.LittleEndian, tmp)
}

func clen(n []byte) int {
	for i := 0; i < len(n); i++ {
		if n[i] == 0 {
			return i
		}
	}
	return len(n)
}
