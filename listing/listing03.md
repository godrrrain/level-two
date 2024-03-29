Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
Вывод:
<nil>
false

Функция Foo возвращает nil значение интерфейса error.
Интерфейс это структура с 2 полями: указатель на метаданные типа и на значение.
В данном случае типом будет *os.PathError, а значение nil.
И при сравнении с nil возвращается false как раз из-за того, что в err дескриптор типа не равен nil.

Интерфейс в Go является контрактом, и каждый тип, который хочет реализовывать конкретный интерфейс, обязан реализовать все методы этого интерфейса.

Интерфейс в исходниках представлен структурой iface из 2-х полей:
tab *itab и data unsafe.Pointer
структура itab хранит метаданные об интерфейсе и о типе (в поле *_type)
data это указатель на значение.

Пустой интерфейс как особый случай представлен упрощенной структурой eface из 2-х полей:
_type *_type и data unsafe.Pointer.
Теперь вместо *itab идет сразу информация о типе (дескриптор типа).
Пустой интерфейс имеет следующую смысловую нагрузку "по такому адресу хранится значение такого типа".

```
