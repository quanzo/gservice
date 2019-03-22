package repository

import (
	"sync"
)

type Repository struct {
	IRepository
	lock  sync.Mutex
	vault map[string]interface{}
}

func New() *Repository {
	this := new(Repository)
	this.vault = make(map[string]interface{})
	return this
}
func (this *Repository) Set(key string, value interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.vault[key] = value
}
func (this *Repository) Get(key string) interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()
	v, exist := this.vault[key]
	if exist {
		return v
	} else {
		return nil
	}
}
func (this *Repository) Delete(key string) {
	this.lock.Lock()
	defer this.lock.Unlock()
	_, exist := this.vault[key]
	if exist {
		delete(this.vault, key)
	}
}
func (this *Repository) Exist(key string) bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	_, exist := this.vault[key]
	return exist
}
func (this *Repository) Walk(f func(index string, value interface{})) {
	this.lock.Lock()
	defer this.lock.Unlock()
	for i, v := range this.vault {
		f(i, v)
	}
}
func (this *Repository) Count() int {
	this.lock.Lock()
	defer this.lock.Unlock()
	return len(this.vault)
}
