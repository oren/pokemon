package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	_ "github.com/cayleygraph/cayley/graph/bolt"
	"github.com/cayleygraph/cayley/quad"
)

func main() {
	dbFile := flag.String("db", "db", "BoltDB file")
	csvFile := flag.String("csv", "pokemon.csv", "csv file with pokemon")
	flag.Parse()

	// Initialize the database
	graph.InitQuadStore("bolt", *dbFile, nil)

	// Open and use the database
	store, err := cayley.NewGraph("bolt", *dbFile, nil)
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Open(*csvFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ",")
		id, err := strconv.Atoi(s[0])
		if err != nil {
			log.Fatal(err)
		}
		speciesId, err := strconv.Atoi(s[2])
		if err != nil {
			log.Fatal(err)
		}
		height, err := strconv.Atoi(s[3])
		if err != nil {
			log.Fatal(err)
		}
		baseExperience, err := strconv.Atoi(s[4])
		if err != nil {
			log.Fatal(err)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		store.AddQuad(quad.Make(id, "name", s[1], "."))
		store.AddQuad(quad.Make(id, "species_id", speciesId, "."))
		store.AddQuad(quad.Make(id, "height", height, "."))
		store.AddQuad(quad.Make(id, "base_experience", baseExperience, "."))
	}

	// Now we create the path, to get to our data
	// p := cayley.StartPath(store, quad.String("id")).Out(quad.String("name"))
	p := cayley.StartPath(store)
	p = p.Has(quad.String("name"))

	it, _ := p.BuildIterator().Optimize()
	defer it.Close()

	// While we have items
	for it.Next() {
		token := it.Result()                // get a ref to a node
		value := store.NameOf(token)        // get the value in the node
		nativeValue := quad.NativeOf(value) // this converts nquad values to normal Go type

		fmt.Println(nativeValue) // print it!
	}
	if err := it.Err(); err != nil {
		log.Fatalln(err)
	}
}
