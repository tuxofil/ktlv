package ktlv

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
)

const (
	Bool   = 0
	Uint8  = 1
	Uint16 = 2
	Uint24 = 3
	Uint32 = 4
	Uint64 = 5
	Double = 6
	String = 7
	Bitmap = 8
	Int8   = 9
	Int16  = 10
	Int24  = 11
	Int32  = 12
	Int64  = 13

	List_of_String = 50
	List_of_Uint8  = 51
	List_of_Uint16 = 52
	List_of_Uint24 = 53
	List_of_Uint32 = 54
	List_of_Uint64 = 55
	List_of_Double = 56
	List_of_Int8   = 57
	List_of_Int16  = 58
	List_of_Int24  = 59
	List_of_Int32  = 60
	List_of_Int64  = 61

	Min_Int8   = int8(-0x80)
	Min_Int16  = int16(-0x8000)
	Min_Int24  = int32(-0x800000)
	Min_Int32  = int32(-0x80000000)
	Min_Int64  = int64(-0x8000000000000000)
	Max_Int8   = int8(0x7f)
	Max_Int16  = int16(0x7fff)
	Max_Int24  = int32(0x7fffff)
	Max_Int32  = int32(0x7fffffff)
	Max_Int64  = int64(0x7fffffffffffffff)
	Min_Uint8  = uint8(0)
	Min_Uint16 = uint16(0)
	Min_Uint24 = uint32(0)
	Min_Uint32 = uint32(0)
	Min_Uint64 = uint64(0)
	Max_Uint8  = uint8(0xff)
	Max_Uint16 = uint16(0xffff)
	Max_Uint24 = uint32(0xffffff)
	Max_Uint32 = uint32(0xffffffff)
	Max_Uint64 = uint64(0xffffffffffffffff)
)

var t2s = map[uint8]string{
	Bool:           "Bool",
	Uint8:          "Uint8",
	Uint16:         "Uint16",
	Uint24:         "Uint24",
	Uint32:         "Uint32",
	Uint64:         "Uint64",
	Double:         "Double",
	String:         "String",
	Bitmap:         "Bitmap",
	Int8:           "Int8",
	Int16:          "Int16",
	Int24:          "Int24",
	Int32:          "Int32",
	Int64:          "Int64",
	List_of_String: "List_of_String",
	List_of_Uint8:  "List_of_Uint8",
	List_of_Uint16: "List_of_Uint16",
	List_of_Uint24: "List_of_Uint24",
	List_of_Uint32: "List_of_Uint32",
	List_of_Uint64: "List_of_Uint64",
	List_of_Double: "List_of_Double",
	List_of_Int8:   "List_of_Int8",
	List_of_Int16:  "List_of_Int16",
	List_of_Int24:  "List_of_Int24",
	List_of_Int32:  "List_of_Int32",
	List_of_Int64:  "List_of_Int64",
}

func FTypeToString(t uint8) string {
	return t2s[t]
}

type Elem struct {
	Key   uint16
	FType uint8
	Value interface{}
}

type Data []*Elem

type DataDict map[uint16]*Elem

var (
	ElementNotFound     = errors.New("no such element")
	TypeAssertionFailed = errors.New("unexpected element type")
)

// Encode input data to byte buffer.
func (d Data) Encode() ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	for _, elem := range d {
		encoded, err := elem.Encode()
		if err != nil {
			return nil, err
		}
		if _, err := buffer.Write(encoded); err != nil {
			return nil, err
		}
	}
	return buffer.Bytes(), nil
}

// Encode dictionary with input data to byte buffer.
func (d DataDict) Encode() ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	for _, elem := range d {
		encoded, err := elem.Encode()
		if err != nil {
			return nil, err
		}
		if _, err := buffer.Write(encoded); err != nil {
			return nil, err
		}
	}
	return buffer.Bytes(), nil
}

func (d DataDict) Put(key uint16, ftype uint8, v interface{}) {
	d[key] = &Elem{key, ftype, v}
}

