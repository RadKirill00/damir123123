package main 

import (
	"context"
	"fmt"
    "time"   
)

func sumArray(arr []int, result chan int, ctx context.Context) {
    sum := 0
    for _, num := range arr {
        select {
            case <- ctx.Done():
                fmt.Println("Error ", ctx.Err())
                return 
            default:
                sum += num
                // time.Sleep(2*time.Second)
        }
    }
    result <- sum
}

func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    mid := len(arr) / 2
    part1 := arr[:mid]
    part2 := arr[mid:]

    stx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

    defer cancel()

    result := make(chan int, 2)

    go sumArray(part1, result, stx)
    go sumArray(part2, result, stx)



    sum1 := 0
    sum2 := 0
    done := 0

    for done < 2 {
        select {
        case num := <-result:
            if done == 0 {
                sum1 = num
            } else {
                sum2 = num
            }
            done++
        case <-stx.Done():
            fmt.Println("Error ", stx.Err())
            return
        }
    }
    total := sum1 + sum2

    fmt.Printf("Общая сумма: %d\n", total)
}