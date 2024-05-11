package saida

import (
	d "campo-minado/aplicacao/dominio"
)

type RepositorioCampoMinado interface {
	Obter(id string) (*d.CampoMinado, error)
}
