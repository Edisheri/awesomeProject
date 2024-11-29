package analyze

import "fmt"

// AnalyzePerformance — Анализирует производительность обработки данных.
// Выводит ускорение и эффективность, а также расчёт по закону Амдала.
// Аргументы:
// - seqDuration: Время последовательной обработки (в секундах).
// - parDuration: Время параллельной обработки (в секундах).
// - N: Количество элементов.
// - M: Количество потоков.
func AnalyzePerformance(seqDuration, parDuration float64, N, M int) {
	speedup := seqDuration / parDuration     // Ускорение
	efficiency := speedup / float64(M) * 100 // Эффективность в процентах

	fmt.Printf("Ускорение: %.2f\n", speedup)
	fmt.Printf("Эффективность: %.2f%%\n", efficiency)

	// Закон Амдала
	P := 0.9 // Доля программы, которая может быть распараллелена
	amdahlSpeedup := 1 / ((1 - P) + P/float64(M))
	fmt.Printf("Теоретическое ускорение по закону Амдала: %.2f\n", amdahlSpeedup)
}
