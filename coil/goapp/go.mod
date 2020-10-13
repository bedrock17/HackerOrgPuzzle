module github.com/bedrock17/HackerOrgPuzzle/coil/goapp

go 1.14

require (
  goapp/game v0.0.0
)

replace (
  goapp/game v0.0.0 => ./game
)