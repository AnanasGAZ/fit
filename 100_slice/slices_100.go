package main

import (
	"fmt"
	"math/rand/v2"
)

// ----------------------------Util---------------------------

// randomSlice возвращает слайс со случайной длиной в [minLen, maxLen]
// и случайными значениями в [minVal, maxVal].
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
func test1append() {

	s := make([]int, 5, 7)
	fmt.Printf("слайс: %v\n длинна %v емкость %v\n", s, len(s), cap(s)) //слайс: [0 0 0 0 0]
	for i := 1; i < 4; i++ {

		fmt.Printf("слайс: %v\n длинна %v емкость %v\n", append(s, i), len(s), cap(s)) //слайс: [0 0 0 0 0]

	}
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
	// s1 := randomSlice(5, 10, 10, 50)
	s2 := []int{}
	fmt.Printf("В слайсе %v\n %d четных элементов", s2, CountEven(s2))

}

// Посчитай количество чётных элементов
func CountEven(s []int) int {
	// del1 := s[0] / 2
	// remainder1 := s[0] % 2
	// println(del1, remainder1)

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

// -------------------------6------------------
func ex6() {
	s1 := randomSlice(1, 10, 5, 10)
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
	s1 := randomSlice(1, 10, 5, 10)
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
	s1 := randomSlice(10, 15, 5, 10)
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
	s := randomSlice(3, 15, 5, 10)
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
	s := randomSlice(1, 10, 5, 10)
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

func ex11() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("Исходный слайс: %v\n", s)
	SwapPairs2(s)
	fmt.Printf("Итоговый слайс %v\n", s)
}

// Поменяй местами соседние элементы: [1,2,3,4,5] -> [2,1,4,3,5]. Работай in-place.
func SwapPairs(s []int) {
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		s[left], s[right] = s[right], s[left]
	}
}
func SwapPairs2(s []int) {
	for left, right := 0, 1; right < len(s); left, right = left+2, right+2 {

		s[left], s[right] = s[right], s[left]
		fmt.Printf("левый %v правй %v \n слайс: %v\n", left, right, s)

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
	fmt.Printf("Исходный слайс: %d\n удалить элемент с индексом: %d\n", s, i)
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

// -----------------------------------
// Удали все target in-place, сохранив порядок.
// Новый backing array выделять нельзя
func ex16() {
	s := randomSlice(5, 10, 7, 10)
	target := 10
	fmt.Printf("Исходный слайс %v\n", s)
	s = RemoveAllInPlace(s, target)
	fmt.Printf("Итоговый слайс %v\n", s)
}
func RemoveAllInPlace(s []int, target int) []int {
	//пойдем как бы двумя индексами.
	// Один там, внтури обхода в for
	// а воторой будет указывать на то место, куда надо писать, как бы сдвиг

	WritenIndex := 0

	for _, value := range s {
		// fmt.Printf("Индекс %v, инд куда пишем %v, значение %v, цель %v\n", i, WritenIndex, value, target)
		if value != target {
			// fmt.Printf("зашли в условие\n")
			s[WritenIndex] = value
			WritenIndex++ //не цель - не перезаписываем
		}
	}
	// fmt.Printf("Слайс внутри %v\n", s)
	//s[начало:конец] вернем только куда записали, мусор отрежем
	return s[:WritenIndex]
}

func ex17() {
	s := randomSlice(7, 9, 5, 10)
	fmt.Printf("Исходный слайс %v\n", s)
	fmt.Printf("Итоговый слайс %v\n", KeepEvenInPlace(s))

}

// Оставь только чётные элементы, сохранив их порядок.
// Работай in-place без нового backing array
func KeepEvenInPlace(s []int) []int {
	writenIndex := 0

	for _, val := range s {
		if val%2 == 0 {
			s[writenIndex] = val
			writenIndex++
		}
	}
	return s[:writenIndex]
}

func ex18() {
	s := randomSlice(7, 10, 0, 3)
	fmt.Printf("Исходный слайс %v\n", s)
	MoveZerosToEnd(s)
	fmt.Printf("Итоговый слайс %v\n", s)

}

// Перемести все нули в конец, сохранив относительный порядок
// ненулевых элементов.
// Работай in-place и O(1) дополнительной памяти
func MoveZerosToEnd(s []int) {
	writenIndex := 0
	for _, val := range s {
		if val != 0 {
			s[writenIndex] = val
			writenIndex++
		}
	}
	for writenIndex < len(s) {
		s[writenIndex] = 0
		writenIndex++
	}
}
func ex19() {
	s := randomSlice(7, 10, 0, 3)
	fmt.Printf("Исходный слайс %v\n", s)

	fmt.Printf("Итоговый слайс %v\n", CompactNonZero(s))
}

// Сдвинь ненулевые элементы в начало слайса, сохрани порядок,
// и верни слайс только с ними.
// Хвост результата должен быть исключён через изменение len.
func CompactNonZero(s []int) []int {
	writenIndex := 0
	for _, val := range s {
		if val != 0 {
			s[writenIndex] = val
			writenIndex++
		}
	}
	return s[:writenIndex]
}

func ex20() {
	s := randomSlice(7, 10, 10, 15)
	i := 6
	val := 42
	fmt.Printf("Исходный слайс %v\n вставить значение %v, по индексу %v\n", s, val, i)
	s, err := InsertAt(s, i, val)
	if err != nil {
		fmt.Printf("не коррректный индекс %d", i)
		return
	}
	fmt.Printf("Итоговый слайс %v\n", s)
}

// Вставь value перед элементом с индексом i; i может быть равен len(s).
// Сохрани порядок. Используй append/copy;
// корректно работай и при realloc.
func InsertAt(s []int, i, value int) ([]int, error) {
	if i < 0 || i > len(s) {
		return s, fmt.Errorf("не коррректный индекс %d", i)
	}
	// добавим еще место
	s = append(s, 0)
	// возьмем кусок массива с i+1 до конца
	// и вставим туда s[i:] - тот кусок, который хотим сохранить
	copy(s[i+1:], s[i:])
	// а теперь вставим value
	s[i] = value

	return s, nil
}
func ex21() {
	s := randomSlice(7, 10, 10, 15)
	//values := randomSlice(3, 5, 40, 45)
	values := []int{}
	i := 2
	fmt.Printf("Исходный слайс %v\n вставить значение %v\n, по индексу %v\n", s, values, i)
	s, err := InsertMany(s, i, values)
	if err != nil {
		fmt.Printf("Ошибка: %v", err)
		return
	}
	fmt.Printf("Итоговый слайс %v\n ", s)

}

// Вставь весь values перед индексом i.
// values может быть пустым.
// Сохрани порядок элементов.
func InsertMany(s []int, i int, values []int) ([]int, error) {
	if len(values) == 0 {
		return s, nil
	}
	if i < 0 || i > len(s) {
		return s, fmt.Errorf("Не корректный индекс")
	}
	s = append(s, make([]int, len(values))...)
	// copy(куда_копировать, что_копировать)
	copy(s[i+len(values):], s[i:])
	copy(s[i:], values)
	return s, nil
}
func ex22() {
	// s := randomSlice(10, 15, 1, 42)
	s := []int{10, 20, 30, 40, 50}
	left := 1
	right := 3
	fmt.Printf("Исходный слайс %v\n удалить значения с индекса %v по индекс %v\n", s, left, right)
	s, err := DeleteRange(s, left, right)
	if err != nil {
		fmt.Printf("Ошибка: %v", err)
	}
	fmt.Printf("Итоговый слайс %v\n ", s)

}

// Удали полуинтервал [left:right), сохранив порядок.
// Для пустого диапазона верни исходный слайс без изменений.
func DeleteRange(s []int, left, right int) ([]int, error) {
	if left < 0 || right < 0 || left > len(s) || right > len(s) || left > right {
		return s, fmt.Errorf("Не корректные индексы")
	}
	if right-left == 0 {
		return s, nil
	}
	cutiNdex := len(s) - (right - left)
	println(cutiNdex)
	copy(s[left:], s[right:])
	return s[:cutiNdex], nil
}

func ex23() {
	s := randomSlice(7, 10, 10, 15)
	values := randomSlice(5, 5, 40, 45)
	// values := []int{}
	left := 2
	right := 5
	fmt.Printf("Исходный слайс %v\n вставить значение %v\n, по индексу от %v и до %v\n", s, values, left, right)
	s, err := ReplaceRange(s, left, right, values)
	if err != nil {
		fmt.Printf("Ошибка: %v", err)
		return
	}
	fmt.Printf("Итоговый слайс %v\n ", s)
}

// Замени полуинтервал [left:right) элементами values.
// Количество удаляемых и вставляемых элементов может различаться.
func ReplaceRange(s []int, left, right int, values []int) ([]int, error) {
	//[left - включен
	//right) - полуинтервал - не включен
	// 1 values  может быть длиннее интервала -> увеличиваем слайс
	// 2 короче  -> уменьшаем слайс
	// 3 values пусто -> просто удаляе из слайса
	// 4 left == right -> просто вставка right+1 т.к. right) - полуинтервал - не включен
	if left < 0 || right < 0 || left > len(s) || right > len(s) || left > right {
		return s, fmt.Errorf("Не корректные индексы")

	}
	result := []int{}
	result = append(result, s[:left]...)
	result = append(result, values...)
	result = append(result, s[right:]...)
	return result, nil

}
func ex24() {
	s := randomSlice(7, 10, 10, 25)
	values := randomSlice(5, 5, 40, 45)
	start := 2
	deleteCount := 3
	fmt.Printf("Исходный слайс %v\n вставить значение %v\n, с индекса %v и сколько элементов удалить %v\n", s, values, start, deleteCount)
	s, err := Splice(s, start, deleteCount, values)
	if err != nil {
		fmt.Printf("Ошибка: %v", err)
		return
	}
	fmt.Printf("Итоговый слайс %v\n ", s)

}

// Реализуй аналог splice: начиная со start удали deleteCount элементов и вставь values.
// Сохрани порядок и минимизируй лишние аллокации.
func Splice(s []int, start, deleteCount int, values []int) ([]int, error) {
	// 	s — исходный слайс;
	// start — с какого индекса начать удаление;
	// deleteCount — сколько элементов удалить;
	// values — какие элементы вставить вместо удалённых.
	if start < 0 || start > len(s) {
		return s, fmt.Errorf("некорректный start: %d", start)
	}
	if deleteCount < 0 || deleteCount > len(s)-start {
		return s, fmt.Errorf("некорректный deleteCount: %d", deleteCount)
	}
	result := []int{} // а надо как-то оптимизировать, высчитывать длинну?
	result = append(result, s[:start]...)
	result = append(result, values...)
	result = append(result, s[(start+deleteCount):]...)
	return result, nil

}

func ex25() {

	s := randomSlice(7, 10, 10, 25)
	fmt.Printf("Исходный слайс %v\n ", s)
	s2 := Clone(s)
	s2[0] = 100
	fmt.Printf("\n")
	fmt.Printf("Исходный слайс %v\n ", s)
	fmt.Printf("Клон слайс %v\n ", s2) //<----------- независимый слайс

	// Остальное просто тесты для себя
	s3 := s
	s3[0] = 30

	fmt.Printf("\n")
	fmt.Printf("Исходный слайс %v\n ", s)
	fmt.Printf("Клон слайс %v\n ", s2) //<----------- независимый слайс
	fmt.Printf("Присвоенный слайс %v\n ", s3)

	s4 := []int{}
	s4 = append(s4, s...)
	s4[0] = 40

	fmt.Printf("\n")
	fmt.Printf("Исходный слайс %v\n ", s)
	fmt.Printf("Клон слайс %v\n ", s2) //<----------- независимый слайс
	fmt.Printf("Присвоенный слайс %v\n ", s3)
	fmt.Printf("Пустой и append слайс %v\n ", s4) //<----------- независимый слайс

	s5 := make([]int, len(s))
	copy(s5, s)
	s5[0] = 50

	fmt.Printf("\n")
	fmt.Printf("Исходный слайс %v\n ", s)
	fmt.Printf("Клон слайс %v\n ", s2) //<----------- независимый слайс
	fmt.Printf("Присвоенный слайс %v\n ", s3)
	fmt.Printf("Пустой и присвоенный слайс %v\n ", s4)
	fmt.Printf("Пустой make и copy слайс %v\n ", s5) //<----------- независимый слайс

}

// Верни полную независимую копию слайса.
// Изменения результата не должны влиять на исходный слайс и наоборот
func Clone(s []int) []int {
	clone := make([]int, len(s))
	copy(clone, s)
	return clone
}
func ex26() {
	s := []int{10, 20, 30, 40, 50}
	left, right := 1, 3

	fmt.Printf("Исходный слайс: %v\n", s)

	clone, err := CloneRange(s, left, right)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	fmt.Printf("Клон диапазона [%d:%d): %v\n", left, right, clone)

	_, err = CloneRange(s, -1, 3)
	if err != nil {
		fmt.Printf("Проверка ошибки: %v\n", err)
	}
}

// Верни независимую копию полуинтервала [left:right).
// Не возвращай subslice исходного backing array.
func CloneRange(s []int, left, right int) ([]int, error) {
	//right) - полуинтервал - не включен
	if left < 0 || right > len(s) || left > right {
		return nil, fmt.Errorf("некорректные границы: left=%d, right=%d", left, right)
	}

	clone := make([]int, right-left)
	copy(clone, s[left:right]) //[left:right)
	// right — это не последний индекс, а граница после последнего включаемого элемента

	return clone, nil
}

func main() {
	ex22()
}
