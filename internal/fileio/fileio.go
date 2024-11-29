package fileio

import (
	"bufio"
	"os"
	"strconv"
)

// GenerateFile — Создаёт файл и записывает числа от 1 до N.
// Аргументы:
// - N: Количество чисел для генерации.
// - filename: Имя создаваемого файла.
// Возвращает ошибку, если файл не удалось создать или записать.
func GenerateFile(N int, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err // Ошибка создания файла
	}
	defer file.Close() // Гарантированное закрытие файла в конце

	writer := bufio.NewWriter(file) // Буферизированная запись для повышения производительности
	for i := 1; i <= N; i++ {
		_, err := writer.WriteString(strconv.Itoa(i) + "\n") // Запись числа в файл
		if err != nil {
			return err // Ошибка записи в файл
		}
	}
	return writer.Flush() // Сохранение данных из буфера в файл
}

// ReadNumbersFromFile — Читает числа из файла и возвращает массив чисел ([]int).
// Аргументы:
// - filename: Имя файла для чтения.
// Возвращает массив чисел и ошибку (если она произошла).
func ReadNumbersFromFile(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err // Ошибка открытия файла
	}
	defer file.Close() // Гарантированное закрытие файла

	var numbers []int
	scanner := bufio.NewScanner(file) // Построчное чтение файла
	for scanner.Scan() {
		numStr := scanner.Text()
		num, err := strconv.Atoi(numStr) // Преобразование строки в число
		if err != nil {
			return nil, err // Ошибка преобразования
		}
		numbers = append(numbers, num) // Добавление числа в массив
	}
	return numbers, scanner.Err() // Возврат массива чисел и ошибок чтения
}

// WriteNumbersToFile — Записывает массив чисел в файл.
// Аргументы:
// - filename: Имя файла для записи.
// - numbers: Массив чисел для записи.
func WriteNumbersToFile(filename string, numbers []int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err // Ошибка создания файла
	}
	defer file.Close() // Гарантированное закрытие файла

	writer := bufio.NewWriter(file) // Буферизированная запись
	for _, num := range numbers {
		_, err := writer.WriteString(strconv.Itoa(num) + "\n") // Запись числа в файл
		if err != nil {
			return err // Ошибка записи
		}
	}
	return writer.Flush() // Сохранение данных из буфера в файл
}
