package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	"github.com/cayleygraph/cayley/quad"
	uuid "github.com/satori/go.uuid"
)

func Setup(dbFile *string) *cayley.Handle {
	// Initialize the database
	graph.InitQuadStore("bolt", *dbFile, nil)

	// Open and use the database
	store, err := cayley.NewGraph("bolt", *dbFile, nil)
	if err != nil {
		log.Fatalln(err)
	}

	return store
}

func LoadPokemons(store *cayley.Handle, csvFile *string) {
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

		// IRI format: <https://my-domain.com/9f1dae45-6d98-11e6-871e-843a4b0f5a10>
		uuid := quad.IRI("https://my-domain.com/" + uuid.NewV1().String())

		store.AddQuad(quad.Make(uuid, quad.IRI("rdf:id"), id, nil))
		store.AddQuad(quad.Make(uuid, quad.IRI("rdf:type"), quad.IRI("https://my-domain.com/pokemon"), nil))
		store.AddQuad(quad.Make(uuid, quad.IRI("schema:name"), s[1], nil))
		store.AddQuad(quad.Make(uuid, quad.IRI("rdf:species_id"), speciesId, nil))
		store.AddQuad(quad.Make(uuid, quad.IRI("rdf:height"), height, nil))
		store.AddQuad(quad.Make(uuid, quad.IRI("rdf:base_experience"), baseExperience, nil))
	}
}

func UpdatePikachu(store *cayley.Handle) {
	// find uuid of pikacho
	p := cayley.StartPath(store).Has(quad.IRI("schema:name"), quad.String("pikacho"))
	vals, err := p.Iterate(nil).AllValues(nil)
	if err != nil {
		log.Fatalln(err)
	} else if len(vals) == 0 {
		log.Fatalln("pikacho not found")
	}
	uuid := vals[0].Native().(quad.IRI)

	// change pikacho to pikachu
	t := cayley.NewTransaction()
	t.RemoveQuad(quad.Make(uuid, quad.IRI("schema:name"), "pikacho", nil))
	t.AddQuad(quad.Make(uuid, quad.IRI("schema:name"), "pikachu", nil))
	err = store.ApplyTransaction(t)

	if err != nil {
		log.Fatalln(err)
	}
}

func LoadEvolutions(store *cayley.Handle, csvFile *string) {
	file, err := os.Open(*csvFile)
	if err != nil {
		log.Fatal("Error while opening evolutions csv file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ",")
		targetPokemon, err := strconv.Atoi(s[0])
		if err != nil {
			log.Fatal("Error converting string to int", err)
		}

		if s[3] == "" {
			continue
		}

		sourcePokemon, err := strconv.Atoi(s[3])
		if err != nil {
			log.Fatal("Error converting string to int", err)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal("Error while scanning", err)
		}

		// find uuid of source and target
		p := cayley.StartPath(store).Has(quad.IRI("rdf:id"), quad.Int(sourcePokemon))
		vals, err := p.Iterate(nil).AllValues(nil)
		if err != nil {
			log.Fatalln(err)
		} else if len(vals) == 0 {
			log.Fatalln("source pokemon not found")
		}
		sourcePokemonUUID := vals[0].Native().(quad.IRI)

		p = cayley.StartPath(store).Has(quad.IRI("rdf:id"), quad.Int(targetPokemon))
		vals, err = p.Iterate(nil).AllValues(nil)
		if err != nil {
			log.Fatalln(err)
		} else if len(vals) == 0 {
			log.Fatalln("target pokemon not found")
		}
		targetPokemonUUID := vals[0].Native().(quad.IRI)

		store.AddQuad(quad.Make(sourcePokemonUUID, quad.IRI("rdf:evolves_to"), targetPokemonUUID, nil))
	}
}

func Print(store *cayley.Handle) {
	// Now we create the path, to get to our data: names of Pokemon
	p := cayley.StartPath(store).Has(quad.IRI("rdf:type"), quad.IRI("https://my-domain.com/pokemon")).Out(quad.IRI("schema:name"))

	it, _ := p.BuildIterator().Optimize()
	it, _ = store.OptimizeIterator(it)
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

func PrintEvolutions(store *cayley.Handle) {
	// Now we create the path, to get to our data: names of Pokemon that evolved twice
	p := cayley.StartPath(store).Out(quad.IRI("rdf:evolves_to")).Out(quad.IRI("rdf:evolves_to")).Out(quad.IRI("schema:name"))

	it, _ := p.BuildIterator().Optimize()
	it, _ = store.OptimizeIterator(it)
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
