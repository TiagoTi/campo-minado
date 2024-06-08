//go:generate mockgen --source=board_generator.go --destination=mocks/board_generator.go --package=mocks
package requires

import "campo-minado/internal/core/domain"

type BoardGenerator interface {
	New(lines, columns, mines int) (domain.Board, error)
}
