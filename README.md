# Documentação do jogo Campo Minado

## Requisitos Funcionais
### RF01 - jogador deve criar nova sessão de jogo
O aplicativo deve permitir que um jogador crie uma nova sessão de jogo, informando dados para que seja possível configurar o tabuleiro de uma partida do jogo campo minado.
> Como jogador, gostaria de criar um novo jogo. Informando a quantidade de minas, a quantidade de linhas e a quantidade de colunas que o tabuleiro do jogo deve ter.
### RF02 - jogador deve conseguir consulta uma sessão de jogo
O aplicativo deve permitir que o jogador consulte um jogo existente, para que ele vizualize o estado atual da partida.
> Eu como jogador gostaria de consultar um jogo já existente para continuar jogando de onde parei, ou apenas para visualizar como está o jogo.
### RF03 - jogador deve confugir revelar um campo
O aplicativo deve permitir que o jogador solicite a revelação de um campo no tabuleiro de uma determinada sessão de jogo, para isso o jogador deve informar o identificador do jogo, e uma posição válida (uma linha e uma coluna), a sessão deve atualizar na base de dados o novo estado do jogo e retornar ao jogador a instancia de como ficou o jogo após a sua ação de revelar um campo.
> Como jogador posso revelar um campo do tabuleiro
### RF04 - jogador deve confugir revelar todos os campos
em andamento
O aplicativo deve permitir que o jogador solicite a revelação de todos os campo no tabuleiro de uma determinada sessão de jogo, para isso o jogador deve informar o identificador do jogo, a sessão deve atualizar na base de dados o novo estado do jogo e retornar ao jogador a instancia de como ficou o jogo após a sua ação de revelar um campo.
> eu como jogador gostaria de revelar todos os campos do jogo
### RF05 - jogador deve confugir listar todas as sessões de jogos
em andamento: eu como jogador, gostaria de poder listar todas as sessões de jogos
### RF06 - jogador deve conseguir adicionar bandeira em um campo
em andamento: eu como jogador, gostaria de poder listar todas as sessões de jogos
O aplicativo deve permitir que o jogador adicione uma bandeira em um campo que ele acha que possui mina terrestre
(esse requisito afeta o requisito, RF04 )
### RF07 - jogador deve confugir listar todas as sessões de jogos
em andamento: O aplicativo de permitir que o usuário remova a bandeira de uma célula marcada com bandeira
(esse requisito afeta o requisito, RF04 )

## Requisitos Não funcionais
### RNF01 o aplicativo deve gerar logs de negócio
O aplicativo deve gerer logs de modo a reproduzir todos os passos tomados pelo jogador. O log deve conter uma presentação do estado do jogo e do tabuleiro, assim como a ação tomada pelo usuário. Logs de negócio e não de aplicação.
## Regras de negócios
- RN001: O tabuleiro do jogo deve ter um quantidade de linhas maior que um.
- RN002: O tabuleiro do jogo deve ter uma quantidade de colunas maior que um.
- RN003: Uma sessão de jogo deve possuir no mínimo uma mina.
- RN004: Uma sessão de jogo deve ter no máximo uma quantidade de mina menor que a relação: quantidade de linhas X quantidade de colunas.
- RN005: O jogador só pode realizar revelações em posições válidas do tabuleiro.
- RN006: O jogador só pode revelar células passiveis de revelação: (células ocultas)
- RN007: Ao realizar a primeira revelação (jogo estado de novo) o jogo de sortear a minas aleatoriamente ignorando a posição informada.
- RN008: Ao revelar uma mina o jogo fica com estado de perda
- RN009: Dado que uma partida esteja em andamento (estado diferente de perda ou ganho), ao consultar um jogo o tabuleiro deve seguir a seguintes regra de exibição: 
  1. Minas Ocultas: As minas existentes no tabuleiro devem permanecer ocultas, ou seja, não devem ser reveladas ao usuário.
  2. Bandeiras Marcadas: Se o jogador tiver marcado uma bandeira em algum campo, essa marcação deve ser mantida na posição correta. Ou seja, a bandeira deve permanecer sobre o campo onde o jogador a colocou.
  3. Campos com Dicas: Campos que já foram revelados e possuem dicas (números indicando a quantidade de minas adjacentes) devem exibir essas dicas ao usuário. Isso ajuda o jogador a entender a situação do jogo e tomar decisões estratégicas.
  4. Campos Ocultos: Campos que ainda não foram revelados devem permanecer ocultos, mantendo o suspense e desafiando o jogador a fazer suas escolhas com base em informações limitadas.
  5. Campos Abertos: Campos que foram abertos e não possuem minas adjacentes devem ser representados como campos abertos, sem dicas, indicando ao jogador que não há minas próximas.
