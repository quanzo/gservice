package bufferint

import (
	"testing"
)

func TestAppend(t *testing.T) {
	buff := New(5, 5)
	buff.Append(1, 2, 3, 4, 5, 6)
	if !buff.Equal(buff.GetCopy(), []int{1, 2, 3, 4, 5, 6}) {
		t.Error("Append result not equal etalon")
	}
}

func TestSubstr(t *testing.T) {
	buff := New(5, 5)
	buff.Append(1, 2, 3, 4, 5, 6)

	if !buff.Equal(buff.Substr(0, 3), []int{1, 2, 3}) {
		t.Error("Substr error [0:3]")
	}
	t.Log(buff.Substr(0, 3))
	if !buff.Equal(buff.Substr(3, 20), []int{4, 5, 6}) {
		t.Error("Substr error [3:20]")
	}
	t.Log(buff.Substr(3, 20))
}
func TestDel(t *testing.T) {
	buff := New(5, 5)
	buff.Append(1, 2, 3, 4, 5, 6)
	buff.Delete(3, 10)
	if !buff.Equal(buff.GetCopy(), []int{1, 2, 3}) {
		t.Error("Delete error [0:3]")
	}
	t.Log(buff.GetCopy())
	buff.Delete(10, 20)
	if !buff.Equal(buff.GetCopy(), []int{1, 2, 3}) {
		t.Error("Delete error [10:20]")
	}
}
func TestWalk(t *testing.T) {
	buff := New(10, 5)
	buff.Append(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	buff.Walk(0, 10, func(i int, v *int) {
		*v++
	})
	if !buff.Equal(buff.GetCopy(), []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}) {
		t.Error("Error walk function. Err increment buffer values")
	}
}

func TestFilter(t *testing.T) {
	buff := New(10, 5)
	buff.Append(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	buff.Filter(func(i, v int) bool { // only even
		if v%2 > 0 {
			return false
		}
		return true
	})

	if !buff.Equal(buff.GetCopy(), []int{2, 4, 6, 8, 10}) {
		t.Error("Error filter function.")
	}
}

func TestPop(t *testing.T) {
	buff := New(10, 1)
	buff.Append(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	t1 := buff.Pop(0, 1)
	if !buff.Equal(t1, []int{1}) || !buff.Equal(buff.GetCopy(), []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}) {
		t.Error("Error pop first element", t1, buff.GetCopy())
	}
	buffLength := buff.Length()
	t2 := buff.Pop(buffLength-1, 1)
	if !buff.Equal(t2, []int{20}) || !buff.Equal(buff.GetCopy(), []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}) {
		t.Error("Error pop end element")
	}
	t3 := buff.Pop(3, 4)
	if !buff.Equal(t3, []int{5, 6, 7, 8}) || !buff.Equal(buff.GetCopy(), []int{2, 3, 4, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}) {
		t.Error("Error pop center elements")
	}
	t4 := buff.Pop(buff.Length(), 10)
	if !buff.Equal(t4, []int{}) || !buff.Equal(buff.GetCopy(), []int{2, 3, 4, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}) {
		t.Error("Error pop out of range elements")
	}

}
