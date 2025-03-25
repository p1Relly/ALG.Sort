package main

import (
	"fmt"
)

func radixSort(sl []int) []int {
	digitCounts := 10      // количество различных цифр
	maxLengthOfNumber := 4 // максимальная длина числа
	p := 1                 // степень 10. Нужна для получения очередного разряда

	pocket := make([][]int, digitCounts) // массив для распределения элементов по "корзинам"
	for i := 0; i < len(pocket); i++ {
		pocket[i] = make([]int, 0)
	}

	for i := 0; i < maxLengthOfNumber; i++ { // проходимся по разрядам
		for j := len(sl); j > 0; j-- { // проходимся по числам
			index := sl[j-1] / p % 10                      // находим индекс корзины
			pocket[index] = append(pocket[index], sl[j-1]) // добавляем
		}

		count := len(sl) - 1               // на каком месте вставляем в первоначальном списке
		for j := 0; j < digitCounts; j++ { // проходимся по корзине
			for k := 0; k < len(pocket[j]); k++ { // проходимся по элементам очередной корзины
				sl[count] = pocket[j][k] // перебрасываем обратно в первоначальный список
				count--                  // уменьшаем место вставки элемента в первоначальном списке
			}
			pocket[j] = pocket[j][:0] // очищаем корзину
		}
		p *= 10 // получаем следующую степень
	}
	return sl
}

func main() {
	var N int
	fmt.Scan(&N)

	sl := make([]int, N)
	for i := range N {
		fmt.Scan(&sl[i])
	}

	for _, v := range radixSort(sl) {
		fmt.Printf("%d ", v)
	}
}
