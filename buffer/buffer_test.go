package buffer

import (
	"testing"

	"github.com/quanzo/gservice/bufferint"
	"github.com/quanzo/bufferstring"
)

type TestingInsert struct {
	name   string
	pos    int
	insert []interface{}
	source []interface{}
	dest   []interface{}
}

type TestingDelete struct {
	name   string
	pos    int
	count  int
	source []interface{}
	dest   []interface{}
}

type TestingSubstr struct {
	source []interface{}
	start  int
	count  int
	dest   []interface{}
}

type TestingFind struct {
	source  []interface{}
	needle  *[]interface{}
	start   int
	reverse bool
	result  int
}

var tInsert []TestingInsert = []TestingInsert{
	TestingInsert{"In begin", 0, []interface{}{10, 20, 30}, []interface{}{"Строка", 0.4, 1}, []interface{}{10, 20, 30, "Строка", 0.4, 1}},
	TestingInsert{"In end", 3, []interface{}{10, 20, 30}, []interface{}{"Строка", 0.4, 1}, []interface{}{"Строка", 0.4, 1, 10, 20, 30}},
	TestingInsert{"Center", 1, []interface{}{10, 20, 30}, []interface{}{"Строка", 0.4, 1}, []interface{}{"Строка", 10, 20, 30, 0.4, 1}},
}

var tDel []TestingDelete = []TestingDelete{
	TestingDelete{"Del first element.", 0, 1, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}, []interface{}{2, 3, 4, 5, 6, 7, 8, 9}},
	TestingDelete{"Del first 3 element.", 0, 3, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}, []interface{}{4, 5, 6, 7, 8, 9}},
	TestingDelete{"Del center 3 element.", 3, 3, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}, []interface{}{1, 2, 3, 7, 8, 9}},
}

var tSubstr []TestingSubstr = []TestingSubstr{
	TestingSubstr{[]interface{}{1, 2, 3, 4, "string#1", "string#2"}, 0, 100, []interface{}{1, 2, 3, 4, "string#1", "string#2"}},
	TestingSubstr{[]interface{}{1, 2, 3, 4, "string#1", "string#2"}, 1, 2, []interface{}{2, 3}},
	TestingSubstr{[]interface{}{1, 2, 3, 4, "string#1", "string#2"}, 5, 100, []interface{}{"string#2"}},
	TestingSubstr{[]interface{}{1, 2, 3, 4, "string#1", "string#2"}, 3, 3, []interface{}{4, "string#1", "string#2"}},
	TestingSubstr{[]interface{}{1, 2, 3, 4, "string#1", "string#2"}, 2, 1, []interface{}{3}},
}

var tFind []TestingFind = []TestingFind{
	TestingFind{[]interface{}{10, 23, 33, 45, "str", 23, 80, "str", 65, 10}, &[]interface{}{23}, 0, false, 1},
	TestingFind{[]interface{}{10, 23, 33, 45, "str", 23, 80, "str", 65, 10}, &[]interface{}{23}, 4, false, 5},
	TestingFind{[]interface{}{10, 23, 33, 45, "str", 23, 80, "str", 65, 10}, &[]interface{}{23}, -1, true, 5},
	TestingFind{[]interface{}{10, 23, 33, 45, "str", 23, 80, "str", 65, 10}, &[]interface{}{"str"}, -1, false, 4},
	TestingFind{[]interface{}{10, 23, 33, 45, "str", 23, 80, "str", 65, 10}, &[]interface{}{"str"}, 6, true, 4},

	TestingFind{[]interface{}{10, 23, 33, 45, "str", 23, 80, "str", 65, 10}, &[]interface{}{"str", 23}, 6, true, 4},
	TestingFind{[]interface{}{10, 23, 33, 45, "str", 23, 80, "str", 65, 10}, &[]interface{}{"str", 65}, 6, true, -1},
	TestingFind{[]interface{}{10, 23, 33, 45, "str", 23, 80, "str", 65, 10}, &[]interface{}{"str", 65, 10}, -1, false, 7},
}

func TestAppend(t *testing.T) {
	buff := NewEmpty(0, 1)
	buff.Append(10)
	buff.Append("Строка", 10.1)
	if buff.Length() != 3 {
		t.Error("Error. Length does not meet the standard.")
	}
	if !buff.Equal(buff.GetCopy(), []interface{}{10, "Строка", 10.1}) {
		t.Error("Error append data to buffer.")
	}
}

