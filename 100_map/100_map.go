package main

import (
	"fmt"
	"maps"
	"math/rand/v2"
	"sort"
)

// --------------------------Util----------------------------------
func randomSlice2(minLen, maxLen, minVal, maxVal int) []int {
	len := rand.IntN(maxLen-minLen+1) + minLen
	s := make([]int, len)
	for i := range s {
		s[i] = rand.IntN(maxVal-minVal+1) + minVal
	}
	return s
}

// --------------------------1----------------------------------
// Напишите функцию CountInts(nums []int) map[int]int,
//
//	которая возвращает количество вхождений
//
// каждого числа.
func ex1() {
	s := randomSlice2(5, 10, 22, 24)
	fmt.Printf("Исходный массив: %v\n", s)
	counts_nums := CountInts(s)
	fmt.Printf("Итоговая мапа: %v\n", counts_nums)
}
func CountInts(nums []int) map[int]int {
	counts := make(map[int]int)

	for _, num := range nums {
		counts[num]++
	}

	return counts
}

// --------------------------2----------------------------------
func ex2() {
	words := []string{"go", "java", "go", "abap", "java"}
	LastIndexes := LastIndex(words)
	fmt.Printf("Итоговая мапа: %v\n", LastIndexes)
}

// Напишите функцию LastIndex(words []string) map[string]int,
// которая для каждой строки хранит
// индекс ее последнего вхождения
func LastIndex(words []string) map[string]int {
	lastIndexes := make(map[string]int) // string  будет строка

	for index, word := range words {
		lastIndexes[word] = index
	}

	return lastIndexes
}

// --------------------------3----------------------------------

func ex3() {
	s := randomSlice2(7, 10, 42, 45)
	fmt.Printf("Исходный: %v\n", s)
	s_Unic := Unique(s)
	fmt.Printf("Исходный: %v\n", s_Unic)
}

// Напишите функцию Unique(nums []int) []int, которая удаляет повторы,
//
//	сохраняя порядок первых вхождений.
//
// Для проверки уникальности используйте map[int]struct{}.
// ключ — число int;
// значение — пустая структура struct{}.
// struct{}{} не занимает дополнительной памяти,
// поэтому её часто используют, когда нужно хранить
// только факт наличия элемента
func Unique(nums []int) []int {
	seen := make(map[int]struct{})
	unique := make([]int, 0, len(nums))

	for _, num := range nums {
		//value, exists := seen[5] - exists - есть ли вообще ключ 5 в map.
		// value нам не нужен. Поэтому вместо имени переменной ставится _
		_, exists := seen[num]
		if exists {
			//if _, exists := seen[num]; exists {
			continue
		}

		seen[num] = struct{}{}
		unique = append(unique, num) //в любом случае обход мапы или сразу сохранять в
	}

	return unique
}

// --------------------------4----------------------------------

func ex4() {
	have := []string{"go", "java", "abap", "питон"}
	need := []string{"go", "java", "abap", "питон"}
	fmt.Printf("Есть ли значения из need в have? : %v\n", ContainsAll(have, need))
	have1 := []string{"go", "java", "abap"}
	need1 := []string{"go", "java", "abap", "питон"} //false
	fmt.Printf("Есть ли значения из need в have? : %v\n", ContainsAll(have1, need1))
	have2 := []string{"go", "java", "abap", "питон"}
	need2 := []string{"go", "java", "abap"} //true
	fmt.Printf("Есть ли значения из need в have? : %v\n", ContainsAll(have2, need2))
}

// Напишите функцию ContainsAll(have, need []string) bool.
// Она должна вернуть true, если каждое
// значение из need встречается в have хотя бы один раз.
func ContainsAll(have, need []string) bool {
	haveSet := make(map[string]struct{}, len(have)) //len(have) - меньше перераспределений памяти

	//соберем из have  уникальные значения
	for _, value := range have {
		haveSet[value] = struct{}{}
	}
	// обойдем need  и посмотрим есть ли значения в have
	for _, value := range need {
		if _, exists := haveSet[value]; !exists {
			return false
		}
	}
	return true
}

// ----------------------------------5--------------------------
func ex5() {
	prices := map[string]int{
		"banana": 120,
		"Яблоко": 80,
		"orange": 100,
		"pear":   150,
	}

	fmt.Println("Мапа:", prices) //Но! мапа уже сама их упорядочевает
	for key := range prices {    // а, нет, это fmt  для строк упорядочил...
		fmt.Println(key)
	}
	fmt.Println("Ключи по алфавиту:", SortedKeys(prices))
}

