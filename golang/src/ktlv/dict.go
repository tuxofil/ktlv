package ktlv

import (
	"bytes"
	"fmt"
)

type Dict map[uint16]*Elem

// Encode dictionary with input data to byte buffer.
func (d Dict) Encode() ([]byte, error) {
	buffer := &bytes.Buffer{}
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

// Decode data from byte buffer to dictionary.
// On error returns non nil value with all successfully decoded
// elements.
func DecodeDict(bytes []byte) (Dict, error) {
	res := Dict{}
	for 0 < len(bytes) {
		elem, tail, err := scan(bytes)
		if err != nil {
			return res, err
		}
		res[elem.Key] = elem
		bytes = tail
	}
	return res, nil
}

// Add new element to data dictionary.
func (d Dict) Add(key uint16, ftype uint8, value interface{}) {
	d[key] = &Elem{key, ftype, value}
}

func (d Dict) Get(key uint16) (ftype uint8, v interface{}, ok bool) {
	if elem, ok := d[key]; ok {
		return elem.FType, elem.Value, true
	}
	return 0, nil, false
}

func (d Dict) String() string {
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

// String field getter.
func (d Dict) GetString(key uint16) (string, error) {
	if elem, ok := d[key]; ok {
		if elem.FType == String {
			return elem.Value.(string), nil
		}
		return "", TypeAssertionFailed
	}
	return "", ElementNotFound
}

// String field getter.
func (d Dict) GetStringDef(key uint16, def string) string {
	if v, err := d.GetString(key); err == nil {
		return v
	}
	return def
}

// bool field getter.
func (d Dict) GetBoolDef(key uint16, def bool) bool {
	if elem, ok := d[key]; ok {
		if elem.FType == Bool {
			return elem.Value.(bool)
		}
	}
	return def
}

// uint8 field getter.
func (d Dict) GetUint8(key uint16) (uint8, error) {
	if elem, ok := d[key]; ok {
		if elem.FType == Uint8 {
			return elem.Value.(uint8), nil
		}
		return 0, TypeAssertionFailed
	}
	return 0, ElementNotFound
}

// uint8 field getter.
func (d Dict) GetUint8Def(key uint16, def uint8) uint8 {
	if elem, ok := d[key]; ok {
		if elem.FType == Uint8 {
			return elem.Value.(uint8)
		}
	}
	return def
}

// uint16 field getter.
func (d Dict) GetUint16Def(key uint16, def uint16) uint16 {
	if elem, ok := d[key]; ok {
		if elem.FType == Uint16 {
			return elem.Value.(uint16)
		}
	}
	return def
}

// uint32 field getter.
func (d Dict) GetUint32(key uint16) (uint32, error) {
	if elem, ok := d[key]; ok {
		if elem.FType == Uint32 {
			return elem.Value.(uint32), nil
		}
		return 0, TypeAssertionFailed
	}
	return 0, ElementNotFound
}

// uint32 field getter.
func (d Dict) GetUint32Def(key uint16, def uint32) uint32 {
	if elem, ok := d[key]; ok {
		if elem.FType == Uint32 {
			return elem.Value.(uint32)
		}
	}
	return def
}

// uint64 field getter.
func (d Dict) GetUint64(key uint16) (uint64, error) {
	if elem, ok := d[key]; ok {
		if elem.FType == Uint64 {
			return elem.Value.(uint64), nil
		}
		return 0, TypeAssertionFailed
	}
	return 0, ElementNotFound
}

// uint64 field getter.
func (d Dict) GetUint64Def(key uint16, def uint64) uint64 {
	if v, err := d.GetUint64(key); err == nil {
		return v
	}
	return def
}

// double field getter.
func (d Dict) GetDouble(key uint16) (float64, error) {
	if elem, ok := d[key]; ok {
		if elem.FType == Double {
			return elem.Value.(float64), nil
		}
		return 0, TypeAssertionFailed
	}
	return 0, ElementNotFound
}

// double field getter.
func (d Dict) GetDoubleDef(key uint16, def float64) float64 {
	if v, err := d.GetDouble(key); err == nil {
		return v
	}
	return def
}

// list of uint8 field getter.
func (d Dict) GetListOfUint8Def(key uint16, def []uint8) []uint8 {
	if elem, ok := d[key]; ok {
		if elem.FType == List_of_Uint8 {
			return elem.Value.([]uint8)
		}
	}
	return def
}

// list of uint32 field getter.
func (d Dict) GetListOfUint32Def(key uint16, def []uint32) []uint32 {
	if elem, ok := d[key]; ok {
		if elem.FType == List_of_Uint32 {
			return elem.Value.([]uint32)
		}
	}
	return def
}

// list of uint64 field getter.
func (d Dict) GetListOfUint64Def(key uint16, def []uint64) []uint64 {
	if elem, ok := d[key]; ok {
		if elem.FType == List_of_Uint64 {
			return elem.Value.([]uint64)
		}
	}
	return def
}

// list of string field getter.
func (d Dict) GetListOfStringDef(key uint16, def []string) []string {
	if elem, ok := d[key]; ok {
		if elem.FType == List_of_String {
			return elem.Value.([]string)
		}
	}
	return def
}

// list of double field getter.
func (d Dict) GetListOfDoubleDef(key uint16, def []float64) []float64 {
	if elem, ok := d[key]; ok {
		if elem.FType == List_of_Double {
			return elem.Value.([]float64)
		}
	}
	return def
}
