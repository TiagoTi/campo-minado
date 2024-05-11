package dominio

type Posicao struct {
	X int
	Y int
}

func (posicao Posicao) Igual(outraPosicao Posicao) bool {
	return outraPosicao.Y == posicao.Y && outraPosicao.X == posicao.X
}
func (posicao Posicao) DiagonalInferiorDireita() Posicao {
	return Posicao{X: posicao.X - 1, Y: posicao.Y + 1}
}
func (posicao Posicao) DiagonalInferiorEsquerda() Posicao {
	return Posicao{X: posicao.X - 1, Y: posicao.Y - 1}
}
func (posicao Posicao) DiagonalSuperiorDireita() Posicao {
	return Posicao{X: posicao.X + 1, Y: posicao.Y + 1}
}
func (posicao Posicao) DiagonalSuperiorEsquerda() Posicao {
	return Posicao{X: posicao.X + 1, Y: posicao.Y - 1}
}
func (posicao Posicao) HorizontalDiretita() Posicao {
	return Posicao{X: posicao.X, Y: posicao.Y + 1}
}
func (posicao Posicao) HorizontalEsquerda() Posicao {
	return Posicao{X: posicao.X, Y: posicao.Y - 1}
}
func (posicao Posicao) VerticalAbaixo() Posicao {
	return Posicao{X: posicao.X - 1, Y: posicao.Y}
}
func (posicao Posicao) VerticalAcima() Posicao {
	return Posicao{X: posicao.X + 1, Y: posicao.Y}
}

type Posicoes []Posicao

func (posicoes Posicoes) Existe(outraPosicao Posicao) bool {
	for _, posicao := range posicoes {
		if posicao.Y == outraPosicao.Y && posicao.X == outraPosicao.X {
			return true
		}
	}
	return false
}
