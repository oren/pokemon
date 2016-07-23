package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

func init() {
}

type Pokemon struct {
	Id             int
	Name           string
	Species_id     int
	Height         int
	BaseExperience int
}

func main() {
	dbFile := flag.String("db", "db", "BoltDB file")
	csvFile := flag.String("csv", "pokemon.csv", "csv file with pokemon")
	flag.Parse()

	storage := New(*dbFile)
	pokemon := importPokemon(*csvFile)
	storage.Insert(pokemon)
	storage.Read()
}

func importPokemon(csvFile string) []Pokemon {
	file, err := os.Open(csvFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pokemon := []Pokemon{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ",")
		id, err := strconv.Atoi(s[0])
		speciesId, err := strconv.Atoi(s[2])
		height, err := strconv.Atoi(s[3])
		baseExperiment, err := strconv.Atoi(s[4])

		if err != nil {
			log.Fatal(err)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		pokemon = append(pokemon, Pokemon{id, s[1], speciesId, height, baseExperiment})
	}

	return pokemon
}
