package ktlv

import "testing"

func TestSearchUint64(t *testing.T) {
	testset := []struct {
		List        List
		SearchKey   uint16
		Max         int
		ExpectValue uint64
		ExpectError bool
	}{
		{List{
			&Elem{0, Uint64, uint64(1)},
		}, 0, 10, 1, false},
		{List{
			&Elem{0, Uint64, uint64(1)},
			&Elem{1, Uint64, uint64(2)},
		}, 1, 10, 2, false},
		{List{
			&Elem{0, Uint64, uint64(1)},
			&Elem{1, String, "abcdef"},
			&Elem{2, Uint64, uint64(3)},
		}, 2, 3, 3, false},
		{List{
			&Elem{0, Uint64, uint64(1)},
			&Elem{1, String, "abcdef"},
			&Elem{2, Uint64, uint64(3)},
			&Elem{3, Uint64, uint64(4)},
		}, 3, 3, 0, true},
		{List{
			&Elem{0, Uint64, uint64(1)},
			&Elem{1, String, "abcdef"},
			&Elem{2, Uint64, uint64(3)},
			&Elem{3, Uint64, uint64(4)},
		}, 3, 4, 4, false},
		{List{
			&Elem{0, Uint64, uint64(1)},
			&Elem{1, String, "abcdef"},
			&Elem{2, Uint64, uint64(3)},
			&Elem{3, String, "4"},
		}, 3, 4, 0, true},
	}
	for n, test := range testset {
		encoded, err := test.List.Encode()
		if err != nil {
			t.Fatalf("#%d> encode: %s", n, err)
		}
		val, err := SearchUint64(encoded, test.SearchKey, test.Max)
		if test.ExpectError && err == nil {
			t.Errorf("#%d> expected error but search succeeded", n)
		} else if !test.ExpectError && err != nil {
			t.Errorf("#%d> unexpected error: %s", n, err)
		} else if !test.ExpectError && val != test.ExpectValue {
			t.Errorf("#%d> expected %d but %d found", n, test.ExpectValue, val)
		}
	}
}
