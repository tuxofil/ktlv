package ktlv

import "errors"

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

var t2s = map[uint8]string{
	Bool:           "Bool",
	Uint8:          "Uint8",
	Uint16:         "Uint16",
	Uint24:         "Uint24",
	Uint32:         "Uint32",
	Uint64:         "Uint64",
	Double:         "Double",
	String:         "String",
	Bitmap:         "Bitmap",
	Int8:           "Int8",
	Int16:          "Int16",
	Int24:          "Int24",
	Int32:          "Int32",
	Int64:          "Int64",
	List_of_String: "List_of_String",
	List_of_Uint8:  "List_of_Uint8",
	List_of_Uint16: "List_of_Uint16",
	List_of_Uint24: "List_of_Uint24",
	List_of_Uint32: "List_of_Uint32",
	List_of_Uint64: "List_of_Uint64",
	List_of_Double: "List_of_Double",
	List_of_Int8:   "List_of_Int8",
	List_of_Int16:  "List_of_Int16",
	List_of_Int24:  "List_of_Int24",
	List_of_Int32:  "List_of_Int32",
	List_of_Int64:  "List_of_Int64",
}

func FTypeToString(t uint8) string {
	return t2s[t]
}

var (
	ElementNotFound     = errors.New("no such element")
	TypeAssertionFailed = errors.New("unexpected element type")
)
