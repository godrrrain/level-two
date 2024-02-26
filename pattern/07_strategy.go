package pattern

import "fmt"

/*
Стратегия является поведенческим паттерном проектирования уровня объекта, который
определяет набор алгоритмов схожих по роду деятельности, помещает их в отдельный класс
и делает подменяемыми.

Применение:
	1. Хотим динамически менять алгоритм решения задачи во время выполнения программы.
	2. Хотим больше гибкости.

Плюсы:
	1. Можем динамически менять алгоритм решения задачи во время выполнения программы.
	2. Легко добавляем новые варианты алгоритмов.

Минусы:
	1. Усложняем код внедрением дополнительных классов.
*/

type StrategySort interface {
	Sort([]int)
}

// сортировка пузырьком
type BubbleSort struct {
}

func (s *BubbleSort) Sort(a []int) {
	fmt.Println("Сортировка пузырьком:", a)
}

// сортировка вставками
type InsertionSort struct {
}

func (s *InsertionSort) Sort(a []int) {
	fmt.Println("Сортировка вставками:", a)
}

// Контекст, использующий стратегию сортировки
type SortContext struct {
	strategy StrategySort
}

func (c *SortContext) ChangeAlgorithm(a StrategySort) {
	c.strategy = a
}

func (c *SortContext) Sort(s []int) {
	c.strategy.Sort(s)
}
