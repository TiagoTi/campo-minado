package servicos

import (
	"campo-minado/aplicacao/dominio"
	"campo-minado/aplicacao/porta/entrada"
	"campo-minado/aplicacao/porta/saida"
	"errors"
	"fmt"
	"log"
)

func NovaSessaoCampoMinado(s dominio.GeradorPosicoes,
	r saida.RepositorioCampoMinado) (entrada.SessaoCampoMinado, error) {
	return &servico{sorteador: s, repositorio: r}, nil
}

type servico struct {
	sorteador   dominio.GeradorPosicoes
	repositorio saida.RepositorioCampoMinado
}

func (s *servico) Novo(nome string, tamanho dominio.Posicao, bombas int) (*dominio.CampoMinado, error) {
	if e := validarEntrada(nome, tamanho, bombas); e != nil {
		return nil, e
	}
	return &dominio.CampoMinado{
		Nome:    nome,
		Estado:  dominio.EstadoJogoNovo,
		Tamanho: tamanho,
		Bombas:  bombas,
	}, nil
}

func (s *servico) Revelar(id string, p dominio.Posicao) (*dominio.CampoMinado, error) {
	jogo, errObt := s.repositorio.Obter(id)
	if errObt != nil {
		log.Printf("erro ao consultar jogo por id: %q", errObt.Error())
		return nil, errors.New("não foi possivél consultar o jogo")
	}
	if (jogo.Estado == dominio.EstadoJogoPerda) || (jogo.Estado == dominio.EstadoJogoGanho) {
		return jogo, errors.New("este jogo já foi encerrado")
	}
	if jogo.Estado == dominio.EstadoJogoNovo {
		tabuleiro, errSorteio := s.sorteador.Tabuleiro(jogo.Tamanho, jogo.Bombas, p)
		if errSorteio != nil {
			log.Printf("erro ao realizar sorteio das bombas %q", errSorteio.Error())
			return nil, errSorteio
		}
		jogo.Tabuleiro = tabuleiro
		// salvar
	}
	if p.X < 0 {
		return jogo, errors.New("a coluna deve iniciar da posicao 0")
	}
	if p.Y < 0 {
		return jogo, errors.New("a linha deve iniciar da posicao 0")
	}

	// limite maximo da coluna  - exportar para o tabuleiro
	if p.X > jogo.Tamanho.X {
		return jogo, fmt.Errorf("a coluna deve estar entre 0 e %d", jogo.Tamanho.X)
	}
	// limite maximo da coluna  - exportar para o tabuleiro
	if p.Y > jogo.Tamanho.Y {
		return jogo, fmt.Errorf("a linha deve estar entre 0 e %d", jogo.Tamanho.X)
	}
	// encontrou bomba? - exportar para tabuleiro
	if jogo.Tabuleiro.Mina(p) {
		// if jogo.Tabuleiro[p.X][p.Y] == dominio.EstadoCelulaMina {
		jogo.Estado = dominio.EstadoJogoPerda
		return jogo, nil
	}
	return &dominio.CampoMinado{Tabuleiro: jogo.Tabuleiro}, nil
}

func (s *servico) Marcar(id string, p dominio.Posicao) (*dominio.CampoMinado, error) {
	return nil, nil
}
func (s *servico) Obter(id string) (*dominio.CampoMinado, error) {
	return nil, nil
}
