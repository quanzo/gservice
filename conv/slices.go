package conv

// Convert []int in []rune.
func ConvSlicesIntRune(input []int) []rune {
	if len(input) > 0 {
		var (
			i, v int
			res  []rune
		)
		res = make([]rune, len(input))
		for i, v = range input {
			res[i] = rune(v)
		}
		return res
	} else {
		return make([]rune, 0)
	}
}

// Convert []rune in []int.
func ConvSlicesRuneInt(input []rune) []int {
	if len(input) > 0 {
		var (
			i   int
			v   rune
			res []int
		)
		res = make([]int, len(input))
		for i, v = range input {
			res[i] = int(v)
		}
		return res
	} else {
		return make([]int, 0)
	}
}

/*
// Convert []int in []rune.
func ConvSlicesIntRune(input []int) []rune {
	return *(*[]rune)(unsafe.Pointer(&input))
}

// Convert []rune in []int.
func ConvSlicesRuneInt(input []rune) []int {
	return *(*[]int)(unsafe.Pointer(&input))
}
*/
