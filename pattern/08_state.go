package pattern

import "fmt"

/*
Состояние является поведенческим паттерном проектирования уровня объекта, который
позволяет объектам менять поведение в зависимости от своего состояния.

Применение:
	1. Хотим менять поведение объектов в зависимости от своего состояния,
	когда таких состояний может быть много.

Плюсы:
	1. Избавляемся от множества больших условных операторов.
	2. Разграничиваем код, связанный с конкретными состояниями объекта.

Минусы:
	1. Можем неоправданно усложнять код при небольшом количестве состояний.
*/

// Интерфейс состояния
type State interface {
	Print()
}

// Первое состояние
type OnState struct{}

func (state OnState) Print() {
	fmt.Println("Цветной режим")
}

// Второе состояние
type OffState struct{}

func (state OffState) Print() {
	fmt.Println("Черно-белый режим")
}

// Контекст, использующий состояние
type Printer struct {
	state State
}

func (printer *Printer) SetState(state State) {
	printer.state = state
}

func (printer *Printer) Print() {
	printer.state.Print()
}
