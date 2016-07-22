package main

import (
	"fmt"
	"log"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	"github.com/cayleygraph/cayley/quad"
)

type Storage struct {
	Store *cayley.Handle
}

func New() *Storage {
	storage := &Storage{}
	// Create a brand new graph
	store, err := cayley.NewMemoryGraph()
	storage.Store = store
	if err != nil {
		log.Fatalln(err)
	}

	return storage
}

func (s *Storage) Insert(pokemon []Pokemon) {
	for _, pok := range pokemon {
		// fmt.Println("pok", pok.Id)
		s.Store.AddQuad(quad.Make(pok.Id, "name", pok.Name, "."))
		s.Store.AddQuad(quad.Make(pok.Id, "species_id", pok.Species_id, "."))
		s.Store.AddQuad(quad.Make(pok.Id, "height", pok.Height, "."))
		s.Store.AddQuad(quad.Make(pok.Id, "base_experience", pok.BaseExperience, "."))
	}
}

func (s *Storage) Read() {
	// Now we create the path, to get to our data
	p := cayley.StartPath(s.Store, quad.Int(1)).Out(quad.String("name"))

	// Now we get an iterator for the path (and optimize it, the second return is if it was optimized,
	// but we don't care for now)
	it, _ := p.BuildIterator().Optimize()
	// remember to cleanup after yourself
	defer it.Close()

	// Now for each time we can go to next iterator
	nxt := graph.AsNexter(it)
	// remember to cleanup after yourself
	defer nxt.Close()

	// While we have items
	for nxt.Next() {
		token := it.Result()                // get a ref to a node
		value := s.Store.NameOf(token)      // get the value in the node
		nativeValue := quad.NativeOf(value) // this converts nquad values to normal Go type

		fmt.Println(nativeValue) // print it!
	}
	if err := nxt.Err(); err != nil {
		log.Fatalln(err)
	}
}
