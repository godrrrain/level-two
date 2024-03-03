package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func unpack(s string) (string, error) {
	arr := []rune(s)

	builder := strings.Builder{}

	ind := 0
	for ind < len(arr) {
		curCh := arr[ind]
		if unicode.IsLetter(curCh) {
			builder.WriteString(string(curCh))
		} else if unicode.IsDigit(curCh) {
			cnt := 0

			j := ind
			prevInd := ind
			for j < len(arr) && unicode.IsDigit(arr[j]) {
				buf, _ := strconv.Atoi(string(arr[j]))
				cnt = cnt*10 + buf
				j++
				ind++
			}
			if prevInd > 0 {
				for i := 0; i < cnt-1; i++ {
					builder.WriteString(string(arr[prevInd-1]))
				}
			} else {
				return "", errors.New("wrong string")
			}
			continue
		} else if string(curCh) == `\` {
			if ind < len(arr)-1 {
				builder.WriteString(string(arr[ind+1]))
				ind++
			} else {
				return "", errors.New("wrong string")
			}
		}
		ind++
	}

	return builder.String(), nil
}

func main() {
	newStr, err := unpack(`a4bc2d5e`)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(newStr)
	}

	newStr, err = unpack(`abcd`)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(newStr)
	}

	newStr, err = unpack(`qwe\45`)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(newStr)
	}

	newStr, err = unpack(`45`)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(newStr)
	}
}
