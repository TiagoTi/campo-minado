package dominio

const (
	EstadoCelulaOculta EstadoCelula = -(iota)
	EstadoCelulaAberta
	EstadoCelulaBandeira
	EstadoCelulaMina
)
const (
	EstadoJogoNovo EstadoJogo = iota
	EstadoJogoAtivo
	EstadoJogoGanho
	EstadoJogoPerda
)
