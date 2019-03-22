package stack

/* Последним пришел - первым ушел
Данные хранятся в виде interface{}
*/

import (
//"sync"
)

type Lifo struct {
	*Fifo
}

func NewLifo() *Lifo {
	this := new(Lifo)
	this.Fifo = NewFifo()
	//this.queue = make(map[uint]interface{})
	//this.counter = 0
	return this
}

func (this *Lifo) Pop() interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.counter--
	return this.popIndex(this.counter)
}
