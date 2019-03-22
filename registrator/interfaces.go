package registrator

/* Интерфейс-регистратор

*/
type IRegistrator interface {
	Register() (int, chan byte) // регистрирует горутину, сопоставляет ей номер и канал
	UnRegister(id int)          // убирает горутину из списка
	Count() int                 // кол-во горутин в списке
	GetChan(id int) chan byte   // возвращает канал горутины
}
