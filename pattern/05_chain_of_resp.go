package pattern

/*
Цепочка обязанностей является поведенческим паттерном проектирования уровня объекта, который
позволяет передавать запросы последовательно по цепочке обработчиков.
Каждый обработчик сам решает, стоит ли передавать запрос дальше по цепочке.

Применение:
	1. Хотим обрабатывать разнообразные запросы, заранее не зная, какие обработчики понадобятся.
	2. Хотим динамически задавать набор объектов, способных обработать запрос.

Плюсы:
	1. Уменьшаем связанность между отправителем запроса и получателем.
	2. Настраиваем обработку запросов более гибко.
	2. Легко добавляем новые обработчики в систему.

Минусы:
	1. Запрос может остаться необработанным.
*/

type Handler interface {
	Handle(message int) string
	SetNext(Handler) Handler
}

// handler "A".
type ConcreteHandlerA struct {
	next Handler
}

func (h *ConcreteHandlerA) SendRequest(message int) (result string) {
	if message == 1 {
		result = "Im handler 1"
	} else if h.next != nil {
		result = h.next.Handle(message)
	}
	return
}

func (h *ConcreteHandlerA) SetNext(next Handler) Handler {
	h.next = next
	return next
}

// handler "B"
type ConcreteHandlerB struct {
	next Handler
}

func (h *ConcreteHandlerB) SendRequest(message int) (result string) {
	if message == 2 {
		result = "Im handler 2"
	} else if h.next != nil {
		result = h.next.Handle(message)
	}
	return
}

func (h *ConcreteHandlerB) SetNext(next Handler) Handler {
	h.next = next
	return next
}

// handler "C"
type ConcreteHandlerC struct {
	next Handler
}

func (h *ConcreteHandlerC) SendRequest(message int) (result string) {
	if message == 3 {
		result = "Im handler 3"
	} else if h.next != nil {
		result = h.next.Handle(message)
	}
	return
}

func (h *ConcreteHandlerC) SetNext(next Handler) Handler {
	h.next = next
	return next
}