func (d DataDict) Get(key uint16) (ftype uint8, v interface{}, ok bool) {
	if elem, ok := d[key]; ok {
		return elem.FType, elem.Value, true
	}
	return 0, nil, false
}

func (d DataDict) String() string {
	s := ""
	for k, e := range d {
		if s != "" {
			s += ","
		}
		s += fmt.Sprintf("%d(%s)=%#v", k,
			FTypeToString(e.FType), e.Value)
	}
	return s
}

// Encode input data to byte buffer.
// Deprecated API.
func Enc(data Data) []byte {
	bytes, err := data.Encode()
	if err != nil {
		log.Fatalf("deprecated API: %s", err)
	}
	return bytes
}

// Decode data from byte buffer.
func Decode(bytes []byte) (Data, error) {
	res := make(Data, 0)
	for 0 < len(bytes) {
		elem, tail, err := scan(bytes)
		if err != nil {
			return nil, err
		}
		res = append(res, elem)
		bytes = tail
	}
	return res, nil
}

// Decode data from byte buffer to dictionary.
func DecodeDict(bytes []byte) (DataDict, error) {
	res := make(DataDict)
	for 0 < len(bytes) {
		elem, tail, err := scan(bytes)
		if err != nil {
			return nil, err
		}
		res[elem.Key] = elem
		bytes = tail
	}
	return res, nil
}

// String field getter.
func (d *DataDict) GetString(key uint16) (string, error) {
	if elem, ok := (*d)[key]; ok {
		if elem.FType == String {
			return elem.Value.(string), nil
		}
		return "", TypeAssertionFailed
	}
	return "", ElementNotFound
}

// String field getter.
func (d *DataDict) GetStringDef(key uint16, def string) string {
	if v, err := d.GetString(key); err == nil {
		return v
	}
	return def
}

// bool field getter.
func (d *DataDict) GetBoolDef(key uint16, def bool) bool {
	if elem, ok := (*d)[key]; ok {
		if elem.FType == Bool {
			return elem.Value.(bool)
		}
	}
	return def
}

// uint8 field getter.
func (d *DataDict) GetUint8(key uint16) (uint8, error) {
	if elem, ok := (*d)[key]; ok {
		if elem.FType == Uint8 {
			return elem.Value.(uint8), nil
		}
		return 0, TypeAssertionFailed
	}
	return 0, ElementNotFound
}

// uint8 field getter.
func (d *DataDict) GetUint8Def(key uint16, def uint8) uint8 {
	if elem, ok := (*d)[key]; ok {
		if elem.FType == Uint8 {
			return elem.Value.(uint8)
		}
	}
	return def
}

// uint16 field getter.
func (d *DataDict) GetUint16Def(key uint16, def uint16) uint16 {
	if elem, ok := (*d)[key]; ok {
		if elem.FType == Uint16 {
			return elem.Value.(uint16)
		}
	}
	return def
}

// uint32 field getter.
func (d *DataDict) GetUint32(key uint16) (uint32, error) {
	if elem, ok := (*d)[key]; ok {
		if elem.FType == Uint32 {
			return elem.Value.(uint32), nil
		}
		return 0, TypeAssertionFailed
	}
	return 0, ElementNotFound
}

// uint32 field getter.
func (d *DataDict) GetUint32Def(key uint16, def uint32) uint32 {
	if elem, ok := (*d)[key]; ok {
		if elem.FType == Uint32 {
			return elem.Value.(uint32)
		}
	}
	return def
}

// uint64 field getter.
func (d *DataDict) GetUint64(key uint16) (uint64, error) {
	if elem, ok := (*d)[key]; ok {
		if elem.FType == Uint64 {
			return elem.Value.(uint64), nil
		}
		return 0, TypeAssertionFailed
	}
	return 0, ElementNotFound
}

// uint64 field getter.
func (d *DataDict) GetUint64Def(key uint16, def uint64) uint64 {
	if v, err := d.GetUint64(key); err == nil {
		return v
	}
	return def
}

// double field getter.
func (d *DataDict) GetDouble(key uint16) (float64, error) {
	if elem, ok := (*d)[key]; ok {
		if elem.FType == Double {
			return elem.Value.(float64), nil
		}
		return 0, TypeAssertionFailed
	}
	return 0, ElementNotFound
}

