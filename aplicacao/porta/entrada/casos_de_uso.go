package entrada

import (
	d "campo-minado/aplicacao/dominio"
)

type SessaoCampoMinado interface {
	Novo(nome string, tamanho d.Posicao, bombas int) (*d.CampoMinado, error)
	Revelar(id string, p d.Posicao) (*d.CampoMinado, error)
	Marcar(id string, p d.Posicao) (*d.CampoMinado, error)
	Obter(id string) (*d.CampoMinado, error)
}

type GeradorPosicoes interface {
	Tabuleiro(tamanho d.Posicao, qntBomba int, posicaoInicial d.Posicao) (d.Tabuleiro, error)
}
