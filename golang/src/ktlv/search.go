package ktlv

import (
	"errors"
	"fmt"
)

// Search specific field in KTLV-encoded message without
// decoding it to the end.
func Search(encoded []byte, key uint16, max int) (*Elem, error) {
	for ; 0 < max; max-- {
		elem, tail, err := scan(encoded)
		if err != nil {
			return nil, err
		}
		if elem.Key == key {
			return elem, nil
		}
		encoded = tail
	}
	return nil, nil
}

// Make search as Search() function does, assume uint64 value.
func SearchUint64(encoded []byte, key uint16, max int) (uint64, error) {
	elem, err := Search(encoded, key, max)
	if err != nil {
		return 0, err
	} else if elem != nil {
		if n, ok := elem.Value.(uint64); ok {
			return n, nil
		} else {
			return 0, fmt.Errorf("expected uint64 but found: %T",
				elem.Value)
		}
	}
	return 0, errors.New("not found")
}
