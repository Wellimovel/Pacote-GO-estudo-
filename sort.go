package main

import (
	"fmt"
	"sort"
)

type Dados struct {
	Nome  string
	Idade int
}

type ParaNome []Dados

func (ps ParaNome) Len() int {
	return len(ps)
}

func (ps ParaNome) Less(i, j int) bool {
	return ps[i].Nome < ps[j].Nome
}
func (ps ParaNome) Swap(i, j int) { ps[i], ps[j] = ps[j], ps[i] }

func main() {
	criancas := []Dados{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	sort.Sort(ParaNome(criancas))
	fmt.Println(criancas)
}
