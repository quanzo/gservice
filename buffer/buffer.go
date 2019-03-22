package buffer

import (
	"sync"

	_ "github.com/quanzo/gservice/bufferint"
	"github.com/quanzo/gservice/status"
)

type Buffer struct {
	IBuffer
	buffer         []interface{}
	size           int
	addSpace       int
	lock           sync.Mutex
	modeThreadSafe bool
}

func (this *Buffer) init(size int, addSpace int) {
	if addSpace <= 0 {
		addSpace = 1
	}
	if size < 0 {
		size = 0
	}
	this.buffer = make([]interface{}, size, size+addSpace)
	this.addSpace = addSpace
	this.size = 0
}

//******************************************************************************

// Увеличить размер буфера.
func (this *Buffer) alloc(newSize int) {
	if len(this.buffer) < newSize {
		this.buffer = append(this.buffer, make([]interface{}, newSize-len(this.buffer)+this.addSpace)...)
	}
}

// Получить одно значение из буфера с индексом i.
func (this *Buffer) one(i int) (interface{}, error) {
	if i >= 0 && i < this.size {
		return this.buffer[i], nil
	} else {
		return -1, status.New(0, "Index out of range.", true, false)
	}
}

// Заменить часть строки начиная с start длинной count на символы sRunes.
func (this *Buffer) replace(start int, count int, sInterface *[]interface{}) int {
	var (
		sLen int
	)
	if start < 0 {
		start = 0
	}
	if start > this.size {
		start = this.size
	}
	if count < 0 {
		count = 0
	}
	if count > (this.size - start) {
		count = this.size - start
	}
	if sInterface != nil {
		sLen = len(*sInterface)
	} else {
		sLen = 0
	}

	newBuffSize := this.size + sLen - count
	if len(this.buffer) < newBuffSize {
		this.alloc(newBuffSize)
	}
	if sLen != count {
		copy(this.buffer[start+sLen:], this.buffer[start+count:])
	}
	this.size = newBuffSize
	if sLen > 0 {
		copy(this.buffer[start:], *sInterface)
	}
	return this.size
} // end replace

// Поиск цепочки значений в буфере.
func (this *Buffer) find(needle *[]interface{}, start int, back bool) int {
	sNeedle := len(*needle)
	if start == -1 {
		if back {
			start = this.size
		} else {
			start = 0
		}
	}
	if (!back && (start+sNeedle) > this.size) || (back && (start-sNeedle+1) < 0) {
		return -1
	} else {
		var (
			i, i_start, i_end int
		)
		i = start
		for i >= 0 && i <= this.size {
			if back {
				i_start = i - sNeedle
				i_end = i
			} else {
				i_start = i
				i_end = i + sNeedle
			}
			if i_start < 0 || i_end > this.size {
				return -1
			} else {
				if this.Equal(this.buffer[i_start:i_end], *needle) {
					return i_start
				}
			}
			if back {
				i--
			} else {
				i++
			}
		}
		return -1
	}
} // end find

// Получить подстроку: начиная с символа start выдать count символов.
func (this *Buffer) substr(start int, count int) []interface{} {
	if start > this.size || count <= 0 || start < 0 {
		return nil
	} else {
		if start+count > this.size {
			return this.buffer[start:this.size]
		} else {
			return this.buffer[start:(start + count)]
		}
	}
} // end substr

//******************************************************************************

// Добавить элемент в буфер.
func (this *Buffer) Append(s ...interface{}) {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}

	sizeInput := len(s)
	if sizeInput > 0 {
		newSize := this.size + sizeInput
		if newSize > len(this.buffer) {
			this.alloc(newSize + this.addSpace)
		}
		copy(this.buffer[this.size:], s)
		this.size += sizeInput
	}
}

// Добавить элементы в буфер.
func (this *Buffer) AppendSlice(s []interface{}) {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}

	sizeInput := len(s)
	if sizeInput > 0 {
		newSize := this.size + sizeInput
		if newSize > len(this.buffer) {
			this.alloc(newSize + this.addSpace)
		}
		copy(this.buffer[this.size:], s)
		this.size += sizeInput
	}
}

// Добавить в буфер значения из другого буфера. Возвращает количество добавленых записей.
func (this *Buffer) AppendBuffer(buff IBuffer, checked func(index int, value interface{}) bool) int {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}
	inputLength := buff.Length()
	if inputLength > 0 {
		// вычислим размер
		newSize := this.size + inputLength
		if newSize > len(this.buffer) {
			this.alloc(newSize)
		}
		if checked == nil {
			checked = func(index int, value interface{}) bool {
				return true
			}
		}
		var (
			i, res_count int
			v            interface{}
			e            error
		)
		for i = 0; i < inputLength; i++ {
			if v, e = buff.One(i); e == nil {
				if checked(i, v) {
					this.buffer[this.size] = v
					this.size++
					res_count++
				}
			} else {
				break
			}
		}
		return res_count
	} else {
		return 0
	}
}

