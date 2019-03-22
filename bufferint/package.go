package bufferint

const (
	BUFFER_SIZE int = 50 // размер буфера по умолчанию
)

// Создать буфер размером size и размером увеличения буфера равным addSpace.
func New(size int, addSpace int) *BufferInt {
	this := new(BufferInt)
	this.init(size)
	if addSpace > 0 {
		this.addSpace = addSpace
	}
	return this
}
