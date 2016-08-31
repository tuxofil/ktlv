package ktlv

import (
	"testing"
)

func encdec(t *testing.T, data0 Decoded) {
	bytes := Enc(data0)
	data1 := Dec(bytes)
	if len(data0) != len(data1) {
		t.Fatalf(
			"encdec: origin data len is %v but encoded-decoded data len is %v",
			len(data0), len(data1))
	}
	for i := 0; i < len(data0); i++ {
		e0 := data0[i]
		e1 := data1[i]
		if !e0.Equals(e1) {
			t.Fatalf("elems differ: %v and %v", e0, e1)
		}
	}
}

func TestBool(t *testing.T) {
	encdec(t, Decoded{
		&Elem{1, Bool, true},
		&Elem{2, Bool, false}})
}

func TestUint8(t *testing.T) {
	encdec(t, Decoded{
		&Elem{1, Uint8, Min_Uint8},
		&Elem{2, Uint8, Max_Uint8}})
}

func TestUint16(t *testing.T) {
	encdec(t, Decoded{
		&Elem{1, Uint16, Min_Uint16},
		&Elem{2, Uint16, Max_Uint16}})
}

func TestUint24(t *testing.T) {
	encdec(t, Decoded{
		&Elem{1, Uint24, Min_Uint24},
		&Elem{2, Uint24, Max_Uint24}})
}

func TestUint32(t *testing.T) {
	encdec(t, Decoded{
		&Elem{1, Uint32, Min_Uint32},
		&Elem{2, Uint32, Max_Uint32}})
}

func TestUint64(t *testing.T) {
	encdec(t, Decoded{
		&Elem{1, Uint64, Min_Uint64},
		&Elem{2, Uint64, Max_Uint64}})
}

func TestInt8(t *testing.T) {
	encdec(t, Decoded{
		&Elem{1, Int8, Min_Int8},
		&Elem{2, Int8, Max_Int8}})
}

func TestInt16(t *testing.T) {
	encdec(t, Decoded{
		&Elem{1, Int16, Min_Int16},
		&Elem{2, Int16, Max_Int16}})
}

func TestInt24(t *testing.T) {
	encdec(t, Decoded{
		&Elem{1, Int24, Min_Int24},
		&Elem{4, Int24, int32(-1000)},
		&Elem{2, Int24, int32(-1)},
		&Elem{3, Int24, int32(0)},
		&Elem{4, Int24, int32(1)},
		&Elem{5, Int24, int32(1000)},
		&Elem{6, Int24, Max_Int24}})
}

func TestInt32(t *testing.T) {
	encdec(t, Decoded{
		&Elem{1, Int32, Min_Int32},
		&Elem{2, Int32, Max_Int32}})
}

func TestInt64(t *testing.T) {
	encdec(t, Decoded{
		&Elem{1, Int64, Min_Int64},
		&Elem{2, Int64, Max_Int64}})
}

func TestDouble(t *testing.T) {
	encdec(t, Decoded{
		&Elem{1, Double, float64(0.0)},
		&Elem{2, Double, float64(-0.0)},
		&Elem{3, Double, float64(-1.0)},
		&Elem{4, Double, float64(1.0)},
		&Elem{5, Double, float64(3.1415927)}})
}

func TestString(t *testing.T) {
	encdec(t, Decoded{
		&Elem{1, String, ""},
		&Elem{2, String, "a"},
		&Elem{3, String, "abc"}})
}

func TestBitmap(t *testing.T) {
	encdec(t, Decoded{
		&Elem{1, Bitmap, []bool{}},
		&Elem{2, Bitmap, []bool{false}},
		&Elem{3, Bitmap, []bool{true}},
		&Elem{4, Bitmap, []bool{true, true}},
		&Elem{5, Bitmap, []bool{true, true, false}},
		&Elem{6, Bitmap, []bool{true, true, false, true, true, false, false, true, true}}})
}

func TestListOfString(t *testing.T) {
	encdec(t, Decoded{
		&Elem{1, List_of_String, []string{}},
		&Elem{2, List_of_String, []string{""}},
		&Elem{3, List_of_String, []string{"", ""}},
		&Elem{4, List_of_String, []string{"a", "b"}}})
}

func TestListOfUint8(t *testing.T) {
	encdec(t, Decoded{
		&Elem{1, List_of_Uint8, []uint8{}},
		&Elem{2, List_of_Uint8, []uint8{0}},
		&Elem{3, List_of_Uint8, []uint8{1}},
		&Elem{4, List_of_Uint8, []uint8{1, 1}},
		&Elem{5, List_of_Uint8, []uint8{1, 2, 3}},
		&Elem{6, List_of_Uint8, []uint8{Min_Uint8, 0, Max_Uint8}}})
}

