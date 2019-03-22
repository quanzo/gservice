package stack

/* Первым пришел - первым ушел
Данные хранятся в виде interface{}
*/

import (
	//"fmt"
	"sync"
)

func NewFifo() *Fifo {
	this := new(Fifo)
	this.queue = make(map[uint]interface{})
	this.counter = 0
	return this
}

type Fifo struct {
	//Stack
	queue   map[uint]interface{}
	lock    sync.Mutex
	counter uint
}

func (this *Fifo) Push(d interface{}) *Fifo {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.queue[this.counter] = d
	this.counter++
	//fmt.Println(this.counter)
	return this
}

func (this *Fifo) popIndex(index2pop uint) interface{} {
	if index2pop >= 0 {
		v, ok := this.queue[index2pop]
		if ok {
			delete(this.queue, index2pop)
			return v
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (this *Fifo) Pop() interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()
	return this.popIndex((this.counter - uint(len(this.queue))))
}

func (this *Fifo) Len() int {
	return len(this.queue)
}
