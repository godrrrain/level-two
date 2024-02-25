package pattern

import "fmt"

/*
Посетитель является поведенческим паттерном проектирования уровня объекта, который
позволяет обойти набор объектов с разнородными интерфейсами, а также
позволяет добавить новый функционал к существующей структуре, не изменяя саму структуру.

Применение:
	1. Хотим выполнить операции над объектами разных классов с разными интерфейсами.
	2. Хотим выполнить набор связанных операций над группой объектов, независимо от их типов.

Плюсы:
	1. Упрощаем добавление операций, работающих со сложными структурами объектов
	2. Посетитель может накапливать состояние при обходе структуры элементов.

Минусы:
	1. Затрудняем добавление новых классов в систему
*/

type Shape interface {
	accept(Visitor)
}

type Visitor interface {
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
}

type Circle struct{}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

type Rectangle struct{}

func (t *Rectangle) accept(v Visitor) {
	v.visitForRectangle(t)
}

type AreaCalculator struct{}

func (a *AreaCalculator) visitForCircle(s *Circle) {
	fmt.Println("Calculating area for a circle")
}
func (a *AreaCalculator) visitForRectangle(s *Rectangle) {
	fmt.Println("Calculating area for a rectangle")
}

type MiddleCoordinates struct{}

func (a *MiddleCoordinates) visitForCircle(c *Circle) {
	fmt.Println("Calculating middle coordinates for a circle")
}
func (a *MiddleCoordinates) visitForRectangle(t *Rectangle) {
	fmt.Println("Calculating middle coordinates for a rectangle")
}
