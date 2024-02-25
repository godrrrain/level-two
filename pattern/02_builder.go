package pattern

/*
Строитель является порождающим паттерном проектирования уровня объекта, который определяет
процесс поэтапного построения сложного продукта.

Применение:
	1. Хотим создавать различные представления сложного объекта, используя один и тот же процесс строительства.
	2. Хотим использовать объект, процесс создания которого состоит из более чем одного независимого этапа.

Плюсы:
	1. Создаем продукты поэтапно.
	2. Переиспользуем код для других объектов.
	3. Изолируем сборку от бизнес-логики.

Минусы:
	1. Усложняем код внедрением дополнительных сущностей.
*/

type Product struct {
	Content string
}

func (p *Product) Show() string {
	return p.Content
}

type Builder interface {
	MakeHeader(str string)
	MakeBody(str string)
	MakeFooter(str string)
}

type Director struct {
	builder Builder
}

func (d *Director) Construct() {
	d.builder.MakeHeader("Header")
	d.builder.MakeBody("Body")
	d.builder.MakeFooter("Footer")
}

type ConcreteBuilder struct {
	product *Product
}

func (b *ConcreteBuilder) MakeHeader(str string) {
	b.product.Content += "<header>" + str + "</header>"
}

func (b *ConcreteBuilder) MakeBody(str string) {
	b.product.Content += "<article>" + str + "</article>"
}

func (b *ConcreteBuilder) MakeFooter(str string) {
	b.product.Content += "<footer>" + str + "</footer>"
}
