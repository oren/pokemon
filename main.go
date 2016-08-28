package main

import (
	"flag"

	_ "github.com/cayleygraph/cayley/graph/bolt"
)

func main() {
	dbFile := flag.String("db", "data/pokemon.boltdb", "BoltDB file")
	csvFile := flag.String("csv", "data/pokemon.csv", "csv file with pokemon")
	evolutionsFile := flag.String("evolutions", "data/evolutions.csv", "csv file with evolutions")
	flag.Parse()

	store := Setup(dbFile)
	LoadPokemons(store, csvFile)
	UpdatePikachu(store)
	LoadEvolutions(store, evolutionsFile)

	// Print(store)
	PrintEvolutions(store)
}
