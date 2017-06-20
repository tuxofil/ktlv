package ktlv

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"math"
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

type Key uint16

type FType uint8

type Elem struct {
	Key
	FType
	Value interface{}
}

type Data []*Elem

type DataDict map[Key]*Elem

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

// Encode input data to byte buffer.
// Deprecated API.
func Enc(data Data) []byte {
	bytes, err := data.Encode()
	if err != nil {
		log.Fatalf("deprecated API: %s", err)
	}
	return bytes
}

// Encode dictionary with input data to byte buffer.
// Deprecated API.
func Encd(data DataDict) []byte {
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

// Decode data from byte buffer.
// Deprecated API.
func Dec(bytes []byte) Data {
	data, err := Decode(bytes)
	if err != nil {
		log.Fatalf("deprecated API: %s", err)
	}
	return data
}

// Decode data from byte buffer to dictionary.
// Deprecated API.
func Decd(bytes []byte) DataDict {
	data, err := DecodeDict(bytes)
	if err != nil {
		log.Fatalf("deprecated API: %s", err)
	}
	return data
}

// String field getter.
func (d *DataDict) GetStringDef(key Key, def string) string {
	if elem, ok := (*d)[key]; ok {
		if elem.FType == String {
			return elem.Value.(string)
		}
	}
	return def
}

// bool field getter.
func (d *DataDict) GetBoolDef(key Key, def bool) bool {
	if elem, ok := (*d)[key]; ok {
		if elem.FType == Bool {
			return elem.Value.(bool)
		}
	}
	return def
}

// uint8 field getter.
func (d *DataDict) GetUint8Def(key Key, def uint8) uint8 {
	if elem, ok := (*d)[key]; ok {
		if elem.FType == Uint8 {
			return elem.Value.(uint8)
		}
	}
	return def
}

// uint16 field getter.
func (d *DataDict) GetUint16Def(key Key, def uint16) uint16 {
	if elem, ok := (*d)[key]; ok {
		if elem.FType == Uint16 {
			return elem.Value.(uint16)
		}
	}
	return def
}

// uint32 field getter.
func (d *DataDict) GetUint32Def(key Key, def uint32) uint32 {
	if elem, ok := (*d)[key]; ok {
		if elem.FType == Uint32 {
			return elem.Value.(uint32)
		}
	}
	return def
}

// uint64 field getter.
func (d *DataDict) GetUint64Def(key Key, def uint64) uint64 {
	if elem, ok := (*d)[key]; ok {
		if elem.FType == Uint64 {
			return elem.Value.(uint64)
		}
	}
	return def
}

// double field getter.
func (d *DataDict) GetDoubleDef(key Key, def float64) float64 {
	if elem, ok := (*d)[key]; ok {
		if elem.FType == Double {
			return elem.Value.(float64)
		}
	}
	return def
}

// list of uint32 field getter.
func (d *DataDict) GetListOfUint32Def(key Key, def []uint32) []uint32 {
	if elem, ok := (*d)[key]; ok {
		if elem.FType == List_of_Uint32 {
			return elem.Value.([]uint32)
		}
	}
	return def
}

// list of string field getter.
func (d *DataDict) GetListOfStringDef(key Key, def []string) []string {
	if elem, ok := (*d)[key]; ok {
		if elem.FType == List_of_String {
			return elem.Value.([]string)
		}
	}
	return def
}

// list of double field getter.
func (d *DataDict) GetListOfDoubleDef(key Key, def []float64) []float64 {
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
func (d *DataDict) Add(key Key, ftype FType, value interface{}) {
	(*d)[key] = &Elem{key, ftype, value}
}

// Check if elements are equal or not.
// Used in tests.
func (e1 *Elem) Equals(e2 *Elem) bool {
	if e1.Key != e2.Key || e1.FType != e2.FType {
		return false
	}
	switch e1.FType {
	case Bitmap:
		v1 := e1.Value.([]bool)
		v2 := e2.Value.([]bool)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_String:
		v1 := e1.Value.([]string)
		v2 := e2.Value.([]string)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Double:
		v1 := e1.Value.([]float64)
		v2 := e2.Value.([]float64)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Int8:
		v1 := e1.Value.([]int8)
		v2 := e2.Value.([]int8)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Int16:
		v1 := e1.Value.([]int16)
		v2 := e2.Value.([]int16)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Int24:
		v1 := e1.Value.([]int32)
		v2 := e2.Value.([]int32)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Int32:
		v1 := e1.Value.([]int32)
		v2 := e2.Value.([]int32)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Int64:
		v1 := e1.Value.([]int64)
		v2 := e2.Value.([]int64)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Uint8:
		v1 := e1.Value.([]uint8)
		v2 := e2.Value.([]uint8)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Uint16:
		v1 := e1.Value.([]uint16)
		v2 := e2.Value.([]uint16)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Uint24:
		v1 := e1.Value.([]uint32)
		v2 := e2.Value.([]uint32)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Uint32:
		v1 := e1.Value.([]uint32)
		v2 := e2.Value.([]uint32)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Uint64:
		v1 := e1.Value.([]uint64)
		v2 := e2.Value.([]uint64)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	default:
		return e1.Value == e2.Value
	}
	return true
}

// Encode data element to bytes.
// Deprecated API.
func (e *Elem) encode() []byte {
	encoded, err := e.Encode()
	if err != nil {
		log.Fatalf("deprecated API: %s", err)
	}
	return encoded
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

// Decode next element from byte slice.
func scan(bytes []byte) (elem *Elem, tail []byte, err error) {
	if len(bytes) == 0 {
		return nil, bytes, errors.New("EOF")
	}
	if len(bytes) < 5 {
		return nil, bytes, errors.New("decode: incomplete element header")
	}
	key := Key(binary.BigEndian.Uint16(bytes[0:2]))
	ftype := FType(bytes[2])
	body_len := binary.BigEndian.Uint16(bytes[3:5])
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
// Deprecated API
func encode_val(t FType, i interface{}) (r []byte) {
	elem := &Elem{
		FType: t,
		Value: i,
	}
	bytes, err := elem.encodeValue()
	if err != nil {
		log.Fatalf("deprecated API: %s", err)
	}
	return bytes
}

// Encode element value to bytes.
func (e *Elem) encodeValue() ([]byte, error) {
	switch e.FType {
	case Bool:
		switch e.Value.(type) {
		case bool:
		default:
			return nil, fmt.Errorf("encode: bad Bool key#%d: %#v (%T)",
				e.Key, e.Value, e.Value)
		}
		if e.Value.(bool) {
			return []byte{1}, nil
		} else {
			return []byte{0}, nil
		}
	case Uint8:
		switch e.Value.(type) {
		case uint8:
		default:
			return nil, fmt.Errorf("encode: bad Uint8 key#%d: %#v (%T)",
				e.Key, e.Value, e.Value)
		}
		return []byte{e.Value.(uint8)}, nil
	case Uint16:
		switch e.Value.(type) {
		case uint16:
		default:
			return nil, fmt.Errorf("encode: bad Uint16 key#%d: %#v (%T)",
				e.Key, e.Value, e.Value)
		}
		res := make([]byte, 2)
		binary.BigEndian.PutUint16(res, e.Value.(uint16))
		return res, nil
	case Uint24:
		switch e.Value.(type) {
		case uint32:
		default:
			return nil, fmt.Errorf("encode: bad Uint24 key#%d: %#v (%T)",
				e.Key, e.Value, e.Value)
		}
		res := make([]byte, 3)
		enc_uint24(res, e.Value.(uint32))
		return res, nil
	case Uint32:
		switch e.Value.(type) {
		case uint32:
		default:
			return nil, fmt.Errorf("encode: bad Uint32 key#%d: %#v (%T)",
				e.Key, e.Value, e.Value)
		}
		res := make([]byte, 4)
		binary.BigEndian.PutUint32(res, e.Value.(uint32))
		return res, nil
	case Uint64:
		switch e.Value.(type) {
		case uint64:
		default:
			return nil, fmt.Errorf("encode: bad Uint64 key#%d: %#v (%T)",
				e.Key, e.Value, e.Value)
		}
		res := make([]byte, 8)
		binary.BigEndian.PutUint64(res, e.Value.(uint64))
		return res, nil
	case Double:
		switch e.Value.(type) {
		case float64:
		default:
			return nil, fmt.Errorf("encode: bad Double key#%d: %#v (%T)",
				e.Key, e.Value, e.Value)
		}
		res := make([]byte, 8)
		binary.BigEndian.PutUint64(res, math.Float64bits(e.Value.(float64)))
		return res, nil
	case String:
		switch e.Value.(type) {
		case string:
		default:
			return nil, fmt.Errorf("encode: bad String key#%d: %#v (%T)",
				e.Key, e.Value, e.Value)
		}
		return []byte(e.Value.(string)), nil
	case Bitmap:
		switch e.Value.(type) {
		case []bool:
		default:
			return nil, fmt.Errorf("encode: bad Bitmap key#%d: %#v (%T)",
				e.Key, e.Value, e.Value)
		}
		v := e.Value.([]bool)
		l := len(v) / 8
		rem := len(v) % 8
		var unused uint8
		if 0 < rem {
			l++
			unused = 8 - uint8(rem)
		}
		res := make([]byte, l+1)
		res[0] = unused
		for i, b := range v {
			if !b {
				continue
			}
			major_bit_offset := int(unused) + i
			byte_offset := major_bit_offset / 8
			minor_bit_offset := major_bit_offset % 8
			mask := uint8(1 << (7 - uint8(minor_bit_offset)))
			res[byte_offset+1] |= mask
		}
		return res, nil
	case Int8:
		switch e.Value.(type) {
		case int8:
		default:
			return nil, fmt.Errorf("encode: bad Int8 key#%d: %#v (%T)",
				e.Key, e.Value, e.Value)
		}
		return []byte{uint8(e.Value.(int8))}, nil
	case Int16:
		switch e.Value.(type) {
		case int16:
		default:
			return nil, fmt.Errorf("encode: bad Int16 key#%d: %#v (%T)",
				e.Key, e.Value, e.Value)
		}
		res := make([]byte, 2)
		binary.BigEndian.PutUint16(res, uint16(e.Value.(int16)))
		return res, nil
	case Int24:
		switch e.Value.(type) {
		case int32:
		default:
			return nil, fmt.Errorf("encode: bad Int24 key#%d: %#v (%T)",
				e.Key, e.Value, e.Value)
		}
		res := make([]byte, 3)
		enc_int24(res, e.Value.(int32))
		return res, nil
	case Int32:
		switch e.Value.(type) {
		case int32:
		default:
			return nil, fmt.Errorf("encode: bad Int32 key#%d: %#v (%T)",
				e.Key, e.Value, e.Value)
		}
		res := make([]byte, 4)
		binary.BigEndian.PutUint32(res, uint32(e.Value.(int32)))
		return res, nil
	case Int64:
		switch e.Value.(type) {
		case int64:
		default:
			return nil, fmt.Errorf("encode: bad Int64 key#%d: %#v (%T)",
				e.Key, e.Value, e.Value)
		}
		res := make([]byte, 8)
		binary.BigEndian.PutUint64(res, uint64(e.Value.(int64)))
		return res, nil
	case List_of_String:
		switch e.Value.(type) {
		case []string:
		default:
			return nil, fmt.Errorf("encode: bad List_of_String key#%d: %#v (%T)",
				e.Key, e.Value, e.Value)
		}
		v := e.Value.([]string)
		tmp := make([][]byte, len(v)*2)
		for i, s := range v {
			tmp[i*2] = make([]byte, 2)
			bytes := []byte(s)
			tmp[i*2+1] = bytes
			binary.BigEndian.PutUint16(tmp[i*2], uint16(len(bytes)))
		}
		return bytes.Join(tmp, []byte{}), nil
	case List_of_Uint8:
		switch e.Value.(type) {
		case []uint8:
		default:
			return nil, fmt.Errorf("encode: bad List_of_Uint8 key#%d: %#v (%T)",
				e.Key, e.Value, e.Value)
		}
		return []byte(e.Value.([]uint8)), nil
	case List_of_Uint16:
		switch e.Value.(type) {
		case []uint16:
		default:
			return nil, fmt.Errorf("encode: bad List_of_Uint16 key#%d: %#v (%T)",
				e.Key, e.Value, e.Value)
		}
		v := e.Value.([]uint16)
		res := make([]byte, len(v)*2)
		for i, n := range v {
			binary.BigEndian.PutUint16(res[i*2:(i+1)*2], n)
		}
		return res, nil
	case List_of_Uint24:
		switch e.Value.(type) {
		case []uint32:
		default:
			return nil, fmt.Errorf("encode: bad List_of_Uint24 key#%d: %#v (%T)",
				e.Key, e.Value, e.Value)
		}
		v := e.Value.([]uint32)
		res := make([]byte, len(v)*3)
		for i, n := range v {
			enc_uint24(res[i*3:(i+1)*3], n)
		}
		return res, nil
	case List_of_Uint32:
		switch e.Value.(type) {
		case []uint32:
		default:
			return nil, fmt.Errorf("encode: bad List_of_Uint32 key#%d: %#v (%T)",
				e.Key, e.Value, e.Value)
		}
		v := e.Value.([]uint32)
		res := make([]byte, len(v)*4)
		for i, n := range v {
			binary.BigEndian.PutUint32(res[i*4:(i+1)*4], n)
		}
		return res, nil
	case List_of_Uint64:
		switch e.Value.(type) {
		case []uint64:
		default:
			return nil, fmt.Errorf("encode: bad List_of_Uint64 key#%d: %#v (%T)",
				e.Key, e.Value, e.Value)
		}
		v := e.Value.([]uint64)
		res := make([]byte, len(v)*8)
		for i, n := range v {
			binary.BigEndian.PutUint64(res[i*8:(i+1)*8], n)
		}
		return res, nil
	case List_of_Double:
		switch e.Value.(type) {
		case []float64:
		default:
			return nil, fmt.Errorf("encode: bad List_of_Double key#%d: %#v (%T)",
				e.Key, e.Value, e.Value)
		}
		v := e.Value.([]float64)
		res := make([]byte, len(v)*8)
		for i, n := range v {
			binary.BigEndian.PutUint64(res[i*8:(i+1)*8], math.Float64bits(n))
		}
		return res, nil
	case List_of_Int8:
		switch e.Value.(type) {
		case []int8:
		default:
			return nil, fmt.Errorf("encode: bad List_of_Int8 key#%d: %#v (%T)",
				e.Key, e.Value, e.Value)
		}
		v := e.Value.([]int8)
		res := make([]byte, len(v))
		for i, n := range v {
			res[i] = uint8(n)
		}
		return res, nil
	case List_of_Int16:
		switch e.Value.(type) {
		case []int16:
		default:
			return nil, fmt.Errorf("encode: bad List_of_Int16 key#%d: %#v (%T)",
				e.Key, e.Value, e.Value)
		}
		v := e.Value.([]int16)
		res := make([]byte, len(v)*2)
		for i, n := range v {
			binary.BigEndian.PutUint16(res[i*2:(i+1)*2], uint16(n))
		}
		return res, nil
	case List_of_Int24:
		switch e.Value.(type) {
		case []int32:
		default:
			return nil, fmt.Errorf("encode: bad List_of_Int24 key#%d: %#v (%T)",
				e.Key, e.Value, e.Value)
		}
		v := e.Value.([]int32)
		res := make([]byte, len(v)*3)
		for i, n := range v {
			enc_int24(res[i*3:(i+1)*3], n)
		}
		return res, nil
	case List_of_Int32:
		switch e.Value.(type) {
		case []int32:
		default:
			return nil, fmt.Errorf("encode: bad List_of_Int32 key#%d: %#v (%T)",
				e.Key, e.Value, e.Value)
		}
		v := e.Value.([]int32)
		res := make([]byte, len(v)*4)
		for i, n := range v {
			binary.BigEndian.PutUint32(res[i*4:(i+1)*4], uint32(n))
		}
		return res, nil
	case List_of_Int64:
		switch e.Value.(type) {
		case []int64:
		default:
			return nil, fmt.Errorf("encode: bad List_of_Int64 key#%d: %#v (%T)",
				e.Key, e.Value, e.Value)
		}
		v := e.Value.([]int64)
		res := make([]byte, len(v)*8)
		for i, n := range v {
			binary.BigEndian.PutUint64(res[i*8:(i+1)*8], uint64(n))
		}
		return res, nil
	}
	return nil, fmt.Errorf("encode: unknown field type for key#%d: %d", e.Key, e.FType)
}

// Decode element value from byte slice.
// Deprecated API.
func decode_val(t FType, bytes []byte) interface{} {
	res, err := decodeValue(t, bytes)
	if err != nil {
		log.Fatalf("deprecated API: %s", err)
	}
	return res
}

// Decode element value from byte slice.
func decodeValue(t FType, b []byte) (interface{}, error) {
	switch t {
	case Bool:
		if len(b) != 1 {
			return nil, fmt.Errorf("decode: bad Bool len: %d", len(b))
		}
		return b[0] == 1, nil
	case Uint8:
		if len(b) != 1 {
			return nil, fmt.Errorf("decode: bad Uint8 len: %d", len(b))
		}
		return b[0], nil
	case Uint16:
		if len(b) != 2 {
			return nil, fmt.Errorf("decode: bad Uint16 len: %d", len(b))
		}
		return binary.BigEndian.Uint16(b), nil
	case Uint24:
		if len(b) != 3 {
			return nil, fmt.Errorf("decode: bad Uint24 len: %d", len(b))
		}
		return dec_uint24(b), nil
	case Uint32:
		if len(b) != 4 {
			return nil, fmt.Errorf("decode: bad Uint32 len: %d", len(b))
		}
		return binary.BigEndian.Uint32(b), nil
	case Uint64:
		if len(b) != 8 {
			return nil, fmt.Errorf("decode: bad Uint64 len: %d", len(b))
		}
		return binary.BigEndian.Uint64(b), nil
	case Double:
		if len(b) != 8 {
			return nil, fmt.Errorf("decode: bad Double len: %d", len(b))
		}
		return math.Float64frombits(binary.BigEndian.Uint64(b)), nil
	case String:
		return string(b), nil
	case Bitmap:
		if len(b) == 0 {
			return nil, fmt.Errorf("decode: bad Bitmap len: %d", len(b))
		}
		unused := b[0]
		bit_len := (len(b)-1)*8 - int(unused)
		r := make([]bool, bit_len)
		for i := 0; i < len(r); i++ {
			major_bit_offset := int(unused) + i
			byte_offset := major_bit_offset / 8
			minor_bit_offset := major_bit_offset % 8
			mask := uint8(1 << (7 - uint8(minor_bit_offset)))
			r[i] = 0 < b[byte_offset+1]&mask
		}
		return r, nil
	case Int8:
		if len(b) != 1 {
			return nil, fmt.Errorf("decode: bad Int8 len: %d", len(b))
		}
		return int8(b[0]), nil
	case Int16:
		if len(b) != 2 {
			return nil, fmt.Errorf("decode: bad Int16 len: %d", len(b))
		}
		return int16(binary.BigEndian.Uint16(b)), nil
	case Int24:
		if len(b) != 3 {
			return nil, fmt.Errorf("decode: bad Int24 len: %d", len(b))
		}
		return dec_int24(b), nil
	case Int32:
		if len(b) != 4 {
			return nil, fmt.Errorf("decode: bad Int32 len: %d", len(b))
		}
		return int32(binary.BigEndian.Uint32(b)), nil
	case Int64:
		if len(b) != 8 {
			return nil, fmt.Errorf("decode: bad Int64 len: %d", len(b))
		}
		return int64(binary.BigEndian.Uint64(b)), nil
	case List_of_String:
		res := make([]string, 0)
		tail := b
		for 0 < len(tail) {
			if len(tail) < 2 {
				return nil, fmt.Errorf("decode: broken List_of_String (elem length)")
			}
			l := int(binary.BigEndian.Uint16(tail))
			if len(tail) < 2+l {
				return nil, fmt.Errorf("decode: broken List_of_String (elem value)")
			}
			res = append(res, string(tail[2:2+l]))
			tail = tail[2+l:]
		}
		return res, nil
	case List_of_Uint8:
		return []uint8(b), nil
	case List_of_Uint16:
		if len(b)%2 != 0 {
			return nil, fmt.Errorf("decode: bad List_of_Uint16 len: %d", len(b))
		}
		r := make([]uint16, len(b)/2)
		for i := 0; i < len(r); i++ {
			r[i] = binary.BigEndian.Uint16(b[i*2 : (i+1)*2])
		}
		return r, nil
	case List_of_Uint24:
		if len(b)%3 != 0 {
			return nil, fmt.Errorf("decode: bad List_of_Uint24 len: %d", len(b))
		}
		r := make([]uint32, len(b)/3)
		for i := 0; i < len(r); i++ {
			r[i] = dec_uint24(b[i*3 : (i+1)*3])
		}
		return r, nil
	case List_of_Uint32:
		if len(b)%4 != 0 {
			return nil, fmt.Errorf("decode: bad List_of_Uint32 len: %d", len(b))
		}
		r := make([]uint32, len(b)/4)
		for i := 0; i < len(r); i++ {
			r[i] = binary.BigEndian.Uint32(b[i*4 : (i+1)*4])
		}
		return r, nil
	case List_of_Uint64:
		if len(b)%8 != 0 {
			return nil, fmt.Errorf("decode: bad List_of_Uint64 len: %d", len(b))
		}
		r := make([]uint64, len(b)/8)
		for i := 0; i < len(r); i++ {
			r[i] = binary.BigEndian.Uint64(b[i*8 : (i+1)*8])
		}
		return r, nil
	case List_of_Double:
		if len(b)%8 != 0 {
			return nil, fmt.Errorf("decode: bad List_of_Double len: %d", len(b))
		}
		r := make([]float64, len(b)/8)
		for i := 0; i < len(r); i++ {
			r[i] = math.Float64frombits(binary.BigEndian.Uint64(b[i*8 : (i+1)*8]))
		}
		return r, nil
	case List_of_Int8:
		r := make([]int8, len(b))
		for i, n := range b {
			r[i] = int8(n)
		}
		return r, nil
	case List_of_Int16:
		if len(b)%2 != 0 {
			return nil, fmt.Errorf("decode: bad List_of_Int16 len: %d", len(b))
		}
		r := make([]int16, len(b)/2)
		for i := 0; i < len(r); i++ {
			r[i] = int16(binary.BigEndian.Uint16(b[i*2 : (i+1)*2]))
		}
		return r, nil
	case List_of_Int24:
		if len(b)%3 != 0 {
			return nil, fmt.Errorf("decode: bad List_of_Int24 len: %d", len(b))
		}
		r := make([]int32, len(b)/3)
		for i := 0; i < len(r); i++ {
			r[i] = dec_int24(b[i*3 : (i+1)*3])
		}
		return r, nil
	case List_of_Int32:
		if len(b)%4 != 0 {
			return nil, fmt.Errorf("decode: bad List_of_Int32 len: %d", len(b))
		}
		r := make([]int32, len(b)/4)
		for i := 0; i < len(r); i++ {
			r[i] = int32(binary.BigEndian.Uint32(b[i*4 : (i+1)*4]))
		}
		return r, nil
	case List_of_Int64:
		if len(b)%8 != 0 {
			return nil, fmt.Errorf("decode: bad List_of_Int64 len: %d", len(b))
		}
		r := make([]int64, len(b)/8)
		for i := 0; i < len(r); i++ {
			r[i] = int64(binary.BigEndian.Uint64(b[i*8 : (i+1)*8]))
		}
		return r, nil
	}
	return nil, fmt.Errorf("decode: unknown field type: %d", t)
}

// Encode unsigned int24 to byte slice.
func enc_uint24(a []byte, n uint32) {
	a[0] = uint8((n &^ 0xff000000) >> 16)
	binary.BigEndian.PutUint16(a[1:], uint16(n&^0xffff0000))
}

// Decode unsigned int24 from byte slice.
func dec_uint24(b []byte) uint32 {
	return (uint32(b[0]) << 16) | uint32(binary.BigEndian.Uint16(b[1:]))
}

// Encode signed int24 to byte slice.
func enc_int24(a []byte, n int32) {
	tmp := make([]byte, 4)
	binary.BigEndian.PutUint32(tmp, uint32(n))
	copy(a, tmp[1:])
}

// Decode signed int24 from byte slice.
func dec_int24(b []byte) int32 {
	major := int16(binary.BigEndian.Uint16(b[0:2]))
	return (int32(major) << 8) + int32(b[2])
}
