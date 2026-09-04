package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	ex5()
}

// ----------------------------Util---------------------------

// randomSlice возвращает слайс со случайной длиной в [minLen, maxLen]
// и случайными значениями в [minVal, maxVal].
func randomSlice(minLen, maxLen, minVal, maxVal int) []int {
	n := rand.IntN(maxLen-minLen+1) + minLen
	s := make([]int, n)
	for i := range s {
		s[i] = rand.IntN(maxVal-minVal+1) + minVal
	}
	return s
}

// -----------------------------1---------------------------
func ex1() {
	s1 := randomSlice(0, 20, -100, 100)
	fmt.Printf("слайс: %v\nсумма: %d\n", s1, Sum(s1))

	// нулевой слайс (nil): len == 0, Sum вернёт 0
	var s2 []int
	fmt.Printf("слайс: %v\nсумма: %d\n", s2, Sum(s2))
}

// Sum возвращает сумму всех элементов слайса.
// Для пустого слайса возвращает 0.
func Sum(s []int) int {
	var total int
	for _, v := range s {
		total += v
	}
	return total
}

// -----------------------------2---------------------------
func ex2() {
	s1 := randomSlice(0, 10, 50, 100)
	min1, ok1 := Min(s1)
	fmt.Printf("слайс: %v\nmin значение: %d, есть минимум: %t\n", s1, min1, ok1)

	var s2 []int
	min2, ok2 := Min(s2)
	fmt.Printf("слайс: %v\nmin значение: %d, есть минимум: %t\n", s2, min2, ok2)

}

// Верни минимальный элемент и true.
// Для пустого слайса верни 0 и false.
func Min(s []int) (int, bool) {
	var min int
	var have_min bool
	if len(s) != 0 {
		min = s[0]
		have_min = true
		for _, v := range s {
			if v < min {
				min = v
			}
		}
	} else {
		min = 0
		have_min = false
	}

	return min, have_min
}

// -----------------------------3---------------------------
func ex3() {
	s1 := randomSlice(1, 20, 1, 25)
	max1, ok1 := Max(s1)
	fmt.Printf("Слайс: %v\n max значение: %d, есть максимум? %t\n", s1, max1, ok1)

	var s2 []int
	max2, ok2 := Max(s2)
	fmt.Printf("Слайс: %v\n max значение: %d, есть максимум? %t\n", s2, max2, ok2)
	//fmt.Printf("слайс: %v\nmin значение: %d, есть минимум: %t\n", s2, min2, ok2)
}

// Верни максимальный элемент и true.
// Для пустого слайса верни 0 и false.
func Max(s []int) (int, bool) {
	var max int
	var have_max bool

	if len(s) != 0 {
		max = s[0]
		have_max = true
		for _, v := range s {
			if v > max {
				max = v
			}
		}
	} else {
		max = 0
		have_max = false
	}
	return max, have_max
}

// -----------------------------4---------------------------
func ex4() {
	s1 := randomSlice(1, 10, 10, 50)
	fmt.Printf("В слайсе %v\n %d четных элементов", s1, CountEven(s1))

}

// Посчитай количество чётных элементов
func CountEven(s []int) int {
	del1 := s[0] / 2
	remainder1 := s[0] % 2
	println(del1, remainder1)

	var even_counter int

	for _, v := range s {
		if v%2 == 0 {
			even_counter++
		}
	}
	return even_counter
}

// -----------------------------5---------------------------
// Верни true, если target встречается в слайсе хотя бы один раз.
func ex5() {
	s := randomSlice(1, 10, 4, 50)
	target := 5
	fmt.Printf("Есть ли в слайсе %v\n элемент %d ?\n  ответ: %t", s, target, Contains(s, target))

}

func Contains(s []int, target int) bool {
	have_target := false
	for _, v := range s {
		if v == target {
			have_target = true
		}
	}
	return have_target
}
