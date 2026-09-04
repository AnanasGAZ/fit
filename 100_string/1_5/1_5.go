package main

import (
	"fmt"
	"unicode/utf8"
)

// строка представляет собой срез байтов
// По умолчанию строки кодируются в UTF-8
// rune - int32 (Unicode UTF-8 ) Руна представляет эти кодовые точки Unicode в Go
func mine() {

}

// GenerateTestString возвращает строку с кириллицей, emoji и комбинируемым символом.
func GenerateTestString() string {
	return "Привет, мир! 👋 Café"
}

// Напишите функцию, которая возвращает количество байт и количество
// Unicode-символов в строке. Строка может содержать кириллицу, emoji и комбинируемые
// символы.
func StringSize(s string) (bytes int, runes int) {
	//len(). Эта функция кодирует строку, а затем возвращает общее количество байтов, занимаемых каждым символом
	return len(s), utf8.RuneCountInString(s)
}

func ex1() {
	s := GenerateTestString()
	bytes, runes := StringSize(s)
	fmt.Printf("байты: %v\n руны: %v", bytes, runes)
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
	FirstRune("Привет!")
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
