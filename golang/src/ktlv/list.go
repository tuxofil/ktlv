package ktlv

import (
	"bytes"
)

type List []*Elem

// Encode input data to byte buffer.
func (d List) Encode() ([]byte, error) {
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

// Decode data from byte buffer.
func DecodeList(bytes []byte) (List, error) {
	res := make(List, 0)
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

// Convert list of elements to dict of elements.
func (d List) Dict() (dict Dict) {
	dict = make(Dict)
	for _, e := range d {
		dict[e.Key] = e
	}
	return dict
}
