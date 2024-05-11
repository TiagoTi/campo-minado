package servicos

import (
	"campo-minado/aplicacao/dominio"
	"errors"
)

func validaNome(s string) error {
	if s == "" {
		return errors.New("nome não informado")
	}
	return nil
}
func validarTamanho(t dominio.Posicao) error {
	if t.X < 0 {
		return errors.New("o jogo não pode ter quantidade negativa de coluna")
	}
	if t.X < 1 {
		return errors.New("o jogo dever ter uma quantidade de coluna maior que zero")
	}
	if t.X < 2 {
		return errors.New("o jogo dever ter uma quantidade de coluna maior que 1")
	}
	if t.Y < 0 {
		return errors.New("o jogo não pode ter quantidade negativa de linha")
	}
	if t.Y < 1 {
		return errors.New("o jogo dever ter uma quantidade de linha maior que zero")
	}
	if t.Y < 2 {
		return errors.New("o jogo dever ter uma quantidade de linha maior que 1")
	}
	return nil
}
func validarQuantidadeBombas(b int) error {
	if b < 0 {
		return errors.New("quantidade de bombas não pode ser negativa")
	}
	if b < 1 {
		return errors.New("o jogo deve ter pelo menos uma bomba")
	}
	return nil
}
func validarEntrada(nome string, tamanho dominio.Posicao, bombas int) error {
	if e := validaNome(nome); e != nil {
		return e
	}
	if e := validarTamanho(tamanho); e != nil {
		return e
	}
	if e := validarQuantidadeBombas(bombas); e != nil {
		return e
	}
	if bombas >= tamanho.X*tamanho.Y {
		return errors.New("a quantidade de bombas deve ser menor que o total de espaço do tabuleiro")
	}
	return nil
}
