package ktlv

import (
	"encoding/binary"
	"fmt"
	"io"
)

type Elem struct {
	Key   uint16
	FType uint8
	Value interface{}
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

// Search and decode one element with given key in octet stream.
func DecodeElem(b []byte, key uint16) (*Elem, error) {
	for {
		bLen := len(b)
		if bLen < 5 {
			break
		}
		fKey := binary.BigEndian.Uint16(b)
		fType := b[2]
		fLen := int(binary.BigEndian.Uint16(b[3:]))
		if tLen := bLen - 5; tLen < fLen {
			return nil, fmt.Errorf("broken "+
				"elem key#%d type=%d. expected body"+
				" len is %d but %d found",
				fKey, fType, fLen, tLen)
		}
		if fKey == key {
			value, err := decodeValue(fType, b[5:5+fLen])
			if err != nil {
				return nil, err
			}
			return &Elem{fKey, fType, value}, nil
		}
		b = b[5+fLen:]
	}
	return nil, ElementNotFound
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

// Encode element value to bytes.
func (e *Elem) encodeValue() ([]byte, error) {
	encoded, err := encodeValue(e.FType, e.Value)
	if err != nil {
		return nil, fmt.Errorf("encode key#%d: %s", e.Key, err)
	}
	return encoded, nil
}

// Check if elements are equal or not.
// Used in tests.
func (e1 *Elem) Equals(e2 *Elem) bool {
	if e1.Key != e2.Key || e1.FType != e2.FType {
		return false
	}
	switch e1.FType {
	case Bitmap:
		v1, _ := e1.Value.([]bool)
		v2, _ := e2.Value.([]bool)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_String:
		v1, _ := e1.Value.([]string)
		v2, _ := e2.Value.([]string)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Double:
		v1, _ := e1.Value.([]float64)
		v2, _ := e2.Value.([]float64)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Int8:
		v1, _ := e1.Value.([]int8)
		v2, _ := e2.Value.([]int8)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Int16:
		v1, _ := e1.Value.([]int16)
		v2, _ := e2.Value.([]int16)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Int24:
		v1, _ := e1.Value.([]int32)
		v2, _ := e2.Value.([]int32)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Int32:
		v1, _ := e1.Value.([]int32)
		v2, _ := e2.Value.([]int32)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Int64:
		v1, _ := e1.Value.([]int64)
		v2, _ := e2.Value.([]int64)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Uint8:
		v1, _ := e1.Value.([]uint8)
		v2, _ := e2.Value.([]uint8)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Uint16:
		v1, _ := e1.Value.([]uint16)
		v2, _ := e2.Value.([]uint16)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Uint24:
		v1, _ := e1.Value.([]uint32)
		v2, _ := e2.Value.([]uint32)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Uint32:
		v1, _ := e1.Value.([]uint32)
		v2, _ := e2.Value.([]uint32)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Uint64:
		v1, _ := e1.Value.([]uint64)
		v2, _ := e2.Value.([]uint64)
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
