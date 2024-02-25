package pattern

import "fmt"

/*
Посетитель является поведенческим паттерном проектирования уровня объекта, который
позволяет представить запрос в виде объекта.

Применение:
	1. Хотим ставить запросы в очередь.
	2. Хотим иметь возможность отменять и возобновлять запросы.

Плюсы:
	1. Можем ставить в очередь, отменять, возобновлять запросы.
	2. Убираем прямую зависимость между объектами, которые вызывают и непосредственно выполняют операции.

Минусы:
	1. Усложняем код внедрением дополнительных классов.
*/

type Command interface {
	Execute()
}

// Конкретная команда
type ConcreteCommand struct {
	Receiver *Receiver
}

func (c *ConcreteCommand) Execute() {
	c.Receiver.Action()
}

// Получатель команды
type Receiver struct{}

func (r *Receiver) Action() {
	fmt.Println("Receiver: выполнение действия")
}

// Инвокер
type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}
