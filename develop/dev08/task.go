package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	goProcesses "github.com/mitchellh/go-ps"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	for {
		fmt.Print("CustomShell> ")

		input := readInput()

		// Выход
		if input == `/q` || input == `\q` || input == `quit` || input == `q` {
			break
		}

		commands := strings.Split(input, "|")

		// Запуск каждой команды
		var output []byte
		var err error
		for _, cmdStr := range commands {
			cmdStr = strings.TrimSpace(cmdStr)
			args := strings.Fields(cmdStr)

			switch args[0] {
			case "cd":
				err := changeDirectory(args)
				if err != nil {
					fmt.Printf("Ошибка: %v\n", err)
				}
			case "pwd":
				dir, err := printWorkingDirectory()
				if err != nil {
					fmt.Printf("Ошибка: %v\n", err)
				} else {
					fmt.Printf("Текущая директория: %s\n", dir)
				}
			case "echo":
				result, err := echo(args)
				if err != nil {
					fmt.Printf("Ошибка: %v\n", err)
				} else {
					fmt.Printf("%s\n", result)
				}
			case "kill":
				err := killProcess(args)
				if err != nil {
					fmt.Printf("Ошибка: %v\n", err)
				}
			case "ps":
				err = printProcesses()
				if err != nil {
					fmt.Printf("Ошибка: %v\n", err)
				}
			default:
				output, err = unknownCommand(args)
				if err != nil {
					fmt.Printf("Ошибка: %v\n", err)
				} else {
					fmt.Println(output)
				}
			}

			if err != nil {
				fmt.Printf("Ошибка: %v\n", err)
				break
			}
		}

		// Выводим результат выполнения цепочки команд
		if len(output) > 0 {
			fmt.Printf("%s\n", string(output))
		}
	}
}

// Функция для считывания ввода пользователя
func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Функция для смены директории
func changeDirectory(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("некорректное использование, попробуй cd <директория>")
	}

	dir := args[1]
	err := os.Chdir(dir)
	if err != nil {
		return err
	}

	return nil
}

// Функция для вывода текущей директории
func printWorkingDirectory() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return dir, nil
}

// Функция для команды echo
func echo(args []string) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("некорректное использование, попробуй echo <текст>")
	}
	return strings.Join(args[1:], " "), nil
}

// Функция для команды kill
func killProcess(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("некорректное использование, попробуй kill <PID>")
	}

	pid, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	err = process.Kill()
	if err != nil {
		return err
	}

	return nil
}

// Функция для команды ps
func printProcesses() error {
	sliceProc, err := goProcesses.Processes()

	if err != nil {
		return err
	}

	for _, proc := range sliceProc {
		fmt.Printf("Process name: %v process id: %v\n", proc.Executable(), proc.Pid())
	}
	return nil
}

// Функция для выполнения внешней команды
func unknownCommand(args []string) ([]byte, error) {
	cmd := exec.Command(args[0], args[1:]...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return output, nil
}
