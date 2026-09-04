package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	ex14()
}

func randomSlice(minLen, maxLen, minVal, maxVal int) []int {
	//генерирует случайное число n в диапазоне от minLen до maxLen
	// !!!!!!!!!
	//rand.IntN(x) возвращает случайное число от 0 до x-1.
	// !!!!!!!!!   без +1 получим minLen ... maxLen-1
	// rand.IntN(5) может вернуть 0, 1, 2, 3 или 4
	// 	Допустим:
	// minLen := 3
	// maxLen := 7
	// тогда:
	// n := rand.IntN(7 - 3 + 1) + 3
	// n := rand.IntN(5) + 3
	// 	rand.IntN(5) даст число от 0 до 4:
	// 0 + 3 = 3
	// 1 + 3 = 4
	// 2 + 3 = 5
	// 3 + 3 = 6
	// 4 + 3 = 7

	len_s := rand.IntN(maxLen-minLen+1) + minLen
	s := make([]int, len_s)

	for i := range s {
		s[i] = rand.IntN(maxVal-minVal+1) + minVal
	}
	return s
}

func ex11() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("Исходный слайс: %v\n", s)
	SwapPairs(s)
	fmt.Printf("Итоговый слайс %v\n", s)
}

// Поменяй местами соседние элементы: [1,2,3,4,5] -> [2,1,4,3,5]. Работай in-place.
func SwapPairs(s []int) {
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		s[left], s[right] = s[right], s[left]
	}
}
func ex12() {
	s := randomSlice(1, 10, 5, 10)
	left := 1
	right := 7
	fmt.Printf("Исходный слайс: %v\n индексы: л - %d, п - %d\n", s, left, right)
	err := ReverseRange(s, left, right)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("Итоговый: %v", s)
	}
}

// Разверни in-place участок с индексами от left до right включительно.
// Некорректные границы
// должны возвращать ошибку, не меняя слайс.
func ReverseRange(s []int, left, right int) error {
	if left < 0 || right >= len(s) || left > right {
		return fmt.Errorf("некорректные границы: left=%d, right=%d", left, right)
	}

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}

	return nil
}
func ex13() {
	s := randomSlice(1, 10, 5, 9)
	i := 2
	fmt.Printf("Исходный слайс: %v\n удалить элемент с индексом: %d\n", s, i)
	s, err := RemoveAt(s, i)
	if err == nil {
		fmt.Printf("Итоговый слайс: %v\n", s)
	} else {
		fmt.Printf("Ошибка: %v", err)
	}
}

// Удали элемент с индексом i, сохранив порядок остальных элементов.
// Используй существующий
// backing array, если это возможно.
func RemoveAt(s []int, i int) ([]int, error) {
	if i < 0 || i >= len(s) {
		return s, fmt.Errorf("некорректный индекс: %d", i)
	}
	// copy(куда_копировать, что_копировать) еще не запомнила
	// пример:
	// индексы:  0   1   2   3   4
	// значения: 10  20  30  40  50
	// i = 2
	// s[i:] - Это срез от индекса 2 до конца: [30 40 50] - куда копируем.
	// s[i+1:] - [40 50] - что копируем.
	// copy([30 40 50], [40 50])
	copy(s[i:], s[i+1:])
	//fmt.Printf("Итоговый слайс: %v\n", s)
	return s[:len(s)-1], nil
}
func ex14() {
	s := randomSlice(1, 5, 11, 25)
	i := 3
	fmt.Printf("Исходный слайс: %d\n удалить элемент с индексом: %d\n", s)
	s, err := RemoveAtFast(s, i)
	if err != nil {
		fmt.Printf("Ошибка %v\n", err)
	} else {
		fmt.Printf("Итоговый массив: %v\n", s)
	}

}

// Удали элемент с индексом i за O(1),
// если порядок остальных элементов сохранять не требуется
func RemoveAtFast(s []int, i int) ([]int, error) {
	if i < 0 || i >= len(s) {
		return s, fmt.Errorf("некорректный индекс: %d", i)
	}
	//Так как порядок элементов сохранять не нужно,
	// можно заменить удаляемый элемент последним элементом слайса.
	// Это работает за O(1)
	last := len(s) - 1 // индекс от 0  индекс от 0 индекс от 0 индекс от 0 индекс от 0 индекс от 0 индекс от 0 индекс от 0 индекс от 0
	s[i] = s[last]
	s[last] = 0 // очищаем ставший лишним элемент

	return s[:last], nil
}
func ex15() {
	s := randomSlice(1, 10, 20, 25)
	target := 22
	fmt.Printf("Исхлдный: %v\n target: %d\n", s, target)
	s1 := RemoveAll(s, target)
	fmt.Printf("Итоговый: %v\n", s1)
}

// Верни новый независимый слайс без элементов, равных target.
// Результат не должен alias-ить
// входной слайс.
// ????????????????????????????????Не должен alias-ить входной слайс????????????????????????????????
// result := s[:0] - создает алиас
// слайс:
// указателя на массив данных;
// длины;
// вместимости.
func RemoveAll(s []int, target int) []int {
	result := []int{}

	for _, value := range s {
		if value != target {
			result = append(result, value)
		}
	}

	return result
}
