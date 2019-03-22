package bufferint

import (
	//	"fmt"
	"sync"

	"github.com/quanzo/gservice/status"
)

type BufferInt struct {
	buffer         []int
	size           int
	addSpace       int
	lock           *sync.Mutex
	modeThreadSafe bool
}

func (this *BufferInt) init(bufferSize int) {
	if bufferSize <= 0 {
		bufferSize = BUFFER_SIZE
	}
	if this.addSpace <= 0 {
		this.addSpace = 10
	}
	this.buffer = make([]int, bufferSize)
	this.size = 0
	this.lock = new(sync.Mutex)
}

//==============================================================================

// Увеличить размер буфера.
func (this *BufferInt) alloc(newSize int) {
	if len(this.buffer) < newSize {
		this.buffer = append(this.buffer, make([]int, newSize-len(this.buffer)+this.addSpace)...)
	}
}

// Заменить часть buffer начиная с start длинной count на sInt
func (this *BufferInt) replace(start int, count int, sInt *[]int) int {
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
	if sInt != nil {
		sLen = len(*sInt)
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
		copy(this.buffer[start:], *sInt)
	}
	return this.size
} // end replace

// Получить срез: начиная с элемента start выдать count элементов.
func (this *BufferInt) substr(start int, count int) []int {
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
func (this *BufferInt) Append(s ...int) {
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

func (this *BufferInt) AppendSlice(s []int) {
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
func (this *BufferInt) AppendBuffer(buff *BufferInt, checked func(index int, value int) bool) int {
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
			checked = func(index int, value int) bool {
				return true
			}
		}
		var (
			i, res_count int
			v            int
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

// Вставить число s перед элементом before_pos. Возвращает новый размер буфера.
func (this *BufferInt) Insert(s int, before_pos int) int {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}
	ss := []int{s}
	return this.replace(before_pos, 0, &ss)
}

// Вставить slice s перед элементом before_pos. Возвращает новый размер буфера.
func (this *BufferInt) InsertSlice(s []int, before_pos int) int {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}
	return this.replace(before_pos, 0, &s)
}

// Удалить count элементов начиная с символа start. Возвращает кол-во удаленных символов.
func (this *BufferInt) Delete(start int, count int) int {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}
	s := this.size
	return s - this.replace(start, count, nil)
} // end func Delete

// Получить срез элементов из буфера: начиная с символа start выдать count символов.
func (this *BufferInt) Substr(start int, count int) []int {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}
	return this.Copy(this.substr(start, count))
}

// Получить один элемент буфера. Будет возвращена ошибка, если элемента с индексом нет в буфере.
func (this *BufferInt) One(i int) (int, error) {
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

// Выбрать и вернуть из буфера цепочку данных. Возвращаемая цепочка из буфера будет удалена.
func (this *BufferInt) Pop(start int, count int) []int {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}
	res := this.Copy(this.substr(start, count))
	_ = this.replace(start, count, nil)
	return res
}

// Поменять два элемента местами. Для возможности использования пакета sort.
func (this *BufferInt) Swap(i, j int) {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}
	if i < this.size && j < this.size && i >= 0 && j >= 0 && i != j {
		this.buffer[i], this.buffer[j] = this.buffer[j], this.buffer[i]
	}
} // end swap

// Сравнить элементы с индексами i и j. Возвражает true если i < j. Для использования пакета sort.
func (this *BufferInt) Less(i, j int) bool {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}
	if i < this.size && j < this.size && i >= 0 && j >= 0 && i != j {
		return this.buffer[i] < this.buffer[j]
	}
	return false
}

// Сравнивает два среза.
func (this *BufferInt) Equal(q, w []int) bool {
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
func (this *BufferInt) Copy(q []int) []int {
	if len(q) > 0 {
		res := make([]int, len(q))
		copy(res, q)
		return res
	} else {
		return []int{}
	}
}

// Очистить буфер.
func (this *BufferInt) Empty() {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}
	this.size = 0
}

// Выполнить функцию к каждому элементу буфера, начиная с позиции start.
func (this *BufferInt) Walk(start int, count int, f func(index int, value *int)) {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}
	if start < this.size && f != nil {
		var (
			i int
			v *int
		)
		for i = start; i < start+count && i < this.size; i++ {
			v = &this.buffer[i]
			f(i, v)
		}
	}
}

// Фильтрация буфера. Если функция filter возвращает false, то элемент удаляется из буфера.
func (this *BufferInt) Filter(filter func(index int, value int) bool) {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}
	if filter != nil {
		var (
			i, v, delCount int
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

}

//==============================================================================

// Длинна буфера.
func (this *BufferInt) Length() int {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}
	return this.size
}

// Вернуть значения буфера.
func (this *BufferInt) GetCopy() []int {
	if this.modeThreadSafe {
		this.lock.Lock()
		defer this.lock.Unlock()
	}
	res := make([]int, this.size)
	copy(res, this.buffer[0:this.size])
	return res
}

// Установить многопоточный режим.
func (this *BufferInt) SetModeThreadSafe(m bool) {
	this.modeThreadSafe = m
}

// Вернуть многопоточный режим.
func (this *BufferInt) GetModeThreadSafe() bool {
	return this.modeThreadSafe
}
