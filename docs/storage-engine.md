## Switch to PostgreSQL Storage Engine

```
  cayley dump --db=bolt --dbpath=data/pokemon.boltdb   # dump the database into a quad file
  cayley init --config=cayley.cfg                      # assumes the database exist but no table
  cayley load --config=cayley.cfg --quads=dbdump.nq    # load a quad file and using a configuration file
```
