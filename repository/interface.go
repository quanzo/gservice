package repository

/* Интерфейс хранилища элементов (по строковому ключу)

*/
type IRepository interface {
	Set(key string, value interface{})          // установить значение
	Get(key string) interface{}                 // получить значение
	Delete(key string)                          // удалить
	Exist(key string) bool                      // проверить на наличие
	Walk(func(index string, value interface{})) // применить функцию ко всем значениям
	Count() int                                 // количество элементов
}
