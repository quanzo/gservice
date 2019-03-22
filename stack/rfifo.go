package stack

/* Первым пришел - первым ушел
Элементы хранятся в виде reflect.Value
*/

import (
	//"fmt"
	"reflect"
	"sync"
)

func NewRFifo() *RFifo {
	this := new(RFifo)
	this.queue = make(map[uint]reflect.Value)
	this.counter = 0
	return this
}

type RFifo struct {
	//Stack
	queue   map[uint]reflect.Value
	lock    sync.Mutex
	counter uint
}

func (this *RFifo) Push(d interface{}) *RFifo {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.queue[this.counter] = reflect.ValueOf(d)
	this.counter++
	return this
}

func (this *RFifo) popIndex(index2pop uint) interface{} {
	if index2pop >= 0 {
		v, ok := this.queue[index2pop]
		if ok {
			delete(this.queue, index2pop)
			return v.Interface()
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (this *RFifo) Pop() interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()
	return this.popIndex((this.counter - uint(len(this.queue))))
}

func (this *RFifo) Len() int {
	return len(this.queue)
}

func (this *RFifo) InStack(f interface{}) bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	f_val := reflect.ValueOf(f)
	for _, value := range this.queue {
		if f_val == value {
			return true
		}
	}
	return false
}
