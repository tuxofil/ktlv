package main

import (
	"ktlv"
	"log"
	"os"
)

func main() {
	fileinfo, err := os.Stat("object.bin")
	if err != nil {
		log.Fatalf("unable to get file size: %v", err)
	}
	file, err := os.Open("object.bin")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}
	b := make([]byte, fileinfo.Size())
	n, err := file.Read(b)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	if int64(n) != fileinfo.Size() {
		log.Fatalf("only %v bytes read but %v expected", n, fileinfo.Size())
	}
	data := ktlv.Dec(b)

	expected := ktlv.Decoded{
		&ktlv.Elem{1, ktlv.Bool, true},
		&ktlv.Elem{2, ktlv.Uint8, uint8(2)},
		&ktlv.Elem{3, ktlv.Uint16, uint16(3)},
		&ktlv.Elem{4, ktlv.Uint24, uint32(4)},
		&ktlv.Elem{5, ktlv.Uint32, uint32(5)},
		&ktlv.Elem{6, ktlv.Uint64, uint64(6)},
		&ktlv.Elem{7, ktlv.Double, 3.1415927},
		&ktlv.Elem{8, ktlv.String, "hello"},
		&ktlv.Elem{9, ktlv.Bitmap, []bool{true, true, false, false, true, false, true, true, true, true}},
		&ktlv.Elem{10, ktlv.List_of_String, []string{"hello", "world", "!"}},
		&ktlv.Elem{11, ktlv.List_of_Uint8, []uint8{ktlv.Min_Uint8, ktlv.Max_Uint8,
			ktlv.Max_Uint8, ktlv.Min_Uint8}},
		&ktlv.Elem{12, ktlv.List_of_Uint16, []uint16{ktlv.Min_Uint16, ktlv.Max_Uint16,
			ktlv.Max_Uint16, ktlv.Min_Uint16}},
		&ktlv.Elem{13, ktlv.List_of_Uint24, []uint32{ktlv.Min_Uint24, ktlv.Max_Uint24,
			ktlv.Max_Uint24, ktlv.Min_Uint24}},
		&ktlv.Elem{14, ktlv.List_of_Uint32, []uint32{ktlv.Min_Uint32, ktlv.Max_Uint32,
			ktlv.Max_Uint32, ktlv.Min_Uint32}},
		&ktlv.Elem{15, ktlv.List_of_Uint64, []uint64{ktlv.Min_Uint64, ktlv.Max_Uint64,
			ktlv.Max_Uint64, ktlv.Min_Uint64}},
		&ktlv.Elem{16, ktlv.List_of_Double, []float64{1.1, -2.2, 3.3}},
		&ktlv.Elem{17, ktlv.Int8, int8(-2)},
		&ktlv.Elem{18, ktlv.Int16, int16(-3)},
		&ktlv.Elem{19, ktlv.Int24, int32(-4)},
		&ktlv.Elem{20, ktlv.Int32, int32(-5)},
		&ktlv.Elem{21, ktlv.Int64, int64(-6)},
		&ktlv.Elem{22, ktlv.List_of_Int8, []int8{0, ktlv.Min_Int8, ktlv.Max_Int8,
			ktlv.Max_Int8, ktlv.Min_Int8}},
		&ktlv.Elem{23, ktlv.List_of_Int16, []int16{0, ktlv.Min_Int16, ktlv.Max_Int16,
			ktlv.Max_Int16, ktlv.Min_Int16}},
		&ktlv.Elem{24, ktlv.List_of_Int24, []int32{0, ktlv.Min_Int24, ktlv.Max_Int24,
			ktlv.Max_Int24, ktlv.Min_Int24}},
		&ktlv.Elem{25, ktlv.List_of_Int32, []int32{0, ktlv.Min_Int32, ktlv.Max_Int32,
			ktlv.Max_Int32, ktlv.Min_Int32}},
		&ktlv.Elem{26, ktlv.List_of_Int64, []int64{0, ktlv.Min_Int64, ktlv.Max_Int64,
			ktlv.Max_Int64, ktlv.Min_Int64}},
	}

	for i := 0; i < len(expected); i++ {
		e0 := expected[i]
		e1 := data[i]
		if !e0.Equals(e1) {
			log.Fatalf("elems differ: expected %v but decode result is %v", e0, e1)
		}
	}
}
