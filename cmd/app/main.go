package main

import (
	"awesomeProject/internal/analyze"
	"awesomeProject/internal/fileio"
	"awesomeProject/internal/processing"
	"fmt"
	"time"
)

// main — Главная точка входа в программу.
// Выполняет:
// 1. Генерацию файла с числами.
// 2. Последовательную обработку чисел.
// 3. Параллельную обработку чисел с заданным числом потоков.
// 4. Анализ производительности последовательной и параллельной обработки.
func main() {
	// Ввод количества чисел для генерации
	var N int
	fmt.Print("Введите N (количество чисел для генерации): ")
	fmt.Scanln(&N)

	// Генерация файла numbers.txt с числами от 1 до N
	err := fileio.GenerateFile(N, "numbers.txt")
	if err != nil {
		fmt.Println("Ошибка при генерации файла:", err)
		return
	}
	fmt.Println("Файл numbers.txt успешно создан.")

	// Чтение чисел из файла numbers.txt
	numbers, err := fileio.ReadNumbersFromFile("numbers.txt")
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	// Определение функции обработки: умножение на 2
	processFunc := func(num int) int {
		return num * 2
	}

	// Последовательная обработка чисел
	start := time.Now()
	seqResult := processing.SequentialProcessing(numbers, processFunc)
	seqDuration := time.Since(start)
	fmt.Printf("Последовательная обработка завершена за %v.\n", seqDuration)

	// Запись результатов последовательной обработки в output_sequential.txt
	err = fileio.WriteNumbersToFile("output_sequential.txt", seqResult)
	if err != nil {
		fmt.Println("Ошибка при записи результатов последовательной обработки:", err)
		return
	}

	// Ввод количества потоков для параллельной обработки
	var M int
	fmt.Print("Введите количество потоков M: ")
	fmt.Scanln(&M)

	// Параллельная обработка чисел
	start = time.Now()
	parResult := processing.ParallelProcessing(numbers, processFunc, M)
	parDuration := time.Since(start)
	fmt.Printf("Параллельная обработка завершена за %v.\n", parDuration)

	// Запись результатов параллельной обработки в output_parallel.txt
	err = fileio.WriteNumbersToFile("output_parallel.txt", parResult)
	if err != nil {
		fmt.Println("Ошибка при записи результатов параллельной обработки:", err)
		return
	}

	// Анализ производительности (ускорение и эффективность)
	analyze.AnalyzePerformance(seqDuration.Seconds(), parDuration.Seconds(), N, M)
}
