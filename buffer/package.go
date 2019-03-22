package buffer

func NewEmpty(size int, addSpace int) *Buffer {
	this := new(Buffer)
	this.init(size, addSpace)
	return this
}
