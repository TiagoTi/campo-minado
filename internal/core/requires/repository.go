//go:generate mockgen -source=repository.go -destination=mocks/repository.go -package=mocks
package requires

import "campo-minado/internal/core/domain"

type GamesRepository interface {
	Save(new domain.Game) (saved *domain.Game, err error)
	Get(id string) (saved *domain.Game, err error)
}