// double field getter.
func (d *DataDict) GetDoubleDef(key uint16, def float64) float64 {
	if v, err := d.GetDouble(key); err == nil {
		return v
	}
	return def
}

// list of uint8 field getter.
func (d *DataDict) GetListOfUint8Def(key uint16, def []uint8) []uint8 {
	if elem, ok := (*d)[key]; ok {
		if elem.FType == List_of_Uint8 {
			return elem.Value.([]uint8)
		}
	}
	return def
}

// list of uint32 field getter.
func (d *DataDict) GetListOfUint32Def(key uint16, def []uint32) []uint32 {
	if elem, ok := (*d)[key]; ok {
		if elem.FType == List_of_Uint32 {
			return elem.Value.([]uint32)
		}
	}
	return def
}

// list of string field getter.
func (d *DataDict) GetListOfStringDef(key uint16, def []string) []string {
	if elem, ok := (*d)[key]; ok {
		if elem.FType == List_of_String {
			return elem.Value.([]string)
		}
	}
	return def
}

// list of double field getter.
func (d *DataDict) GetListOfDoubleDef(key uint16, def []float64) []float64 {
	if elem, ok := (*d)[key]; ok {
		if elem.FType == List_of_Double {
			return elem.Value.([]float64)
		}
	}
	return def
}

// Convert list of elements to dict of elements.
func (d *Data) Dict() (dict DataDict) {
	dict = make(DataDict)
	for _, e := range *d {
		dict[e.Key] = e
	}
	return dict
}

// Add new element to data dictionary.
func (d *DataDict) Add(key uint16, ftype uint8, value interface{}) {
	(*d)[key] = &Elem{key, ftype, value}
}

// Encode data element to bytes.
func (e *Elem) Encode() ([]byte, error) {
	body, err := e.encodeValue()
	if err != nil {
		return nil, err
	}
	res := make([]byte, len(body)+5)
	copy(res[5:], body)
	binary.BigEndian.PutUint16(res[0:2], uint16(e.Key))
	res[2] = uint8(e.FType)
	binary.BigEndian.PutUint16(res[3:5], uint16(len(body)))
	return res, nil
}

func (e *Elem) WriteTo(writer io.Writer) (n int, err error) {
	encodedValue, err := e.encodeValue()
	if err != nil {
		return 0, err
	}
	header := make([]byte, 5)
	binary.BigEndian.PutUint16(header, uint16(e.Key))
	header[2] = uint8(e.FType)
	binary.BigEndian.PutUint16(header[3:], uint16(len(encodedValue)))
	n1, err := writer.Write(header)
	if err != nil {
		return n1, err
	}
	n2, err := writer.Write(encodedValue)
	if err != nil {
		return n2, err
	}
	return n1 + n2, nil
}

// Decode next element from byte slice.
func scan(bytes []byte) (elem *Elem, tail []byte, err error) {
	if len(bytes) == 0 {
		return nil, bytes, errors.New("EOF")
	}
	if len(bytes) < 5 {
		return nil, bytes, errors.New("decode: incomplete element header")
	}
	key := binary.BigEndian.Uint16(bytes[0:])
	ftype := bytes[2]
	body_len := binary.BigEndian.Uint16(bytes[3:])
	if len(bytes) < int(body_len)+5 {
		return nil, nil, fmt.Errorf(
			"decode: broken elem key#%d ftype=%d. expected body len %d but %d found",
			key, ftype, body_len, len(bytes)-5)
	}
	value, err := decodeValue(ftype, bytes[5:5+body_len])
	if err != nil {
		return nil, nil, err
	}
	return &Elem{key, ftype, value}, bytes[5+body_len:], nil
}

// Encode element value to bytes.
func (e *Elem) encodeValue() ([]byte, error) {
	encoded, err := encodeValue(e.FType, e.Value)
	if err != nil {
		return nil, fmt.Errorf("encode key#%d: %s", e.Key, err)
	}
	return encoded, nil
}