func TestInsert(t *testing.T) {
	buff := NewEmpty(0, 1)
	for i, test := range tInsert {
		buff.Empty()
		buff.AppendSlice(test.source)
		buff.InsertSlice(test.insert, test.pos)
		if !buff.Equal(buff.GetCopy(), test.dest) {
			t.Error("Error insert #", i)
		}
	}
}

func TestDelete(t *testing.T) {
	buff := NewEmpty(0, 1)
	for i, test := range tDel {
		buff.Empty()
		buff.AppendSlice(test.source)
		buff.Delete(test.pos, test.count)
		if !buff.Equal(buff.GetCopy(), test.dest) {
			t.Error("Error insert #", i)
		}
	}
}

func TestOne(t *testing.T) {
	buff := NewEmpty(0, 1)
	buff.Append(10)
	buff.Append("Строка", 10.1)
	if v, e := buff.one(1); v != "Строка" || e != nil {
		t.Error("Error method one!")
	}
	if v, e := buff.One(1); v != "Строка" || e != nil {
		t.Error("Error method one!")
	}
}

func TestSubstr(t *testing.T) {
	buff := NewEmpty(0, 1)
	for i, test := range tSubstr {
		buff.Empty()
		buff.AppendSlice(test.source)
		if !buff.Equal(buff.substr(test.start, test.count), test.dest) {
			t.Error("Error test #", i)
		}
		if !buff.Equal(buff.Substr(test.start, test.count), test.dest) {
			t.Error("Error test #", i)
		}
	}
}

func TestFind(t *testing.T) {
	buff := NewEmpty(0, 1)
	for i, test := range tFind {
		buff.Empty()
		buff.AppendSlice(test.source)
		if buff.find(test.needle, test.start, test.reverse) != test.result {
			t.Error("Error test #", i, buff.find(test.needle, test.start, test.reverse))
		}
	}
}

func TestWalk(t *testing.T) {
	buff := NewEmpty(0, 1)
	buff_res := NewEmpty(0, 1)
	buff.Append(1, 3, 2, 4, 5, 6, 7, 8, 8, 7, 7, 4, 4, 10)
	buff.Walk(0, buff.Length(), func(index int, v *interface{}) {
		vi := (*v).(int)
		if vi-(vi/2)*2 > 0 {
			buff_res.Append(vi)
		}
	})
	if !buff.Equal(buff_res.GetCopy(), []interface{}{1, 3, 5, 7, 7, 7}) {
		t.Error("Error walk", buff_res.GetCopy())
	}
}

// BENCHMARK

func BenchmarkEqualInterface(b *testing.B) {
	buff := new(Buffer)
	ar_1 := make([]interface{}, 1000)
	ar_2 := make([]interface{}, 1000)
	for i := 0; i < 1000; i++ {
		ar_1[i] = i
		ar_2[i] = i
	}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buff.Equal(ar_1, ar_2)
	}

}
func BenchmarkEqualRune(b *testing.B) {
	buff := new(bufferstring.BufferString)
	ar_1 := make([]rune, 1000)
	ar_2 := make([]rune, 1000)
	for i := 0; i < 1000; i++ {
		ar_1[i] = rune(i)
		ar_2[i] = rune(i)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buff.Equal(ar_1, ar_2)
	}
}

func BenchmarkAppendInterface(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	buff := NewEmpty(100, 100)
	for i := 0; i < b.N; i++ {
		buff.Append(i)
	}
}

func BenchmarkAppendIntbuffer(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	buff := bufferint.New(100, 100)
	for i := 0; i < b.N; i++ {
		buff.Append(i)
	}
}
func BenchmarkAppendIntbufferBigSize(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	buff := bufferint.New(b.N, 100)
	for i := 0; i < b.N; i++ {
		buff.Append(i)
	}
}
func BenchmarkAppendIntbufferBigIncSize(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	buff := bufferint.New(b.N/10, b.N/10)
	for i := 0; i < b.N; i++ {
		buff.Append(i)
	}
}
func BenchmarkAppendSlice(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	buff := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		buff[i] = i
	}
}
func BenchmarkStandartAppendSlice(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	buff := make([]int, 10)
	for i := 0; i < b.N; i++ {
		buff = append(buff, i)
	}
}
func BenchmarkAppendSliceCopy(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	buff := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		copy(buff[i:], []int{i})
	}
}
