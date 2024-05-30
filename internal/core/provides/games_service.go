package provides

import "campo-minado/internal/core/domain"

type GamesService interface {
	Create(lines, columns uint, mines uint) (*domain.Game, error)
}
