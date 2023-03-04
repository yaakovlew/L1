package twenty_second_task

import (
	"fmt"
	"math/big"
)

// Функцция суммы
func Sum(a, b big.Float) *big.Float {
	s := new(big.Float)
	s.Add(&a, &b)
	return s
}

// Функция произведения
func Multiply(a, b big.Float) *big.Float {
	m := new(big.Float)
	m.Mul(&a, &b)
	return m
}

// Функция частного
func Division(a, b big.Float) *big.Float {
	q := new(big.Float)
	q.Quo(&a, &b)
	return q
}

// Функция разности
func Diference(a, b big.Float) *big.Float {
	d := new(big.Float)
	d.Sub(&a, &b)
	return d
}

// Проверка
func TwentyTwo() {
	a := big.NewFloat(987.654321)
	b := big.NewFloat(123.456789)
	fmt.Println(Sum(*a, *b).String())
	fmt.Println(Diference(*a, *b).String())
	fmt.Println(Multiply(*a, *b).String())
	fmt.Println(Division(*a, *b).String())
}
