package main

import (
	"fmt"
)

type Comp struct {
	Data int
	Next *Comp
}


func s_push(top **Comp, D int) {
	q := &Comp{
		Data: D,
		Next: nil,
	}

	if *top == nil {
		*top = q
	} else {
		q.Next = *top
		*top = q
	}
}

func main() {
	var top *Comp
	s_push(&top, 10)
	s_push(&top, 20)
	s_push(&top, 30)

	// Выводим элементы стека
	for current := top; current != nil; current = current.Next {
		fmt.Println(current.Data)
	}

}