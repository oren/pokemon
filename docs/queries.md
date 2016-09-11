## Web console
```
cayley http --config=cayley.cfg
```

Find what pichu evolves into after 2 phases of evolution

```
g.V("pichu").In("<schema:name>").Out("<rdf:evolves_to>").Out("<rdf:evolves_to>").Out("<schema:name>").All()

{
  "result": [
    {
    "id": "raichu"
    }
  ]
}
```

Find all pokemons that are the result of 2 phases of evolution

```
g.V().In("<schema:name>").Out("<rdf:evolves_to>").Out("<rdf:evolves_to>").Out("<schema:name>").All()
```

Find all the evolutions of eevee
```
g.V("eevee").In("<schema:name>").Out("<rdf:evolves_to>").Out("<schema:name>").All()

{
  "result": [
    {
    "id": "flareon"
    },
    {
    "id": "jolteon"
    },
    {
    "id": "vaporeon"
    }
  ]
}
```

## Repl
```
cayley repl --config=cayley.cfg
```

## HTTP API

```
curl http://localhost:64210/api/v1/query/gremlin -d 'g.V("eevee").In("name").Out("evolves_to").Out("name").All()'
```
