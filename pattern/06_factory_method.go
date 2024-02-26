package pattern

import "fmt"

/*
Фабричный метод является порождающим паттерном проектирования, который определяет общий интерфейс
для создания объектов, позволяя подклассам изменять тип создаваемых объектов.

Применение:
	1. Хотим создавать различные объекты, не указывая конкретных типов объектов.

Плюсы:
	1. Избавляем главный класс от привязки к конкретным типам объектов.
	2. Легко добавляем новые типы объектов в систему.

Минусы:
	1. Усложняем код внедрением дополнительных классов.
*/

// Фабрика
type LoggerFactory interface {
	CreateLogger(action int) Logger
}

// Абстрактный класс логгера
type Logger interface {
	Log(message string)
}

type ConcreteLoggerFactory struct{}

func NewCreator() LoggerFactory {
	return &ConcreteLoggerFactory{}
}

func (f *ConcreteLoggerFactory) CreateLogger(action int) Logger {
	var logger Logger

	switch action {
	case 1:
		logger = &FileLogger{}
	case 2:
		logger = &ConsoleLogger{}
	default:
		fmt.Println("Unknown action")
	}

	return logger
}

// Файловый логгер
type FileLogger struct{}

func (logger *FileLogger) Log(message string) {
	fmt.Println("Запись в файл:", message)
}

// Консольный логгер
type ConsoleLogger struct{}

func (logger *ConsoleLogger) Log(message string) {
	fmt.Println("Вывод в консоль:", message)
}
