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
			fmt.Printf("%s", emojis[val*-1])
		}
		fmt.Println()
	}
}

func (t Tabuleiro) DebugDefualt() {
	mapa := []string{"o", "a", "b", "m"}
	for _, row := range t {
		for _, val := range row {
			fmt.Printf(" %s", mapa[val*-1])
		}
		fmt.Println()
	}
}

func (t Tabuleiro) DebugRaw() {
	for _, row := range t {
		for _, val := range row {
			if val == 0 {
				fmt.Printf(" %d", val)
			} else {
				fmt.Print(val)
			}

		}
		fmt.Println()
	}
}
