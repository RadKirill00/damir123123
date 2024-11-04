package main

import (
    "fmt"
    "sync"
)

var once sync.Once
var config string

func initialize() {
    fmt.Println("Initializing...")
    config = "Configuration data"
}

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            fmt.Printf("Goroutine %d attempting to initialize\n", id)
            once.Do(initialize)
            fmt.Printf("Goroutine %d: config = %s\n", id, config)
        }(i)
    }

    wg.Wait()
    fmt.Println("All goroutines completed.")
}