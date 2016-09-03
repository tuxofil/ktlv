package ktlv

import (
	"bytes"
	"encoding/binary"
	"errors"
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
func Enc(data Data) []byte {
	parts := make([][]byte, len(data))
	for k, v := range data {
		parts[k] = v.encode()
	}
	return bytes.Join(parts, []byte{})
}

// Encode dictionary with input data to byte buffer.
func Encd(data DataDict) []byte {
	parts := make([][]byte, len(data))
	i := 0
	for _, v := range data {
		parts[i] = v.encode()
		i++
	}
	return bytes.Join(parts, []byte{})
}

// Decode data from byte buffer.
func Dec(bytes []byte) Data {
	var elem *Elem
	var err error
	res := make(Data, 0, 100)
	for {
		elem, bytes, err = scan(bytes)
		if err != nil {
			break
		}
		res = append(res, elem)
	}
	return res
}

// Decode data from byte buffer to dictionary.
func Decd(bytes []byte) DataDict {
	var elem *Elem
	var err error
	var res = make(DataDict)
	for {
		elem, bytes, err = scan(bytes)
		if err != nil {
			break
		}
		res[elem.Key] = elem
	}
	return res
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
func (e *Elem) encode() []byte {
	body := encode_val(e.FType, e.Value)
	res := make([]byte, len(body)+5)
	copy(res[5:], body)
	binary.BigEndian.PutUint16(res[0:2], uint16(e.Key))
	res[2] = uint8(e.FType)
	binary.BigEndian.PutUint16(res[3:5], uint16(len(body)))
	return res
}

// Decode next element from byte slice.
func scan(bytes []byte) (elem *Elem, tail []byte, err error) {
	if len(bytes) == 0 {
		return nil, bytes, errors.New("EOF")
	}
	if len(bytes) < 5 {
		return nil, bytes, errors.New("incomplete element header")
	}
	key := Key(binary.BigEndian.Uint16(bytes[0:2]))
	ftype := FType(bytes[2])
	body_len := binary.BigEndian.Uint16(bytes[3:5])
	value := decode_val(ftype, bytes[5:5+body_len])
	return &Elem{key, ftype, value}, bytes[5+body_len:], nil
}

// Encode element value to bytes.
func encode_val(t FType, v0 interface{}) (r []byte) {
	switch t {
	case Bool:
		if v0.(bool) {
			r = []byte{1}
		} else {
			r = []byte{0}
		}
	case Uint8:
		r = []byte{v0.(uint8)}
	case Uint16:
		r = make([]byte, 2)
		binary.BigEndian.PutUint16(r, v0.(uint16))
	case Uint24:
		r = make([]byte, 3)
		enc_uint24(r, v0.(uint32))
	case Uint32:
		r = make([]byte, 4)
		binary.BigEndian.PutUint32(r, v0.(uint32))
	case Uint64:
		r = make([]byte, 8)
		binary.BigEndian.PutUint64(r, v0.(uint64))
	case Double:
		r = make([]byte, 8)
		binary.BigEndian.PutUint64(r, math.Float64bits(v0.(float64)))
	case String:
		r = []byte(v0.(string))
	case Bitmap:
		v := v0.([]bool)
		l := len(v) / 8
		rem := len(v) % 8
		var unused uint8
		if 0 < rem {
			l++
			unused = 8 - uint8(rem)
		}
		r = make([]byte, l+1)
		r[0] = unused
		for i, b := range v {
			if !b {
				continue
			}
			major_bit_offset := int(unused) + i
			byte_offset := major_bit_offset / 8
			minor_bit_offset := major_bit_offset % 8
			mask := uint8(1 << (7 - uint8(minor_bit_offset)))
			r[byte_offset+1] |= mask
		}
	case Int8:
		r = []byte{uint8(v0.(int8))}
	case Int16:
		r = make([]byte, 2)
		binary.BigEndian.PutUint16(r, uint16(v0.(int16)))
	case Int24:
		r = make([]byte, 3)
		enc_int24(r, v0.(int32))
	case Int32:
		r = make([]byte, 4)
		binary.BigEndian.PutUint32(r, uint32(v0.(int32)))
	case Int64:
		r = make([]byte, 8)
		binary.BigEndian.PutUint64(r, uint64(v0.(int64)))
	case List_of_String:
		v := v0.([]string)
		tmp := make([][]byte, len(v)*2)
		for i, s := range v {
			tmp[i*2] = make([]byte, 2)
			bytes := []byte(s)
			tmp[i*2+1] = bytes
			binary.BigEndian.PutUint16(tmp[i*2], uint16(len(bytes)))
		}
		r = bytes.Join(tmp, []byte{})
	case List_of_Uint8:
		r = []byte(v0.([]uint8))
	case List_of_Uint16:
		v := v0.([]uint16)
		r = make([]byte, len(v)*2)
		for i, n := range v {
			binary.BigEndian.PutUint16(r[i*2:(i+1)*2], n)
		}
	case List_of_Uint24:
		v := v0.([]uint32)
		r = make([]byte, len(v)*3)
		for i, n := range v {
			enc_uint24(r[i*3:(i+1)*3], n)
		}
	case List_of_Uint32:
		v := v0.([]uint32)
		r = make([]byte, len(v)*4)
		for i, n := range v {
			binary.BigEndian.PutUint32(r[i*4:(i+1)*4], n)
		}
	case List_of_Uint64:
		v := v0.([]uint64)
		r = make([]byte, len(v)*8)
		for i, n := range v {
			binary.BigEndian.PutUint64(r[i*8:(i+1)*8], n)
		}
	case List_of_Double:
		v := v0.([]float64)
		r = make([]byte, len(v)*8)
		for i, n := range v {
			binary.BigEndian.PutUint64(r[i*8:(i+1)*8], math.Float64bits(n))
		}
	case List_of_Int8:
		v := v0.([]int8)
		r = make([]byte, len(v))
		for i, n := range v {
			r[i] = uint8(n)
		}
	case List_of_Int16:
		v := v0.([]int16)
		r = make([]byte, len(v)*2)
		for i, n := range v {
			binary.BigEndian.PutUint16(r[i*2:(i+1)*2], uint16(n))
		}
	case List_of_Int24:
		v := v0.([]int32)
		r = make([]byte, len(v)*3)
		for i, n := range v {
			enc_int24(r[i*3:(i+1)*3], n)
		}
	case List_of_Int32:
		v := v0.([]int32)
		r = make([]byte, len(v)*4)
		for i, n := range v {
			binary.BigEndian.PutUint32(r[i*4:(i+1)*4], uint32(n))
		}
	case List_of_Int64:
		v := v0.([]int64)
		r = make([]byte, len(v)*8)
		for i, n := range v {
			binary.BigEndian.PutUint64(r[i*8:(i+1)*8], uint64(n))
		}
	default:
		log.Fatalf("ktlv encoder> unknown field type: %v", t)
	}
	return r
}

// Decode element value from byte slice.
func decode_val(t FType, b []byte) interface{} {
	var value interface{} = nil
	switch t {
	case Bool:
		return b[0] == 1
	case Uint8:
		return b[0]
	case Uint16:
		return binary.BigEndian.Uint16(b)
	case Uint24:
		return dec_uint24(b)
	case Uint32:
		return binary.BigEndian.Uint32(b)
	case Uint64:
		return binary.BigEndian.Uint64(b)
	case Double:
		return math.Float64frombits(binary.BigEndian.Uint64(b))
	case String:
		return string(b)
	case Bitmap:
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
		return r
	case Int8:
		return int8(b[0])
	case Int16:
		return int16(binary.BigEndian.Uint16(b))
	case Int24:
		return dec_int24(b)
	case Int32:
		return int32(binary.BigEndian.Uint32(b))
	case Int64:
		return int64(binary.BigEndian.Uint64(b))
	case List_of_String:
		r := make([]string, 0)
		for i := 0; i < len(b); {
			l := int(binary.BigEndian.Uint16(b[i : i+2]))
			r = append(r, string(b[i+2:i+2+l]))
			i += l + 2
		}
		return r
	case List_of_Uint8:
		return []uint8(b)
	case List_of_Uint16:
		r := make([]uint16, len(b)/2)
		for i := 0; i < len(r); i++ {
			r[i] = binary.BigEndian.Uint16(b[i*2 : (i+1)*2])
		}
		return r
	case List_of_Uint24:
		r := make([]uint32, len(b)/3)
		for i := 0; i < len(r); i++ {
			r[i] = dec_uint24(b[i*3 : (i+1)*3])
		}
		return r
	case List_of_Uint32:
		r := make([]uint32, len(b)/4)
		for i := 0; i < len(r); i++ {
			r[i] = binary.BigEndian.Uint32(b[i*4 : (i+1)*4])
		}
		return r
	case List_of_Uint64:
		r := make([]uint64, len(b)/8)
		for i := 0; i < len(r); i++ {
			r[i] = binary.BigEndian.Uint64(b[i*8 : (i+1)*8])
		}
		return r
	case List_of_Double:
		r := make([]float64, len(b)/8)
		for i := 0; i < len(r); i++ {
			r[i] = math.Float64frombits(binary.BigEndian.Uint64(b[i*8 : (i+1)*8]))
		}
		return r
	case List_of_Int8:
		r := make([]int8, len(b))
		for i, n := range b {
			r[i] = int8(n)
		}
		return r
	case List_of_Int16:
		r := make([]int16, len(b)/2)
		for i := 0; i < len(r); i++ {
			r[i] = int16(binary.BigEndian.Uint16(b[i*2 : (i+1)*2]))
		}
		return r
	case List_of_Int24:
		r := make([]int32, len(b)/3)
		for i := 0; i < len(r); i++ {
			r[i] = dec_int24(b[i*3 : (i+1)*3])
		}
		return r
	case List_of_Int32:
		r := make([]int32, len(b)/4)
		for i := 0; i < len(r); i++ {
			r[i] = int32(binary.BigEndian.Uint32(b[i*4 : (i+1)*4]))
		}
		return r
	case List_of_Int64:
		r := make([]int64, len(b)/8)
		for i := 0; i < len(r); i++ {
			r[i] = int64(binary.BigEndian.Uint64(b[i*8 : (i+1)*8]))
		}
		return r
	default:
		log.Fatalf("ktlv decoder> unknown field type: %v", t)
	}
	return value
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
