package dominio

import "fmt"

type Tabuleiro [][]EstadoCelula

func (t Tabuleiro) Mina(p Posicao) bool {
	return t[p.X][p.Y] == EstadoCelulaMina
}

func (t Tabuleiro) Debug() {
	emojis := []string{"\U0001F7EA", "\U00002b1c", "\U0001F6A9", "\U0001F4A3"}
	for _, row := range t {
		for _, val := range row {
			fmt.Print(emojis[val])
			fmt.Print("|")
		}
		fmt.Println()
	}
}