- RN010: A quantidade máxima de bandeira é a mesma quantidade minas existentes no jogo
- RN011: O jogador só marcar um campo com bandeira dado que a primeira revelação já tenha sido realizada.
- RN012: O jogador só pode marcar um campo com bandeira em posições válidas do tabuleiro.
- RN013: O jogador só pode marcar um campo com bandeira em células passiveis de revelação: (células ocultas)
- RN014: O jogo deve ser gravado a cada alteração no estado do jogo ou do tabuleiro (revelar, marcar ou descmar celula)
- RN015: Ao revelar todas as celular o jogo deve analisar se o jogo foi ganho (todos os campos com minas estão ocultos ou marcados com bandeira e não exite outros campos a sem mina a revelar) ou perda (ainda existe algum campo a revelar e pelo menos uma mina sem bandeira)
- RN016: o jogador só pode remover bandeira de um campo  marcado com bandeira
- RN017: ao consultar um jogo finalizado (perda ou ganho) todos os campos devem ser exibitos ao usuáro (inclundo bandeiras aonde tinha minas e bandeira aonde não haviam minas), campos que foram revelados e não revelados sem minas, todas as minas existentes
## Modelo de caso de uso
### Atores
#### Usuário Terminal Linux
Pessoa que consegue executar o binário local utilizando o jogo no modo de linha de comando.
#### Usuário Web
Pessoa que consegue acessar o jogo por meio do aplicativo web
#### Usuário Grífico Linux
usuário da máquina local que consegue executar o jogo no modo gráfico
#### Jogador
Qualquer pessoa ou aplicativo que consegue interagir com o jogo com alguma interface disponível do jogo.

### Caso de uso
#### CSU01: Criar jogo
##### Requisitos
RF01
##### Importância
risco alto, prioridade alta: sem esse não existem os demais casos
##### Sumário: 
>	para o jogador consiga joga, a qualquer momento o sistema deve permitir que ele configure um novo jogo. O jogador solicita ao sistema uma nova partida passando as informações de quantas linhas e colunas deve ter seu novo tabuleiro assim como a quantidade de minas.
##### Ator Primário
Jogador
#####  Fluxo

| Jogador                                                                                | Sistema                                          |
| -------------------------------------------------------------------------------------- | ------------------------------------------------ |
| jogador informa: a quantidade linha, coluna do tabuleiro e a quantidade de minas minas | sistema retorna o a identificação do jogo criado |
##### Fluxo de exceção
- Erros de usuário: se o jogo tiver algum problema de validação de dados um erro com o problema sera retornado ao jogador.
- Erros de sistema: se o jogo tiver algum problema de rede, disco processamento  outras naturezas sistêmicas, uma mensagem informativa de falha de operação sera enviada ao usuário.
##### Pós condições
o sistema de persistir (gravar/salvar)o estado do jogo em um repositório de dados para que o jogador possa consultar e voltar ao jogo num momento futuro.
##### Regras de negócio
RN001, RN002, RN003, RN004

