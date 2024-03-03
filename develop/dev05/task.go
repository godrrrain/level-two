package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	after := flag.Int("A", 0, "печатать +N строк после совпадения")
	before := flag.Int("B", 0, "печатать +N строк до совпадения")
	context := flag.Int("C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	count := flag.Bool("c", false, "количество строк")
	ignoreCase := flag.Bool("i", false, "игнорировать регистр")
	invert := flag.Bool("v", false, "вместо совпадения, исключать")
	fixed := flag.Bool("F", false, "точное совпадение со строкой")
	lineNum := flag.Bool("n", false, "напечатать номер строки")

	flag.Parse()

	pattern := flag.Arg(0)
	file := flag.Arg(1)

	matchLines := grep(file, pattern, *after, *before, *context, *count, *ignoreCase, *invert, *fixed, *lineNum)
	if *count {
		fmt.Println(len(matchLines))
	} else {
		printLines(matchLines)
	}
}

func grep(file, pattern string, after, before, context int, count, ignoreCase, invert, fixed, lineNum bool) []string {
	matchLines := make([]string, 0)

	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Ошибка при чтении файла: %s\n", err)
		return matchLines
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineNumber := 0
	prevLines := make([]string, 0)
	found := false

	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++

		match := false

		if fixed {
			match = strings.Contains(line, pattern)
		} else {
			if ignoreCase {
				match = strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
			} else {
				match = strings.Contains(line, pattern)
			}
		}

		if match && !invert {
			if count {
				matchLines = append(matchLines, line)
			} else {
				// Строки до совпадения
				if before > 0 {
					matchLines = append(matchLines, prevLines...)
				}

				if lineNum {
					line = fmt.Sprintf("%d:%s", lineNumber, line)
				}
				matchLines = append(matchLines, line)

				// Строки после совпадения
				if after > 0 {
					afterLines := make([]string, 0)
					for i := 0; i < after; i++ {
						if scanner.Scan() {
							afterLine := scanner.Text()
							matchLines = append(matchLines, afterLine)
							afterLines = append(afterLines, afterLine)
						} else {
							break
						}
					}
					prevLines = afterLines
				} else {
					prevLines = make([]string, 0)
				}

				// Контекстные строки
				if context > 0 {
					contextLines := make([]string, 0)
					for i := 0; i < context; i++ {
						if scanner.Scan() {
							contextLine := scanner.Text()
							matchLines = append(matchLines, contextLine)
							contextLines = append(contextLines, contextLine)
						} else {
							break
						}
					}
					prevLines = contextLines
				}
			}
			found = true

		} else if !match && invert {
			if count {
				matchLines = append(matchLines, line)
			} else {
				// Строки до и после совпадения
				if (before > 0 || after > 0) && (prevLines != nil || after > 0) {
					if lineNum {
						line = fmt.Sprintf("%d:%s", lineNumber, line)
					}
					matchLines = append(matchLines, line)
				}
				// Контекстные строки
				if context > 0 {
					contextLines := make([]string, 0)
					for i := 0; i < context; i++ {
						if scanner.Scan() {
							contextLine := scanner.Text()
							matchLines = append(matchLines, contextLine)
							contextLines = append(contextLines, contextLine)
						} else {
							break
						}
					}
					prevLines = contextLines
				} else {
					prevLines = make([]string, 0)
				}
			}
			found = true

		} else {
			// Строки до и после совпадения
			if (before > 0 || after > 0) && (prevLines != nil || after > 0) {
				prevLines = append(prevLines, line)

				if len(prevLines) > before+after {
					prevLines = prevLines[1:]
				}
			} else if context > 0 {
				prevLines = append(prevLines, line)

				if len(prevLines) > context {
					prevLines = prevLines[1:]
				}
			}
		}
	}

	if count && !found {
		matchLines = []string{"0"}
	}

	return matchLines
}

func printLines(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}
