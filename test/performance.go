package main

import (
	"awesomeProject/internal/fileio"
	"awesomeProject/internal/processing"
	"fmt"
	"os"
	"time"
)

// main — Точка входа для выполнения автоматического тестирования производительности.
// Задачи:
// 1. Тестирование производительности для различных значений N (количество элементов).
// 2. Тестирование производительности для различных значений M (количество потоков).
// 3. Запись результатов тестов в файл performance_results.txt.
func main() {
	// Открываем файл для записи результатов
	resultsFile, err := os.Create("performance_results.txt")
	if err != nil {
		fmt.Println("Ошибка при создании файла для результатов:", err)
		return
	}
	defer resultsFile.Close() // Закрываем файл после завершения

	// Параметры для тестирования:
	ns := []int{10, 100, 1000, 100000, 1000000, 10000000, 100000000} // Значения N
	ms := []int{1, 2, 3, 4, 5, 10, 20, 30, 100}                      // Значения M

	for _, N := range ns {
		// 1. Генерация файла с числами от 1 до N
		err := fileio.GenerateFile(N, "numbers.txt")
		if err != nil {
			fmt.Printf("Ошибка при генерации файла для N=%d: %v\n", N, err)
			continue
		}

		// 2. Чтение чисел из файла
		numbers, err := fileio.ReadNumbersFromFile("numbers.txt")
		if err != nil {
			fmt.Printf("Ошибка при чтении файла для N=%d: %v\n", N, err)
			continue
		}

		// 3. Определение функции обработки
		// Здесь используется умножение каждого числа на 2
		processFunc := func(num int) int {
			return num * 2
		}

		// 4. Последовательная обработка
		start := time.Now()
		processing.SequentialProcessing(numbers, processFunc)
		seqDuration := time.Since(start) // Время последовательной обработки

		// 5. Параллельная обработка для каждого значения M
		for _, M := range ms {
			start = time.Now()
			processing.ParallelProcessing(numbers, processFunc, M)
			parDuration := time.Since(start) // Время параллельной обработки

			// Пропускаем тест, если время слишком мало для анализа
			if seqDuration.Seconds() == 0 || parDuration.Seconds() == 0 {
				fmt.Fprintf(resultsFile, "N=%d, M=%d, Seq=%.4fs, Par=%.4fs, Speedup=N/A, Efficiency=N/A\n",
					N, M, seqDuration.Seconds(), parDuration.Seconds())
				continue
			}

			// 6. Анализ производительности
			speedup := seqDuration.Seconds() / parDuration.Seconds() // Ускорение
			efficiency := speedup / float64(M) * 100                 // Эффективность

			// Сохраняем результаты в файл
			fmt.Fprintf(resultsFile, "N=%d, M=%d, Seq=%.4fs, Par=%.4fs, Speedup=%.2f, Efficiency=%.2f%%\n",
				N, M, seqDuration.Seconds(), parDuration.Seconds(), speedup, efficiency)
		}
	}

	fmt.Println("Тестирование завершено. Результаты сохранены в performance_results.txt")
}
