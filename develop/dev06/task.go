package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	f := flag.Int("f", 1, "выбрать поля (колонки)")
	d := flag.String("d", "\t", "использовать другой разделитель")
	s := flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()

	if *f <= 0 {
		log.Fatal("f <= 0")
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(`Для выхода, введите \q`)
	fmt.Println("Текст:")
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == `\q` || txt == `q` || txt == `/q` {
			break
		}
		splitTxt := strings.Split(txt, *d)
		if *s && !strings.Contains(txt, *d) {
			fmt.Println("")
		} else if len(splitTxt) < *f {
			fmt.Println(txt)
		} else {
			fmt.Println(splitTxt[*f-1])
		}
	}
}
