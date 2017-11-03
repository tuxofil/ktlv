package ktlv

// Check if elements are equal or not.
// Used in tests.
func (e1 *Elem) Equals(e2 *Elem) bool {
	if e1.Key != e2.Key || e1.FType != e2.FType {
		return false
	}
	switch e1.FType {
	case Bitmap:
		v1 := e1.Value.([]bool)
		v2 := e2.Value.([]bool)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_String:
		v1 := e1.Value.([]string)
		v2 := e2.Value.([]string)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Double:
		v1 := e1.Value.([]float64)
		v2 := e2.Value.([]float64)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Int8:
		v1 := e1.Value.([]int8)
		v2 := e2.Value.([]int8)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Int16:
		v1 := e1.Value.([]int16)
		v2 := e2.Value.([]int16)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Int24:
		v1 := e1.Value.([]int32)
		v2 := e2.Value.([]int32)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Int32:
		v1 := e1.Value.([]int32)
		v2 := e2.Value.([]int32)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Int64:
		v1 := e1.Value.([]int64)
		v2 := e2.Value.([]int64)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Uint8:
		v1 := e1.Value.([]uint8)
		v2 := e2.Value.([]uint8)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Uint16:
		v1 := e1.Value.([]uint16)
		v2 := e2.Value.([]uint16)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Uint24:
		v1 := e1.Value.([]uint32)
		v2 := e2.Value.([]uint32)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Uint32:
		v1 := e1.Value.([]uint32)
		v2 := e2.Value.([]uint32)
		if len(v1) != len(v2) {
			return false
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != v2[i] {
				return false
			}
		}
	case List_of_Uint64:
		v1 := e1.Value.([]uint64)
		v2 := e2.Value.([]uint64)
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
