// Реализуем структуру данных стек
// Структура работает по принципу «последним пришёл — первым ушёл» (LIFO — last in, first out)
package main

import (
	"errors"
)

type Stack []int

func (s *Stack) Add(a int) {
	*s = append(*s, a)
}
func (s *Stack) Pop() (int, error) {
	if len(*s) == 0 {
		return 0, errors.New("Stack is empty")
	}
	element := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return element, nil
}
func (s *Stack) IsExist() bool {
	return len(*s) != 0
}

