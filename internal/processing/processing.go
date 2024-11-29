package processing

import "sync"

// SequentialProcessing — Последовательная обработка массива чисел.
// Аргументы:
// - numbers: Массив чисел для обработки.
// - processFunc: Функция обработки одного числа.
// Возвращает массив обработанных чисел.
func SequentialProcessing(numbers []int, processFunc func(int) int) []int {
	result := make([]int, len(numbers)) // Результирующий массив
	for i, num := range numbers {
		result[i] = processFunc(num) // Обработка каждого числа
	}
	return result
}

// ParallelProcessing — Параллельная обработка массива чисел.
// Аргументы:
// - numbers: Массив чисел для обработки.
// - processFunc: Функция обработки одного числа.
// - M: Количество потоков.
// Возвращает массив обработанных чисел.
func ParallelProcessing(numbers []int, processFunc func(int) int, M int) []int {
	result := make([]int, len(numbers))     // Результирующий массив
	var wg sync.WaitGroup                   // Синхронизация потоков
	chunkSize := (len(numbers) + M - 1) / M // Размер каждой части (округление вверх)

	for i := 0; i < M; i++ {
		startIdx := i * chunkSize
		endIdx := (i + 1) * chunkSize
		if endIdx > len(numbers) {
			endIdx = len(numbers) // Последний поток обрабатывает остаток
		}

		wg.Add(1) // Увеличиваем счётчик потоков
		go func(start, end int) {
			defer wg.Done() // Уменьшаем счётчик после завершения потока
			for j := start; j < end; j++ {
				result[j] = processFunc(numbers[j]) // Обработка числа
			}
		}(startIdx, endIdx)
	}
	wg.Wait() // Ожидание завершения всех потоков
	return result
}
