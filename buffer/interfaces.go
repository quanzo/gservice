package buffer

type IBuffer interface {
	Append(s ...interface{})                                                        // Добавить элемент в буфер.
	AppendSlice(s []interface{})                                                    // Добавить элементы в буфер.
	AppendBuffer(buff IBuffer, checked func(index int, value interface{}) bool) int // Добавить в буфер значения из другого буфера. Возвращает количество добавленых записей.

	One(i int) (interface{}, error) // Получить один элемент буфера. Будет возвращена ошибка, если элемента с индексом нет в буфере.
	Swap(i1, i2 int)                // Поменять два элемента местами.

	Insert(s interface{}, before_pos int) int        // Вставить s перед элементом before_pos. Возвращает новый размер буфера.
	InsertSlice(s []interface{}, before_pos int) int // Вставить slice s перед элементом before_pos. Возвращает новый размер буфера.

	Find(s interface{}, start int) int               // Найти s начиная с позиции start.
	FindReverse(s interface{}, start int) int        // Найти s начиная с позиции start и до начала строки.
	FindSlice(s []interface{}, start int) int        // Найти slice s начиная с позиции start.
	FindSliceReverse(s []interface{}, start int) int // Найти slice s начиная с позиции start и до начала строки.

	Substr(start int, count int) []interface{} // Получить срез элементов из буфера: начиная с символа start выдать count символов.
	Pop(start int, count int) []interface{}    // Выбрать и вернуть из буфера цепочку данных. Возвращаемая цепочка из буфера будет удалена.

	Delete(start int, count int) interface{} // Удалить count элементов начиная с символа start. Возвращает кол-во удаленных символов.
	Empty()                                  // Очистить буфер.

	Walk(start int, count int, f func(index int, value *interface{})) // Выполнить функцию к каждому элементу буфера, начиная с позиции start.
	Filter(filter func(index int, value interface{}) bool)            // Фильтрация буфера. Если функция filter возвращает false, то элемент удаляется из буфера.

	Length() int // Длинна буфера.
}
