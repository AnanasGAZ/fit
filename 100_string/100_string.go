package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// строка представляет собой срез байтов
// По умолчанию строки кодируются в UTF-8
// rune - int32 (Unicode UTF-8 ) Руна представляет эти кодовые точки Unicode в Go

func test1str() {
	str := "Hello"
	str2 := "Привет"
	println(len(str), len(str2))
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c   ", str2[i])
		// fmt.Println(str[i])
	}
	for _, val := range str2 {
		fmt.Printf("%c   ", val)

	}
}

// GenerateTestString возвращает строку с кириллицей, emoji и комбинируемым символом.
func GenerateTestString() string {
	return "Привет, мир! 👋 Café Go 🚀 Java Art"
}

// -----------------------------1------------------------
func ex1() {
	s := GenerateTestString()
	bytes, runes := StringSize(s)
	fmt.Printf("байты: %v\n руны: %v", bytes, runes)
}

// Напишите функцию, которая возвращает количество байт и количество
// Unicode-символов в строке. Строка может содержать кириллицу, emoji и комбинируемые
// символы.
func StringSize(s string) (bytes int, runes int) {
	//len(). Эта функция кодирует строку, а затем возвращает общее количество байтов, занимаемых каждым символом
	return len(s), RuneCountInStringMy(s)
}

func RuneCountInStringMy(s string) (n int) {
	for range s {
		n++
	}
	return n
}

// -----------------------------2------------------------
// Напишите функцию, которая проверяет, состоит ли строка только из ASCII-символов.
func IsASCII(s string) bool {
	//ASCII-символы имеют значения от 0 до 127
	for i := 0; i < len(s); i++ {
		if s[i] > 127 {
			return false
		}
	}

	return true
}

// ----------------------------3--------------------------
func ex3() {
	s := "Привет!"
	r, have_rune := FirstRune(s)
	fmt.Printf("Исходная строка %v , \n первый Unicode-символ %c, первый Unicode-символ?': %v", s, r, have_rune)
}

// Напишите функцию, которая возвращает первый Unicode-символ строки и признак его
// наличия. Пустую строку обработайте без panic.

func FirstRune(s string) (rune, bool) {
	if s == "" {
		return 0, false
	}

	r, _ := utf8.DecodeRuneInString(s)
	return r, true
}

// --------------------------4-----------
func ex4() {
	str := GenerateTestString()
	r, have_rune := LastRune(str)
	fmt.Printf("Исходная строка %v , \n руна %c, есть ли руна: %v", str, r, have_rune)
}

// 4. Напишите функцию, которая возвращает последний Unicode-символ строки и признак
// его наличия. Не считайте, что один символ занимает один байт.
func LastRune(s string) (rune, bool) {
	// 	 t одновременно является:
	// ASCII-символом; (116)
	// Unicode-символом;
	// rune в Go со значением 116.

	if len(s) == 0 {
		return 0, false
	}
	r, _ := utf8.DecodeLastRuneInString(s)
	return r, true
}

// func LastRune(s string) (rune, bool) {
// 	runes := []rune(s)

// 	if len(runes) == 0 {
// 		return 0, false
// 	}

//		return runes[len(runes)-1], true
//	}
//
// ----------------------------5-------------
func ex5() {
	fmt.Println(Reverse("Привет 🚀"))
}

// 5. Разверните строку по Unicode-символам. Результат для корректной UTF-8 строки тоже
// должен быть корректной UTF-8 строкой.
func Reverse(s string) string {
	runes := []rune(s)

	for left, right := 0, len(runes)-1; left < right; left, right = left+1, right-1 {
		runes[left], runes[right] = runes[right], runes[left]
	}

	return string(runes)
}

// ----------------------------6-------------
func ex6() {

	fmt.Println(RemoveRune("Привет Долли 🚀", '🚀'))
}

// Удалите из строки все вхождения заданного Unicode-символа.
func RemoveRune(s string, target rune) string {
	result := make([]rune, 0, utf8.RuneCountInString(s))

	for _, val := range s {
		if val != target {
			result = append(result, val)
		}
	}

	return string(result)
}

// ----------------------------7-------------
func ex7() {
	fmt.Println(ReplaceRune("Привет 🚀 hello 🚀 Aloha 🚀", '🚀', '👋'))
}

// 7. Замените все вхождения одного Unicode-символа другим, не используя
// strings.ReplaceAll.
func ReplaceRune(s string, old, new rune) string {
	r := []rune(s)
	for i, val := range r {
		if val == old {
			r[i] = new
		}
	}
	return string(r)
}

// ----------------------------8-------------
func ex8() {
	s := "топот"

	fmt.Printf("%q — палиндром: %v\n", s, IsPalindrome(s))
}

//  8. Проверьте, является ли строка палиндромом по Unicode-символам. Регистр и пробелы
//     учитываются.
func IsPalindrome(s string) bool {
	r := []rune(s)
	// r := make([]rune, 0, len(r))
	// for _, val := range r {
	// 	if !unicode.IsSpace(val) {
	// 		r = append(r, unicode.ToLower(val))
	// 	}
	// }

	for left, right := 0, len(r)-1; left < right; left, right = left+1, right-1 {
		if r[left] != r[right] {
			return false
		}
	}

	return true
}

// --------------------------9---------------------
func ex9() {
	fmt.Println(TrimUnicodeSpace("    Привет 🚀 hello 🚀 Aloha     "))
	fmt.Println("Пустой:")
	fmt.Println(TrimUnicodeSpace("   "))
}

// 9. Удалите пробельные символы в начале и конце строки без использования
// strings.TrimSpace. Учитывайте Unicode-пробелы.
func TrimUnicodeSpace(s string) string {
	r := []rune(s)

	start := 0
	end := len(r) //если len(r)= 0, то  len(r)- 1 - паника
	//а так же если s := "   ", то end > start // 2 > 3 == false

	for start < len(r) && unicode.IsSpace(r[start]) {
		start++
	}
	for end > start && unicode.IsSpace(r[end-1]) {
		end--
	}
	return string(r[start:end])
}

// --------------------------------10-----------------------------
// 10. Замените любую последовательность Unicode-пробелов одним обычным пробелом и
// удалите пробелы по краям.
func ex10() {
	fmt.Print(NormalizeSpaces("    Привет     🚀         hello    🚀 Aloha     "))
	fmt.Println("|")
	fmt.Print("Пустой:")
	fmt.Print(NormalizeSpaces("   "))
	fmt.Println("|")
}
func NormalizeSpaces(s string) string {
	r := []rune(s)
	result := make([]rune, 0, len(r))
	haveSpace := false

	for _, val := range r {
		if unicode.IsSpace(val) { // пробел всегда пропускаем
			if len(result) > 0 { //не начало строки
				haveSpace = true // пометим если пробел есть
			}
			continue
		}
		if haveSpace {
			result = append(result, ' ')
			haveSpace = false
		}
		result = append(result, val)
	}
	return string(result)
}

// ----------------------------------------11-------------------------------------
func main() {
	ex8()
}
