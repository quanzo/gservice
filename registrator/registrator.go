/* Класс для регистрации работающих горутин

*/

package registrator

import (
	"sync"
)

type GoRegistrator struct {
	IRegistrator
	gorutines map[int]chan byte
	lock      sync.Mutex
}

func NewGo() *GoRegistrator {
	this := new(GoRegistrator)
	this.gorutines = make(map[int]chan byte)
	return this
}

func (this *GoRegistrator) Register() (int, chan byte) {
	this.lock.Lock()
	defer this.lock.Unlock()
	i := len(this.gorutines)
	this.gorutines[i] = make(chan byte)
	return i, this.gorutines[i]
}
func (this *GoRegistrator) closeChan(id int) {
	defer func() {
		recover()
	}()
	close(this.gorutines[id])
}
func (this *GoRegistrator) UnRegister(id int) {
	this.lock.Lock()
	defer this.lock.Unlock()
	_, exist := this.gorutines[id]
	if exist {
		this.closeChan(id)
		delete(this.gorutines, id)
	}
}
func (this *GoRegistrator) Count() int {
	return len(this.gorutines)
}
func (this *GoRegistrator) GetChan(id int) chan byte {
	this.lock.Lock()
	defer this.lock.Unlock()
	c, exist := this.gorutines[id]
	if exist {
		return c
	}
	return nil
}
