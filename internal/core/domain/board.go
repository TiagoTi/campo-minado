package domain

type CellStatus int

const (
	Hidden CellStatus = -(iota)
	Reveled
	Flag
	Mine
)

type Board [][]CellStatus
