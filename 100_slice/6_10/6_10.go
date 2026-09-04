package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	ex10()
}

// -------------------------Util---------------
func randomSlice2(minLen, maxLen, minVal, maxVal int) []int {
	n := rand.IntN(maxLen-minLen+1) + minLen
	s := make([]int, n)
	for i := range s {
		s[i] = rand.IntN(maxVal-minVal+1) + minVal
	}
	return s
}

// -------------------------6------------------
func ex6() {
	s1 := randomSlice2(1, 10, 5, 10)
	target := 7
	fmt.Printf("В слайсе %v\n target(%d) встречается %d раз ", s1, target, CountOccurrences(s1, target))
}

// Посчитай, сколько раз target встречается в слайсе.
func CountOccurrences(s []int, target int) int {
	var target_counter int
	for _, v := range s {
		if v == target {
			target_counter++
		}
	}
	return target_counter
}

// -------------------------7------------------00:49
func ex7() {
	s1 := randomSlice2(1, 10, 5, 10)
	target := 7
	fmt.Printf("В срезе %v\n ищем: %d\n", s1, target)
	i := IndexOf(s1, target)
	if i == -1 {
		fmt.Printf("нет такого IndexOf: %d", i)
	} else {
		fmt.Printf("Индекс первого найденного элемента: %d", i)
	}
}

// Верни индекс первого вхождения target. Если элемента нет, верни -1.
func IndexOf(s []int, target int) int {
	first_i_target := -1
	for i, v := range s {
		if v == target {
			first_i_target = i
			break
		}
	}
	return first_i_target
}

// -------------------------8------------------ 00:54
func ex8() {
	s1 := randomSlice2(10, 15, 5, 10)
	target := 9
	fmt.Printf("Слайс: %v\n target %d\n слайс индексов %v\n", s1, target, IndexesOf(s1, target))

}

// Верни слайс всех индексов,
// на которых встречается target,
// в возрастающем порядке.
func IndexesOf(s []int, target int) []int {
	var Indexes []int
	for i, v := range s {
		if v == target {
			Indexes = append(Indexes, i)
		}
	}
	//порядок индексов и так возрастающий, вроде не надо сортировать
	return Indexes
}

// -------------------------9------------------ 1:33
func ex9() {
	s := randomSlice2(3, 15, 5, 10)
	old := 5
	new := 11
	fmt.Printf("Исходный слайс: %v\n старое: %d\n новое: %d\n ", s, old, new)
	ReplaceAll(s, old, new)
	fmt.Printf("Итоговый слайс: %v", s)
}

// Замени все значения old на new прямо в исходном слайсе. Новый слайс создавать нельзя.
func ReplaceAll(s []int, old, new int) {
	for i, v := range s {
		if v == old {
			s[i] = new
		}
	}
}

// -------------------------10------------------ 1:41
func ex10() {
	s := randomSlice2(1, 10, 5, 10)
	fmt.Printf("исходный слайс: %v\n", s)
	Reverse(s)
	fmt.Printf("реверс: %v", s)
}

// Разверни слайс in-place. Дополнительный слайс создавать нельзя;
// O(1) дополнительной памяти.
func Reverse(s []int) {
	//sort.Reverse(sort.IntSlice(s))
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		s[left], s[right] = s[right], s[left]
	}
}

//1:59
