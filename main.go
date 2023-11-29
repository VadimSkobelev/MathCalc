package main

import (
	"log"

	maths "MathsCalc/pkg"
)

func main() {
	in := "input.txt"   // Входной файл.
	out := "output.txt" // Выходной файл.
	err := maths.Maths(in, out)
	if err != nil {
		log.Fatalln("Ошибка: ", err)
	}
}
