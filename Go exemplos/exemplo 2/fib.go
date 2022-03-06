package main

import "fmt"

func fib(n int) int {
    if n  <= 1 {
        return n
    }
    return fib(n - 1)  +  fib(n - 2)
}

func main() {
   var numero int
   fmt.Println("tamanho da sequencia de fibonacci")
   fmt.Scan(&numero);
   seq :=  make([]int, numero+1)
   for i := 1;  i <= numero;  i++ {
        seq [i]= fib(i)	
   }
   fmt.Println("sequencia de fibonacci  = ", seq)	
}