---
#### CSU02: Consultar um jogo
##### Requisitos
RF02
##### Importância
risco alto, prioridade media: sem esse não existem os demais casos
##### Sumário
>	O aplicativo deve permitir que o jogador consulte uma sessão de jogo existente, para que ele possa visualizar ou então continuar jogando. Para isso o jogador informa qual é o identificador do jogo, o aplicativo retornara o jogo no ultimo estado de gravação.

##### Ator Primário
Jogador
##### Fluxo

| Jogador                               | Sistema                                                                  |
| :------------------------------------ | :----------------------------------------------------------------------- |
| Informa o identificador<br>de um jogo | retorna o jogo (estado do jogo e tabuleiro)<br>conforme o estado do jogo |
##### Fluxo de exceção
- Erros de usuário: se o jogador enviar um código inválido, o jogo não retornará nenhum jogo, e uma mensagem de erro vai ser devolvida ao usuário;
- Erros de sistema: se o jogo tiver algum problema de rede, disco processamento  outras naturezas sistêmicas, uma mensagem informativa de falha de operação sera enviada ao usuário.

##### Pós condições
o sistema de retornar uma cópia do jogo no estado do jogo em que foi repositório de dados na ultima operação de alteração/criação
##### Regras de negócio
RN009, RN017

---
#### CSU03: Revelar um campo
##### Requisitos
RF03
##### Importância
risco médio, prioridade media -
##### Sumário
>	O aplicativo deve permitir que o jogador o jogador revele um campo de uma sessão de jogo existente, para que ele possa revelar um campo, deve informar uma posição válida ao sistema assim como o id do jogo que ele deseja revelar.

##### Ator Primário
Jogador
##### Fluxo

| Jogador                                             | Sistema                                                                                                       |
| :-------------------------------------------------- | :------------------------------------------------------------------------------------------------------------ |
| Informa o identificador<br>de um jogo e uma posição | sistema valida as entradas<br><br>retorna o jogo (estado do jogo e tabuleiro)<br>conforme novo estado do jogo |
##### Fluxo de exceção
- Erros de usuário: se o jogador enviar um código inválido, o jogo não retornará nenhum jogo, e uma mensagem de erro vai ser devolvida ao usuário;
- Erros de usuário: se o jogador enviar um uma posição inválida, o jogo não retornará nenhum jogo, e uma mensagem de erro vai ser devolvida ao usuário;
- Erros de sistema: se o jogo tiver algum problema de rede, disco processamento  outras naturezas sistêmicas, uma mensagem informativa de falha de operação sera enviada ao usuário.

##### Pós condições
o sistema de retornar uma cópia do jogo no estado do jogo em que foi repositório de dados na ultima operação de alteração/criação
##### Regras de negócio
RN005, RN006, RN007, RN008, RN014, RN015, RN017

---
## Repositório
### Ferramentas

