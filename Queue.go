// Реализуем структуру данных очередь
// Структура работает по принципу «первым пришёл — первым ушёл» (FIFO — first in, first out)
package main

import (
	"errors"
)

type Queue []int

func (q *Queue) Add(a int) {
	*q = append(*q, a)
}
func (q *Queue) Pop() (int, error) {
	if len(*q) == 0 {
		return 0, errors.New("Queue is empty")
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element, nil
}

func (q *Queue) IsExist() bool {
	return len(*q) != 0
}
