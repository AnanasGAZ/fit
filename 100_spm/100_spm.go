// 100-struct-pointer-method
package main

import "fmt"

//------------------------ex1----------------
func ex1() {
	user := NewUser(125, "Ваня", 42)
	fmt.Println(user)
}

// Определите структуру User с полями
// ID int,
// Name string,
// Age int.
// Напишите функцию
// NewUser(id int, name string, age int) User,
// которая создает и возвращает пользователя.
type User struct {
	ID    int
	Name  string
	Age   int
	Email string
}

func NewUser(id int, name string, age int) User {
	return User{
		ID:   id,
		Name: name,
		Age:  age,
	}
}

//------------------------ex2----------------
func ex2() {
	p := NewPoint(10, 20)
	fmt.Println(p)
}

// Определите структуру
// Point с координатами X int и Y int.
// Напишите функцию NewPoint(x, y int) Point.

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) Point {
	return Point{
		x: x,
		y: y}
}

//------------------------ex3----------------
// Определите структуру Rectangle
// с полями Width int и Height int.
// Напишите функцию Area(r Rectangle) int.
type Rectangle struct {
	Width  int
	Height int
}

func ex3() {
	r := Rectangle{
		Width:  10,
		Height: 20,
	}
	fmt.Println(Area(r))
}
func Area(r Rectangle) int {
	return r.Width * r.Height
}

//------------------------ex4----------------
// Для структуры Point напишите функцию
// IsOrigin(p Point) bool,
// которая возвращает true, если обекоординаты равны нулю.
type Point1 struct {
	x int
	y int
}

func ex4() {
	p := Point1{
		x: 6,
		y: 5,
	}
	fmt.Println(IsOrigin(p))
}
func IsOrigin(p Point1) bool {
	if p.x == p.y {
		return true
	}
	return false
}

//------------------------ex5----------------

func ex5() {
	a := NewUser(125, "Ваня", 42)
	b := NewUser(126, "Петя", 42)
	fmt.Println(Older(a, b))
}

//Для структуры User напишите функцию
// Older(a, b User) User,
// которая возвращает более старшего пользователя.
// При одинаковом возрасте возвращайте a.

func Older(a, b User) User {
	if b.Age > a.Age {
		return b
	}
	return a
}

//--------------------------ex6-----------------------
// Для структуры User добавьте поле Email string.
// Напишите функцию WithEmail(u User, email string) User,
// которая возвращает измененную копию пользователя
// и не меняет исходный u.
func ex6() {
	u := NewUser(125, "Маша", 33)
	fmt.Println(WithEmail(u, "ыыы@mmm.vv"))
	fmt.Println(u)
}
func WithEmail(u User, email string) User {
	return User{
		ID:    u.ID,
		Name:  u.Name,
		Age:   u.Age,
		Email: email,
	}
}

//--------------------------ex7---------------
// Определите структуру Pair с полями A int и B int.
// Напишите функцию SwapPair(p Pair) Pair,
// возвращающую структуру с переставленными значениями.
type Pair struct {
	A int
	B int
}

func ex7() {
	p := Pair{A: 10, B: 20}
	fmt.Println(SwapPair(p))

}
func SwapPair(p Pair) Pair {
	return Pair{
		A: p.B,
		B: p.A,
	}
}

//--------------------------ex8---------------
// Для двух точек Point напишите функцию
// ManhattanDistance(a, b Point) int.
//в обычном евклидовом расстоянии используется квадратный корень,
// а в манхэттенском — просто сумма расстояний по координатам.
func ex8() {
	a := Point{x: 1, y: 2}
	b := Point{x: 4, y: 6}
	fmt.Println(ManhattanDistance(a, b))
}
func ManhattanDistance(a, b Point) int {
	//|1 - 4| + |2 - 6| = 3 + 4 = 7
	dx := a.x - b.x
	if dx < 0 {
		dx = -dx
	}

	dy := a.y - b.y
	if dy < 0 {
		dy = -dy
	}

	return dx + dy
}

//--------------------------ex9---------------
// ------------------???не поняла что надо--------------------------------
// Определите структуру Range с полями From int и To int.
// Напишите функцию NormalizeRange(r Range)Range,
// которая гарантирует From <= To.
type Range struct {
	From int
	To   int
}

// func NormalizeRange(r Range) Range {

// }

//--------------------------ex10---------------
// Напишите функцию MovePoint(p Point, dx, dy int) Point,
// которая возвращает новую точку, неизменяя исходную.
func MovePoint(p Point, dx, dy int) Point {
	return Point{
		x: dx,
		y: dy,
	}
}

//--------------------------ex11---------------

// Определите структуру Student с полями Name string и Grades []int.
// Напишите функцию
// AverageGrade(s Student) float64. Для пустого списка оценок возвращайте 0.
type Student struct {
	Name   string
	Grades []int
}

func AverageGrade(s Student) float64 {
	if len(s.Grades) == 0 {
		return 0
	}
	sum := 0
	for _, grade := range s.Grades {
		sum += grade
	}
	//нет автоматического преобразования...
	return float64(sum) / float64(len(s.Grades)) // надо  в float64?
}

//--------------------------ex12---------------
// Определите структуру Product с полями Name string и Price int.
// Напишите функцию Discounted(p Product, percent int) Product,
// возвращающую копию товара со сниженной ценой.
// Процент считать в диапазоне от 0 до 100.
type Product struct {
	Name  string
	Price int
}

func ex12() {
	p := Product{
		Name:  "Ноутбук",
		Price: 1000,
	}
	fmt.Println(Discounted(p, 0))
	fmt.Println(Discounted(p, 20))

}
func Discounted(p Product, percent int) Product {
	return Product{
		Name:  p.Name,
		Price: p.Price * (100 - percent) / 100,
	}
}
func main() {
	ex12()
}
