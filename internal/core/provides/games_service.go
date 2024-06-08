package provides

import "campo-minado/internal/core/domain"

type GamesService interface {
	Create(lines, columns int, mines int) (*domain.Game, error)
	Get(id string) (*domain.Game, error)
	Revel(id string, pos domain.Position) (*domain.Game, error)
}
