package maths

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Считывание из входного файла списка математических выражений вида 5-2=?, расчёт их и запись в выходной файл.
func Maths(in, out string) error {
	// Читаем из файла.
	fileIn, err := os.ReadFile(in)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	// Регулярное выражение для поиска.
	re := regexp.MustCompile("([0-9]+)([+-/*//])([0-9])([=])[?]")

	// Поиск подходящих строк.
	submatches := re.FindAllStringSubmatch(string(fileIn), -1)

	// Открытие выходного файла.
	fileOut, err := os.OpenFile(out, os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	defer fileOut.Close()

	// Буферизация.
	writer := bufio.NewWriter(fileOut)

	// Поиск чисел и операций в строках. Расчёт. Запись в буфер.
	for _, s := range submatches {

		a, err := strconv.Atoi(s[1])
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		b, err := strconv.Atoi(s[3])
		if err != nil {
			return fmt.Errorf("%v", err)
		}

		var res int

		switch s[2] {
		case "+":
			res = a + b
		case "-":
			res = a - b
		case "*":
			res = a * b
		case "/":
			if b == 0 {
				return fmt.Errorf("деление на ноль. %v", err)
			}
			res = a / b
		}

		resOut := []byte(strconv.Itoa(res)) // Конвертация результата в слайс байтов.

		// Запись в буфер.
		writer.Write([]byte(s[1]))
		writer.Write([]byte(s[2]))
		writer.Write([]byte(s[3]))
		writer.Write([]byte(s[4]))
		writer.Write(resOut)
		writer.WriteString("\n")

	}
	writer.Flush() // Запись в файл.
	return nil
}