func TestListOfUint16(t *testing.T) {
	encdec(t, Decoded{
		&Elem{1, List_of_Uint16, []uint16{}},
		&Elem{2, List_of_Uint16, []uint16{0}},
		&Elem{3, List_of_Uint16, []uint16{1}},
		&Elem{4, List_of_Uint16, []uint16{1, 1}},
		&Elem{5, List_of_Uint16, []uint16{1, 2, 3}},
		&Elem{6, List_of_Uint16, []uint16{Min_Uint16, 0, Max_Uint16}}})
}

func TestListOfUint24(t *testing.T) {
	encdec(t, Decoded{
		&Elem{1, List_of_Uint24, []uint32{}},
		&Elem{2, List_of_Uint24, []uint32{0}},
		&Elem{3, List_of_Uint24, []uint32{1}},
		&Elem{4, List_of_Uint24, []uint32{1, 1}},
		&Elem{5, List_of_Uint24, []uint32{1, 2, 3}},
		&Elem{6, List_of_Uint24, []uint32{Min_Uint24, 0, Max_Uint24}}})
}

func TestListOfUint32(t *testing.T) {
	encdec(t, Decoded{
		&Elem{1, List_of_Uint32, []uint32{}},
		&Elem{2, List_of_Uint32, []uint32{0}},
		&Elem{3, List_of_Uint32, []uint32{1}},
		&Elem{4, List_of_Uint32, []uint32{1, 1}},
		&Elem{5, List_of_Uint32, []uint32{1, 2, 3}},
		&Elem{6, List_of_Uint32, []uint32{Min_Uint32, 0, Max_Uint32}}})
}

func TestListOfUint64(t *testing.T) {
	encdec(t, Decoded{
		&Elem{1, List_of_Uint64, []uint64{}},
		&Elem{2, List_of_Uint64, []uint64{0}},
		&Elem{3, List_of_Uint64, []uint64{1}},
		&Elem{4, List_of_Uint64, []uint64{1, 1}},
		&Elem{5, List_of_Uint64, []uint64{1, 2, 3}},
		&Elem{6, List_of_Uint64, []uint64{Min_Uint64, 0, Max_Uint64}}})
}

func TestListOfInt8(t *testing.T) {
	encdec(t, Decoded{
		&Elem{1, List_of_Int8, []int8{}},
		&Elem{2, List_of_Int8, []int8{0}},
		&Elem{3, List_of_Int8, []int8{1}},
		&Elem{4, List_of_Int8, []int8{1, 1}},
		&Elem{5, List_of_Int8, []int8{1, -2, 3}},
		&Elem{6, List_of_Int8, []int8{Min_Int8, 0, Max_Int8}}})
}

func TestListOfInt16(t *testing.T) {
	encdec(t, Decoded{
		&Elem{1, List_of_Int16, []int16{}},
		&Elem{2, List_of_Int16, []int16{0}},
		&Elem{3, List_of_Int16, []int16{1}},
		&Elem{4, List_of_Int16, []int16{1, 1}},
		&Elem{5, List_of_Int16, []int16{1, -2, 3}},
		&Elem{6, List_of_Int16, []int16{Min_Int16, 0, Max_Int16}}})
}

func TestListOfInt24(t *testing.T) {
	encdec(t, Decoded{
		&Elem{1, List_of_Int24, []int32{}},
		&Elem{2, List_of_Int24, []int32{0}},
		&Elem{3, List_of_Int24, []int32{1}},
		&Elem{4, List_of_Int24, []int32{1, 1}},
		&Elem{5, List_of_Int24, []int32{1, -2, 3}},
		&Elem{6, List_of_Int24, []int32{Min_Int24, 0, Max_Int24}}})
}

func TestListOfInt32(t *testing.T) {
	encdec(t, Decoded{
		&Elem{1, List_of_Int32, []int32{}},
		&Elem{2, List_of_Int32, []int32{0}},
		&Elem{3, List_of_Int32, []int32{1}},
		&Elem{4, List_of_Int32, []int32{1, 1}},
		&Elem{5, List_of_Int32, []int32{1, -2, 3}},
		&Elem{6, List_of_Int32, []int32{Min_Int32, 0, Max_Int32}}})
}

func TestListOfInt64(t *testing.T) {
	encdec(t, Decoded{
		&Elem{1, List_of_Int64, []int64{}},
		&Elem{2, List_of_Int64, []int64{0}},
		&Elem{3, List_of_Int64, []int64{1}},
		&Elem{4, List_of_Int64, []int64{1, 1}},
		&Elem{5, List_of_Int64, []int64{1, -2, 3}},
		&Elem{6, List_of_Int64, []int64{Min_Int64, 0, Max_Int64}}})
}
