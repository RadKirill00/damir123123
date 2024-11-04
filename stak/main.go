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
    } 

     q.Next = *top
    *top = q
    
}

func s_delete_key(top **Comp, N int) {
 q := *top
 var prev *Comp

 for q != nil {
	if q.Data == N {
		if q == *top {
			*top = q.Next
		} else {
			prev.Next = q.Next
		}
		q.Data = 0
		q.Next = nil
		return 
	}
	prev = q
	q = q.Next
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

	s_delete_key(&top, 20)

	fmt.Println("After deletion:")
	for current := top; current != nil; current = current.Next {
		fmt.Println(current.Data)
	}

}