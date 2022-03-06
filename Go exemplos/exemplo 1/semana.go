package main

import "fmt"
import "time"

func main() {
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("é final de semana")
    default:
        fmt.Println("é um dia da semana")
    }
}