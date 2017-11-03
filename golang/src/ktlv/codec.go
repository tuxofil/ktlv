package ktlv

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
)

// Encode element value to bytes.
func encodeValue(ftype uint8, value interface{}) ([]byte, error) {
	switch ftype {
	case Bool:
		switch value.(type) {
		case bool:
		default:
			return nil, fmt.Errorf("bad Bool: %#v (%T)",
				value, value)
		}
		if value.(bool) {
			return []byte{1}, nil
		} else {
			return []byte{0}, nil
		}
	case Uint8:
		switch value.(type) {
		case uint8:
		default:
			return nil, fmt.Errorf("bad Uint8: %#v (%T)",
				value, value)
		}
		return []byte{value.(uint8)}, nil
	case Uint16:
		switch value.(type) {
		case uint16:
		default:
			return nil, fmt.Errorf("bad Uint16: %#v (%T)",
				value, value)
		}
		res := make([]byte, 2)
		binary.BigEndian.PutUint16(res, value.(uint16))
		return res, nil
	case Uint24:
		switch value.(type) {
		case uint32:
		default:
			return nil, fmt.Errorf("bad Uint24: %#v (%T)",
				value, value)
		}
		res := make([]byte, 3)
		enc_uint24(res, value.(uint32))
		return res, nil
	case Uint32:
		switch value.(type) {
		case uint32:
		default:
			return nil, fmt.Errorf("bad Uint32: %#v (%T)",
				value, value)
		}
		res := make([]byte, 4)
		binary.BigEndian.PutUint32(res, value.(uint32))
		return res, nil
	case Uint64:
		switch value.(type) {
		case uint64:
		default:
			return nil, fmt.Errorf("bad Uint64: %#v (%T)",
				value, value)
		}
		res := make([]byte, 8)
		binary.BigEndian.PutUint64(res, value.(uint64))
		return res, nil
	case Double:
		switch value.(type) {
		case float64:
		default:
			return nil, fmt.Errorf("bad Double: %#v (%T)",
				value, value)
		}
		res := make([]byte, 8)
		binary.BigEndian.PutUint64(res, math.Float64bits(value.(float64)))
		return res, nil
	case String:
		switch value.(type) {
		case string:
		default:
			return nil, fmt.Errorf("bad String: %#v (%T)",
				value, value)
		}
		return []byte(value.(string)), nil
	case Bitmap:
		switch value.(type) {
		case []bool:
		default:
			return nil, fmt.Errorf("bad Bitmap: %#v (%T)",
				value, value)
		}
		v := value.([]bool)
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
		switch value.(type) {
		case int8:
		default:
			return nil, fmt.Errorf("bad Int8: %#v (%T)",
				value, value)
		}
		return []byte{uint8(value.(int8))}, nil
	case Int16:
		switch value.(type) {
		case int16:
		default:
			return nil, fmt.Errorf("bad Int16: %#v (%T)",
				value, value)
		}
		res := make([]byte, 2)
		binary.BigEndian.PutUint16(res, uint16(value.(int16)))
		return res, nil
	case Int24:
		switch value.(type) {
		case int32:
		default:
			return nil, fmt.Errorf("bad Int24: %#v (%T)",
				value, value)
		}
		res := make([]byte, 3)
		enc_int24(res, value.(int32))
		return res, nil
	case Int32:
		switch value.(type) {
		case int32:
		default:
			return nil, fmt.Errorf("bad Int32: %#v (%T)",
				value, value)
		}
		res := make([]byte, 4)
		binary.BigEndian.PutUint32(res, uint32(value.(int32)))
		return res, nil
	case Int64:
		switch value.(type) {
		case int64:
		default:
			return nil, fmt.Errorf("bad Int64: %#v (%T)",
				value, value)
		}
		res := make([]byte, 8)
		binary.BigEndian.PutUint64(res, uint64(value.(int64)))
		return res, nil
	case List_of_String:
		switch value.(type) {
		case []string:
		default:
			return nil, fmt.Errorf("bad List_of_String: %#v (%T)",
				value, value)
		}
		v := value.([]string)
		tmp := make([][]byte, len(v)*2)
		for i, s := range v {
			tmp[i*2] = make([]byte, 2)
			bytes := []byte(s)
			tmp[i*2+1] = bytes
			binary.BigEndian.PutUint16(tmp[i*2], uint16(len(bytes)))
		}
		return bytes.Join(tmp, []byte{}), nil
	case List_of_Uint8:
		switch value.(type) {
		case []uint8:
		default:
			return nil, fmt.Errorf("bad List_of_Uint8: %#v (%T)",
				value, value)
		}
		return []byte(value.([]uint8)), nil
	case List_of_Uint16:
		switch value.(type) {
		case []uint16:
		default:
			return nil, fmt.Errorf("bad List_of_Uint16: %#v (%T)",
				value, value)
		}
		v := value.([]uint16)
		res := make([]byte, len(v)*2)
		for i, n := range v {
			binary.BigEndian.PutUint16(res[i*2:(i+1)*2], n)
		}
		return res, nil
	case List_of_Uint24:
		switch value.(type) {
		case []uint32:
		default:
			return nil, fmt.Errorf("bad List_of_Uint24: %#v (%T)",
				value, value)
		}
		v := value.([]uint32)
		res := make([]byte, len(v)*3)
		for i, n := range v {
			enc_uint24(res[i*3:(i+1)*3], n)
		}
		return res, nil
	case List_of_Uint32:
		switch value.(type) {
		case []uint32:
		default:
			return nil, fmt.Errorf("bad List_of_Uint32: %#v (%T)",
				value, value)
		}
		v := value.([]uint32)
		res := make([]byte, len(v)*4)
		for i, n := range v {
			binary.BigEndian.PutUint32(res[i*4:(i+1)*4], n)
		}
		return res, nil
	case List_of_Uint64:
		switch value.(type) {
		case []uint64:
		default:
			return nil, fmt.Errorf("bad List_of_Uint64: %#v (%T)",
				value, value)
		}
		v := value.([]uint64)
		res := make([]byte, len(v)*8)
		for i, n := range v {
			binary.BigEndian.PutUint64(res[i*8:(i+1)*8], n)
		}
		return res, nil
	case List_of_Double:
		switch value.(type) {
		case []float64:
		default:
			return nil, fmt.Errorf("bad List_of_Double: %#v (%T)",
				value, value)
		}
		v := value.([]float64)
		res := make([]byte, len(v)*8)
		for i, n := range v {
			binary.BigEndian.PutUint64(res[i*8:(i+1)*8], math.Float64bits(n))
		}
		return res, nil
	case List_of_Int8:
		switch value.(type) {
		case []int8:
		default:
			return nil, fmt.Errorf("bad List_of_Int8: %#v (%T)",
				value, value)
		}
		v := value.([]int8)
		res := make([]byte, len(v))
		for i, n := range v {
			res[i] = uint8(n)
		}
		return res, nil
	case List_of_Int16:
		switch value.(type) {
		case []int16:
		default:
			return nil, fmt.Errorf("bad List_of_Int16: %#v (%T)",
				value, value)
		}
		v := value.([]int16)
		res := make([]byte, len(v)*2)
		for i, n := range v {
			binary.BigEndian.PutUint16(res[i*2:(i+1)*2], uint16(n))
		}
		return res, nil
	case List_of_Int24:
		switch value.(type) {
		case []int32:
		default:
			return nil, fmt.Errorf("bad List_of_Int24: %#v (%T)",
				value, value)
		}
		v := value.([]int32)
		res := make([]byte, len(v)*3)
		for i, n := range v {
			enc_int24(res[i*3:(i+1)*3], n)
		}
		return res, nil
	case List_of_Int32:
		switch value.(type) {
		case []int32:
		default:
			return nil, fmt.Errorf("bad List_of_Int32: %#v (%T)",
				value, value)
		}
		v := value.([]int32)
		res := make([]byte, len(v)*4)
		for i, n := range v {
			binary.BigEndian.PutUint32(res[i*4:(i+1)*4], uint32(n))
		}
		return res, nil
	case List_of_Int64:
		switch value.(type) {
		case []int64:
		default:
			return nil, fmt.Errorf("bad List_of_Int64: %#v (%T)",
				value, value)
		}
		v := value.([]int64)
		res := make([]byte, len(v)*8)
		for i, n := range v {
			binary.BigEndian.PutUint64(res[i*8:(i+1)*8], uint64(n))
		}
		return res, nil
	}
	return nil, fmt.Errorf("unknown field type: %d", ftype)
}

