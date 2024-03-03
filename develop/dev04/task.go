package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func sortString(s string) string {
	sorted := strings.Split(s, "")
	sort.Strings(sorted)
	return strings.Join(sorted, "")
}

func findAnagram(words *[]string) *map[string][]string {
	anagramSets := make(map[string][]string)

	for _, word := range *words {
		sortedWord := sortString(strings.ToLower(word))
		anagramSets[sortedWord] = append(anagramSets[sortedWord], word)
	}

	result := make(map[string][]string)

	for _, value := range anagramSets {
		if len(value) > 1 {
			sort.Strings(value)
			result[value[0]] = value
		}
	}

	return &result
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	anagramSets := findAnagram(&words)

	for key, value := range *anagramSets {
		fmt.Println(key, value)
	}
}