// Напишите функцию SortedKeys(m map[string]int) []string, которая возвращает ключи мапы в
// лексикографическом порядке.
// ========================================???????=========================================
// В Go сравнение строк выполняется лексикографически по UTF-8-байтам.
// Для обычных корректных UTF-8 строк порядок Unicode-символов сохраняется,
//
//	но это не локализованная сортировка. Например, кириллица будет идти после латиницы.
func SortedKeys(m map[string]int) []string {
	keys := []string{}

	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

// ----------------------------------6--------------------------
func ex6() {
	prices := map[string]int{
		"banana":  100,
		"apple":   50,
		"яблоко":  200,
		"gjvbljh": 150,
	}
	fmt.Printf("Мапа: %#v\n Итого: %d", prices, SumValues(prices))
}

// Напишите функцию SumValues(m map[string]int) int,
// которая вычисляет сумму всех значений мапы.
// Пустая и nil-мапа должны давать ноль.
func SumValues(m map[string]int) int {
	total := 0

	for _, val := range m {
		total = total + val
	}
	return total
}

// ----------------------------------7--------------------------

// Напишите функцию Invert(m map[string]int) (map[int]string, error).
// Если два исходных ключа
// имеют одинаковое значение, функция должна вернуть ошибку.
func ex7() {
	fruts := map[string]int{
		"apple":  100,
		"orange": 200,
		"ananas": 300,
	}
	frut_invert, err := Invert(fruts)
	if err != nil {
		fmt.Printf("Ошибка: %v", err)
	} else {
		fmt.Printf("Мапа фрукты: %v\n Инвертированная мапа: %v", fruts, frut_invert)
	}
}
func Invert(m map[string]int) (map[int]string, error) {
	m_inv := map[int]string{}
	// := make(map[string]struct{}, len(have))
	for key, value := range m {
		// v, exists := m[k]
		if _, exists := m_inv[value]; exists {
			return nil, fmt.Errorf("Значение %d встречается у нескольких ключей", value)
		}
		m_inv[value] = key

	}
	return m_inv, nil
}

// ----------------------------------8--------------------------
func ex8() {
	a := map[string]int{
		"apple":  100,
		"orange": 200,
		"персик": 100,
	}
	b := map[string]int{
		"apple":  100,
		"orange": 200,
		"груша":  300,
	}
	fmt.Printf("Итоговая мапа: %v", MergeCounts(a, b))
}

// Напишите функцию MergeCounts(a, b map[string]int) map[string]int,
// которая создает новую мапу и
// складывает значения совпадающих ключей. Исходные мапы изменять нельзя
func MergeCounts(a, b map[string]int) map[string]int {
	result := make(map[string]int)
	for key_a, val_a := range a {
		val_b, exists := b[key_a]
		if exists {
			result[key_a] = (val_a + val_b)
		}
	}
	return result
}

// ----------------------------------9--------------------------
func ex9() {
	m := map[string]int{
		"apple":  0,
		"orange": 100,
		"груша":  0,
	}
	DeleteZeroValues(m)
	fmt.Printf("Удалили нулевые %v", m)
}

// Напишите функцию DeleteZeroValues(m map[string]int),
// которая на месте удаляет все элементы с
// нулевым значением. Функция должна корректно работать с nil
func DeleteZeroValues(m map[string]int) {
	for key, value := range m {
		if value == 0 {
			delete(m, key)
		}
	}
}

// -------------------------------10-----------------------------
func ex10() {
	original := map[string][]int{
		"key": {1, 2, 3},
	}
	clone := Clone(original)
	clone["key"][1] = 42
	fmt.Println(original["key"]) // [1 2 3]
	fmt.Println(clone["key"])    // [1 42 3]

}

// Напишите функцию Clone(m map[string][]int) map[string][]int,
// создающую полностью
// независимую копию мапы и каждого вложенного слайса.
func Clone(m map[string][]int) map[string][]int {
	// начиная с Go 1.21
	// m_clone := maps.Clone(m)// но слайсы не клонирует
	if m == nil {
		return nil
	}
	m_clone := make(map[string][]int, len(m))

	for key, val := range m {
		if val == nil {
			m_clone[key] = nil
			continue
		}

		valCopy := make([]int, len(val))
		copy(valCopy, val) // копия слайсов
		m_clone[key] = valCopy
	}

	return m_clone

}

// -------------------------------11-----------------------------
func ex11() {
	a := map[string]int{
		"apple":  100,
		"orange": 200,
		// "персик": 100,
		"груша": 300,
	}
	b := map[string]int{
		"apple":  100,
		"orange": 200,
		"груша":  300,
	}
	fmt.Printf("Равны ли мапы? %v", Equal2(a, b))
}

// Напишите функцию Equal(a, b map[string]int) bool,
// которая сравнивает две мапы по наборам
// ключей и значениям, не используя reflect.DeepEqual
func Equal1(a, b map[string]int) bool {
	// 	maps.Equal сравнивает:
	//  количество ключей;
	//  наличие одинаковых ключей;
	//  одинаковые значения у этих ключей.
	return maps.Equal(a, b)
}
func Equal2(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for key, valueA := range a {
		valueB, exists := b[key]

		if !exists || valueA != valueB {
			return false
		}
	}
	return maps.Equal(a, b)
}

func main() {
	ex11()
}
