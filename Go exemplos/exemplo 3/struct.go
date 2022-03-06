package main

import "fmt"

type aluno struct {
    nome string
    disciplina string
     prova1 float64
     prova2 float64
}

func main() {

    al := aluno{nome: "Daniel", disciplina: "EDL", prova1: 7.5, prova2: 6.5 }
    fmt.Println(al)
}
