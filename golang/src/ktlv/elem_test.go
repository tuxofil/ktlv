package ktlv

import "testing"

func TestDecodeElem(t *testing.T) {
	testset := []struct {
		List      List
		SearchKey uint16
		Expect    *Elem
		Error     bool
	}{
		{List{}, 1, nil, true},
		{List{
			&Elem{0, Bool, true},
			&Elem{1, Uint8, uint8(1)},
			&Elem{2, Uint16, uint16(0xffff)},
		},
			3, nil, true},
		{List{
			&Elem{0, Bool, true},
			&Elem{1, Uint8, uint8(1)},
			&Elem{2, Uint16, uint16(0xffff)},
		},
			0,
			&Elem{0, Bool, true},
			false},
		{List{
			&Elem{0, Bool, true},
			&Elem{1, Uint8, uint8(1)},
			&Elem{2, Uint16, uint16(0xffff)},
		},
			1,
			&Elem{1, Uint8, uint8(1)},
			false},
		{List{
			&Elem{0, Bool, true},
			&Elem{1, Uint8, uint8(1)},
			&Elem{2, Uint16, uint16(0xffff)},
		},
			2,
			&Elem{2, Uint16, uint16(0xffff)},
			false},
	}
	for n, test := range testset {
		encoded, err := test.List.Encode()
		if err != nil {
			t.Fatalf("#%d> encode: %s", n, err)
		}
		elem, err := DecodeElem(encoded, test.SearchKey)
		if test.Error {
			if err == nil {
				t.Errorf("#%d> expected error but"+
					" decoding succeeded", n)
			}
		} else {
			if err != nil {
				t.Errorf("#%d> unexpected decode"+
					" error: %s", n, err)
			} else if !elem.Equals(test.Expect) {
				t.Errorf("#%d> unexpected result"+
					": %+v (expected %+v)",
					n, elem, test.Expect)
			}
		}
	}
}
