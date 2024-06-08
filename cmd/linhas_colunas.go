package main

import "fmt"

const linhas int = 4
const colunas int = 5

func fabrica(linhas, colunas int) {
	var valor int = 0
	var matriz [][]int
	fmt.Printf("matriz linhas: %d x colunas: %d \n", linhas, colunas)
	matriz = make([][]int, linhas)
	for l := 0; l < len(matriz); l++ {
		matriz[l] = make([]int, colunas)
		for c := 0; c < colunas; c++ {
			matriz[l][c] = valor
			valor += 1
		}

	}

	for l := 0; l < len(matriz); l++ {
		for c := 0; c < len(matriz[l]); c++ {
			fmt.Printf(" %d ", matriz[l][c])
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {

	for l := 1; l <= linhas; l++ {
		for c := 1; c <= colunas; c++ {
			fabrica(l, c)
		}
		fmt.Println()
	}
}
