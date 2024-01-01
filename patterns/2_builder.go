package patterns

/*
Паттерн Builder определяет поэтапное создание экземпляра сложного объекта.
Такой подход уменьшает вероятность опечатки и повышает читаемость кода, но
увеличивает объём бойлерплейта
*/

type Box struct {
	length int
	width  int
	height int
}

type Builder interface {
	setLength()
	setWidth()
	setHeight()
	getBox() Box
}

type BoxBuilder struct {
	length int
	width  int
	height int
}

func (b *BoxBuilder) setLength(length int) {
	b.length = length
}

func (b *BoxBuilder) setWidth(width int) {
	b.width = width
}

func (b *BoxBuilder) setHeight(height int) {
	b.height = height
}

func (b *BoxBuilder) getBox() Box {
	return Box{
		length: b.length,
		width:  b.width,
		height: b.height,
	}
}
