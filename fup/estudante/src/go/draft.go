package main

import (
	"bufio"
	"fmt"
	"os"
)

type Aluno struct {
	nome        string
	n1, n2, n3 float64
	media       float64
}

func main() {
	entrada := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(entrada, &n)
	entrada.ReadByte()

	alunos := make([]Aluno, n)

	for i := 0; i < n; i++ {
		nome, _ := entrada.ReadString('\n')
		nome = nome[:len(nome)-1]

		alunos[i].nome = nome

		fmt.Fscan(entrada, &alunos[i].n1, &alunos[i].n2, &alunos[i].n3)
		entrada.ReadByte()

		alunos[i].media = (alunos[i].n1 + alunos[i].n2 + alunos[i].n3) / 3
	}

	
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if alunos[j].media < alunos[j+1].media {
				alunos[j], alunos[j+1] = alunos[j+1], alunos[j]
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d: %s\n", i, alunos[i].nome)
		fmt.Printf("   Media: %.2f\n", alunos[i].media)
		fmt.Printf("   N1: %.2f, N2: %.2f, N3: %.2f\n",
			alunos[i].n1,
			alunos[i].n2,
			alunos[i].n3)
	}
}