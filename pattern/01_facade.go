package pattern

import "strings"

/*
Фасад является структурным паттерном проектирования, который предоставляет
упрощенный интерфейс к сложной системе.

Применение:
	1. Хотим облегчить взаимодействие со сложной системой простым интерфейсом.
	2. Хотим добавить новый уровень абстракции.

Плюсы:
	1. На выходе простой в использовании интерфейс.
	2. Изолируем пользователей от сложных компонентов.

Минусы:
	1. Есть вероятность превращения в божественный объект.

Примеры использования:
	1. Библиотеки, фреймворки, пакеты, которые реализуют и предоставляют API к использованию.
*/

type DataLoader struct {
}

func (h *DataLoader) Load() string {
	return "Data loaded"
}

type DataChanger struct {
}

func (t *DataChanger) Change() string {
	return "Data changed"
}

type DataSaver struct {
}

func (c *DataSaver) Save() string {
	return "Data saved"
}

// фасад
type Data struct {
	dataLoader  *DataLoader
	dataChanger *DataChanger
	dataSaver   *DataSaver
}

func NewData() *Data {
	return &Data{
		dataLoader:  &DataLoader{},
		dataChanger: &DataChanger{},
		dataSaver:   &DataSaver{},
	}
}

func (d *Data) CreateData() string {
	result := []string{
		d.dataLoader.Load(),
		d.dataChanger.Change(),
		d.dataSaver.Save(),
	}
	return strings.Join(result, "\n")
}
