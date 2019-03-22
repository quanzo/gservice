package stack

/* Последним пришел - первым ушел
Элементы хранятся в виде reflect.Value
*/

import ()

type RLifo struct {
	*RFifo
}

func NewRLifo() *RLifo {
	this := new(RLifo)
	this.RFifo = NewRFifo()
	return this
}

func (this *RLifo) Pop() interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.counter--
	return this.popIndex(this.counter)
}
