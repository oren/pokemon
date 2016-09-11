## Switch to PostgreSQL Storage Engine

```
  cayley dump --db=bolt --dbpath=data/pokemon.boltdb   # dump the database into a quad file
  cayley init --config=cayley.cfg                      # assumes the database exist but no table
  cayley load --config=cayley.cfg --quads=dbdump.nq    # load a quad file and using a configuration file
```

## Web console
```
cayley http --config=cayley.cfg
```

Find what pichu evolves into after 2 phases of evolution

```
  g.V("pichu").In("name").Out("evolves_to").Out("evolves_to").Out("name").All()
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
  g.V().In("name").Out("evolves_to").Out("evolves_to").Out("name").All()
```

Find all the evolutions of eevee
```
  g.V("eevee").In("name").Out("evolves_to").Out("name").All()
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