[asdf](https://asdf-vm.com/pt-br/guide/introduction.html)
[visual code studio](https://code.visualstudio.com/)

### Referencias:

[Hexagonal Architecture in Go](https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3)

[git@github.com:HunCoding/golang-architecture.git](https://github.com/HunCoding/golang-architecture)

[git@github.com:TiagoTi/go-apperror.git](https://github.com/TiagoTi/go-apperror)

[Debugging go in vscode with environment variables](https://reuvenharrison.medium.com/using-visual-studio-code-to-debug-a-go-program-with-environment-variables-523fea268271)

[git@github.com:matiasvarela/minesweeper-hex-arch-sample.git](https://github.com/matiasvarela/minesweeper-hex-arch-sample/tree/master)

[https://campo-minado.com/](https://campo-minado.com/)

[git@github.com:golang/mock.git](https://github.com/golang/mock)

[Go Fuzzing](https://go.dev/doc/security/fuzz/#glossary) / [Tutorial: Getting started with fuzzing](https://go.dev/doc/tutorial/fuzz)

[Descomplicando "Arquitetura Hexagonal"](https://youtu.be/V7JnDDQY1m0?si=VHTZf71KaXOe2uYa)

---
## Informação adicionais (em construção)
### Versão das tecnologias com asdf
`/.tool-versions`
```.txt
golang 1.22.0

```

### Bibliotecas

### matiasvarela/errors
utilizada para criar erros da aplicação
`github.com/matiasvarela/errors`

```go
package apperrors
import "github.com/matiasvarela/errors"
var (
	NotFound         = errors.Define("not_found")
	IllegalOperation = errors.Define("illegal_operation")
	InvalidInput     = errors.Define("invalid_input")
	Internal         = errors.Define("internal")
)
```

---


### O que é o jogo campo minado segundo o chat-gpt
Campo Minado é um jogo clássico de quebra-cabeça e estratégia. O objetivo do jogo é limpar um campo minado sem detonar nenhuma mina. Aqui estão os principais pontos sobre como jogar:

    Grade: O jogo é jogado em uma grade retangular com células ocultas que podem conter uma mina ou não.

    Revelar células: Você revela as células clicando nelas. Se uma célula contém uma mina, e você a revela, o jogo termina.

    Números: Se você revelar uma célula que não contém uma mina, um número aparecerá. Esse número indica quantas minas estão escondidas imediatamente adjacentes a essa célula (em todas as oito direções).

    Marcação de minas: Você pode marcar células que você acha que contêm minas com uma bandeira. Geralmente, isso é feito com um clique direito do mouse. Marcar minas ajuda a evitar clicar acidentalmente nelas e permite que você identifique mais facilmente quais outras células são seguras para revelar.

    Ganhar o jogo: Para ganhar, você precisa revelar todas as células que não contêm minas.

    Estratégia: O jogo exige cuidado e raciocínio. Uma boa estratégia é começar pelas células nas quais você pode determinar facilmente se há ou não uma mina com base nas informações numéricas fornecidas pelas células adjacentes.

Campo Minado não apenas testa sua paciência e atenção, mas também sua capacidade de pensar logicamente sob pressão. Você gostaria de dicas específicas sobre estratégias para melhorar no jogo ou tem alguma outra pergunta sobre ele?

---

### Hexagonal Architecture in Go
[medium](https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3)

---
Principio da abstração: só as caracteristicas relevantes a resolução de um problema devem ser consideradas

o que me levou a ler o livro foi revisitar/visitar o que há entre a ideia do que quero construir do ponto de visto de uma ideia, sem nenhuma referencia tecnológica, até um plano em que outra pessoa consiga ler e saber o que tem que fazer e como ela pode fazer.
Lendo o livro, notei que gerenciamento de complexidade também é uma boa vantagem dados que passo dividir meu problema em aspectos individuais para tratar especificidades...

cada elemento gráfico possuem uma sintaxe e uma semantica
sintase: como eu desenho
semantica: o que significa
(como eu extendo a sintaxe? como eu extendo a semantica)

--
as 5 visões que posso ter do meu software:
- visão de casos de uso:vista externa, interações entre  agentes externos e o sistema (guia a proximas visões).
- visão de projeto: suporte a visão de extrutura e de comportamento das funcionalidades externas
- visão de implementação: versões do sistema por agrupamento de modulos
- visão de implantação: forma qfisica que o sistemas e subsitema detem e como se conectame essas partes
- visão de processo: condições de paralelismo, concorrencias sincronização e desempenho do sistema.
----

Idenifdicando os atores do campo minado:

usuario linux via linha de comando
aplicativo que faça requisições web/json
qualquer aplicativo que interplete serviços soap
qualquer aplicativo que leia e escreva no pubsub
qualquer aplicativo que consiga lidar com socket no linux (ipc)

0 jogador pode consultar um sessão de jogo existente. (para poder saber os detalhes de uma partida/sessão de jogo ou poder interagir com o jogo)

o jogador pode solicitar a revelação de um campo em um jogo existente
O jogoador pode solicitar revelar todos campos sem bandeira em um jogo existente
o jogador pode criar um novo jogo

Usuario Local Linux -> Jogador
App React (cliente http) -> Jogador
App GTK (cliente socket linux) -> Jogador