// Decode element value from byte slice.
func decodeValue(t uint8, b []byte) (interface{}, error) {
	switch t {
	case Bool:
		if len(b) != 1 {
			return nil, fmt.Errorf("bad Bool len: %d", len(b))
		}
		return b[0] == 1, nil
	case Uint8:
		if len(b) != 1 {
			return nil, fmt.Errorf("bad Uint8 len: %d", len(b))
		}
		return b[0], nil
	case Uint16:
		if len(b) != 2 {
			return nil, fmt.Errorf("bad Uint16 len: %d", len(b))
		}
		return binary.BigEndian.Uint16(b), nil
	case Uint24:
		if len(b) != 3 {
			return nil, fmt.Errorf("bad Uint24 len: %d", len(b))
		}
		return dec_uint24(b), nil
	case Uint32:
		if len(b) != 4 {
			return nil, fmt.Errorf("bad Uint32 len: %d", len(b))
		}
		return binary.BigEndian.Uint32(b), nil
	case Uint64:
		if len(b) != 8 {
			return nil, fmt.Errorf("bad Uint64 len: %d", len(b))
		}
		return binary.BigEndian.Uint64(b), nil
	case Double:
		if len(b) != 8 {
			return nil, fmt.Errorf("bad Double len: %d", len(b))
		}
		return math.Float64frombits(binary.BigEndian.Uint64(b)), nil
	case String:
		return string(b), nil
	case Bitmap:
		if len(b) == 0 {
			return nil, fmt.Errorf("bad Bitmap len: %d", len(b))
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
			return nil, fmt.Errorf("bad Int8 len: %d", len(b))
		}
		return int8(b[0]), nil
	case Int16:
		if len(b) != 2 {
			return nil, fmt.Errorf("bad Int16 len: %d", len(b))
		}
		return int16(binary.BigEndian.Uint16(b)), nil
	case Int24:
		if len(b) != 3 {
			return nil, fmt.Errorf("bad Int24 len: %d", len(b))
		}
		return dec_int24(b), nil
	case Int32:
		if len(b) != 4 {
			return nil, fmt.Errorf("bad Int32 len: %d", len(b))
		}
		return int32(binary.BigEndian.Uint32(b)), nil
	case Int64:
		if len(b) != 8 {
			return nil, fmt.Errorf("bad Int64 len: %d", len(b))
		}
		return int64(binary.BigEndian.Uint64(b)), nil
	case List_of_String:
		res := make([]string, 0)
		tail := b
		for 0 < len(tail) {
			if len(tail) < 2 {
				return nil, fmt.Errorf("broken List_of_String (elem length)")
			}
			l := int(binary.BigEndian.Uint16(tail))
			if len(tail) < 2+l {
				return nil, fmt.Errorf("broken List_of_String (elem value)")
			}
			res = append(res, string(tail[2:2+l]))
			tail = tail[2+l:]
		}
		return res, nil
	case List_of_Uint8:
		return []uint8(b), nil
	case List_of_Uint16:
		if len(b)%2 != 0 {
			return nil, fmt.Errorf("bad List_of_Uint16 len: %d", len(b))
		}
		r := make([]uint16, len(b)/2)
		for i := 0; i < len(r); i++ {
			r[i] = binary.BigEndian.Uint16(b[i*2 : (i+1)*2])
		}
		return r, nil
	case List_of_Uint24:
		if len(b)%3 != 0 {
			return nil, fmt.Errorf("bad List_of_Uint24 len: %d", len(b))
		}
		r := make([]uint32, len(b)/3)
		for i := 0; i < len(r); i++ {
			r[i] = dec_uint24(b[i*3 : (i+1)*3])
		}
		return r, nil
	case List_of_Uint32:
		if len(b)%4 != 0 {
			return nil, fmt.Errorf("bad List_of_Uint32 len: %d", len(b))
		}
		r := make([]uint32, len(b)/4)
		for i := 0; i < len(r); i++ {
			r[i] = binary.BigEndian.Uint32(b[i*4 : (i+1)*4])
		}
		return r, nil
	case List_of_Uint64:
		if len(b)%8 != 0 {
			return nil, fmt.Errorf("bad List_of_Uint64 len: %d", len(b))
		}
		r := make([]uint64, len(b)/8)
		for i := 0; i < len(r); i++ {
			r[i] = binary.BigEndian.Uint64(b[i*8 : (i+1)*8])
		}
		return r, nil
	case List_of_Double:
		if len(b)%8 != 0 {
			return nil, fmt.Errorf("bad List_of_Double len: %d", len(b))
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
			return nil, fmt.Errorf("bad List_of_Int16 len: %d", len(b))
		}
		r := make([]int16, len(b)/2)
		for i := 0; i < len(r); i++ {
			r[i] = int16(binary.BigEndian.Uint16(b[i*2 : (i+1)*2]))
		}
		return r, nil
	case List_of_Int24:
		if len(b)%3 != 0 {
			return nil, fmt.Errorf("bad List_of_Int24 len: %d", len(b))
		}
		r := make([]int32, len(b)/3)
		for i := 0; i < len(r); i++ {
			r[i] = dec_int24(b[i*3 : (i+1)*3])
		}
		return r, nil
	case List_of_Int32:
		if len(b)%4 != 0 {
			return nil, fmt.Errorf("bad List_of_Int32 len: %d", len(b))
		}
		r := make([]int32, len(b)/4)
		for i := 0; i < len(r); i++ {
			r[i] = int32(binary.BigEndian.Uint32(b[i*4 : (i+1)*4]))
		}
		return r, nil
	case List_of_Int64:
		if len(b)%8 != 0 {
			return nil, fmt.Errorf("bad List_of_Int64 len: %d", len(b))
		}
		r := make([]int64, len(b)/8)
		for i := 0; i < len(r); i++ {
			r[i] = int64(binary.BigEndian.Uint64(b[i*8 : (i+1)*8]))
		}
		return r, nil
	}
	return nil, fmt.Errorf("unknown field type: %d", t)
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

// Decode next element from byte slice.
func scan(bytes []byte) (elem *Elem, tail []byte, err error) {
	if len(bytes) == 0 {
		return nil, bytes, errors.New("EOF")
	}
	if len(bytes) < 5 {
		return nil, bytes,
			errors.New("decode: incomplete element header")
	}
	key := binary.BigEndian.Uint16(bytes[0:])
	ftype := bytes[2]
	body_len := binary.BigEndian.Uint16(bytes[3:])
	if len(bytes) < int(body_len)+5 {
		return nil, nil, fmt.Errorf("decode: broken "+
			"elem key#%d ftype=%d. expected body"+
			" len %d but %d found",
			key, ftype, body_len, len(bytes)-5)
	}
	value, err := decodeValue(ftype, bytes[5:5+body_len])
	if err != nil {
		return nil, nil, err
	}
	return &Elem{key, ftype, value}, bytes[5+body_len:], nil
}
