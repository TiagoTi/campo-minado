package dominio

type EstadoCelula int
type EstadoJogo int

type CampoMinado struct {
	ID        string
	Nome      string
	Estado    EstadoJogo
	Tabuleiro Tabuleiro
	Tamanho   Posicao
	Bombas    int
}

func NovoCampoMinado(id string, nome string, tamanho Posicao, bombas int) CampoMinado {
	return CampoMinado{
		Bombas:    bombas,
		Estado:    EstadoJogoNovo,
		ID:        id,
		Nome:      nome,
		Tabuleiro: Tabuleiro{},
		Tamanho:   tamanho,
	}
}