// Вставить s перед элементом before_pos. Возвращает новый размер буфера.
func (this *Buffer) Insert(s interface{}, before_pos int) int {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}
	ss := []interface{}{s}
	return this.replace(before_pos, 0, &ss)
}

// Вставить slice s перед элементом before_pos. Возвращает новый размер буфера.
func (this *Buffer) InsertSlice(s []interface{}, before_pos int) int {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}
	return this.replace(before_pos, 0, &s)
}

// Найти s начиная с позиции start.
func (this *Buffer) Find(s interface{}, start int) int {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}
	return this.find(&[]interface{}{s}, start, false)
}

// Найти s начиная с позиции start и до начала строки.
func (this *Buffer) FindReverse(s interface{}, start int) int {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}
	return this.find(&[]interface{}{s}, start, true)
}

// Найти slice s начиная с позиции start.
func (this *Buffer) FindSlice(s []interface{}, start int) int {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}
	return this.find(&s, start, false)
}

// Найти slice s начиная с позиции start и до начала строки.
func (this *Buffer) FindSliceReverse(s []interface{}, start int) int {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}
	return this.find(&s, start, true)
}

// Удалить count элементов начиная с символа start. Возвращает кол-во удаленных символов.
func (this *Buffer) Delete(start int, count int) interface{} {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}
	s := this.size
	return s - this.replace(start, count, nil)
} // end func Delete

// Получить срез элементов из буфера: начиная с символа start выдать count символов.
func (this *Buffer) Substr(start int, count int) []interface{} {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}
	return this.Copy(this.substr(start, count))
}

// Получить один элемент буфера. Будет возвращена ошибка, если элемента с индексом нет в буфере.
func (this *Buffer) One(i int) (interface{}, error) {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}
	if i >= 0 && i < this.size {
		return this.buffer[i], nil
	} else {
		return -1, status.New(0, "Index out of range.", true, false)
	}
}

// Поменять два элемента местами.
func (this *Buffer) Swap(i1, i2 int) {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}
	if (i1 >= 0 && i1 < this.size) && (i2 >= 0 && i2 < this.size) && i1 != i2 {
		this.buffer[i1], this.buffer[i2] = this.buffer[i2], this.buffer[i1]
	}
}

// Выбрать и вернуть из буфера цепочку данных. Возвращаемая цепочка из буфера будет удалена.
func (this *Buffer) Pop(start int, count int) []interface{} {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}
	res := this.Copy(this.substr(start, count))
	_ = this.replace(start, count, nil)
	return res
}

// Сравнивает два среза.
func (this *Buffer) Equal(q, w []interface{}) bool {
	if q == nil && w == nil {
		return true
	}
	if q == nil || w == nil || len(q) != len(w) {
		return false
	} else {
		s := len(q)
		for i := 0; i < s; i++ {
			if q[i] != w[i] {
				return false
			}
		}
		return true
	}
} // end Equal

// Копировать slice.
func (this *Buffer) Copy(q []interface{}) []interface{} {
	if len(q) > 0 {
		res := make([]interface{}, len(q))
		copy(res, q)
		return res
	} else {
		return []interface{}{}
	}
}

// Очистить буфер.
func (this *Buffer) Empty() {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}
	this.size = 0
}

// Выполнить функцию к каждому элементу буфера, начиная с позиции start.
func (this *Buffer) Walk(start int, count int, f func(index int, value *interface{})) {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}
	if start < this.size && f != nil {
		var (
			i int
			v *interface{}
		)
		for i = start; i < start+count && i < this.size; i++ {
			v = &this.buffer[i]
			f(i, v)
		}
	}
} // end Walk

// Фильтрация буфера. Если функция filter возвращает false, то элемент удаляется из буфера.
func (this *Buffer) Filter(filter func(index int, value interface{}) bool) {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}
	if filter != nil {
		var (
			i, delCount int
			v           interface{}
		)
		for i = 0; i < this.size; i++ {
			v = this.buffer[i]
			if !filter(i+delCount, v) { // excluding from the buffer
				this.replace(i, 1, nil)
				i--
				delCount++
			}
		}
	}
} // end Filter

//******************************************************************************

// Длинна буфера.
func (this *Buffer) Length() int {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}
	return this.size
}

// Вернуть значения буфера.
func (this *Buffer) GetCopy() []interface{} {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}
	res := make([]interface{}, this.size)
	copy(res, this.buffer[0:this.size])
	return res
}

// Установить многопоточный режим.
func (this *Buffer) SetModeThreadSafe(m bool) {
	this.modeThreadSafe = m
}

// Вернуть многопоточный режим.
func (this *Buffer) GetModeThreadSafe() bool {
	return this.modeThreadSafe
}
