This is an outline of some of the mistakes we made / will make and sort or project steps, this will map very sort of roughly to tags.

**IMPORTANT:** Anytime you make a mistake, that should be corrected and have an additional tag made for it adn a writeup of it --
there are "planned" mistakes, but I am guessing we will have a lot of unplanned ones too!

- Task: Import Pokemon (generation 1)
  - _(maybe, but actually code)_ we import using no UUID and then show a problem... correction or something?
    - Story could be someone messed up ID by accident in source file and we need to correct... but we used ID as primary ID and now everything references that...
    - **Zero tag** _(if we decide to do it)_
  - Naked import of data (no namespaces, no type predicate)
    - **1st tag** _(remember tags have to pushed explicitly)_
  - Realize we need type predicate cause they will more other types of data in here -- add it
    - **2nd tag**
  - Realize that "type" will likely conflict with the "type" of pokemon, realize we need some sort of namespaces
    - Do we namespace our data, or the data we import, or both?
    - Do namespaces already exist in the world?  Introduce RDF and the idea of formal namespaces
    - Since we already made the mistake, the "naked" namespace belongs to us, and we import data from the pokedex with pokedex:WHATEVER
    - **3rd tag**
   - Show a few basic queries of information
- Task: Add locations
  - Add in locations
  - Show a few basic intersection queries (at location Y, weight above X)
  - **4th tag**
- Task: Add evolutions
  - Add in evolutions
  - Show that you can "walk" the evolutions to all the end points (8 with Eevee)
  - **5th tag**
- Task: Add labels
  - Add in label of source data: https://github.com/PokeAPI/pokeapi/tree/4803d224eddce9ddcdf49b402b1bcd58c873a25d/data/v2/csv
  - Talk about how vague labels are, etc
  - **6th tag**
- Task: Web interface
  - Build interface, realize that if you want to update while having interface up, bolt will no longer work
  - Dump to RDF quad file
  - Switch backend to PostgreSQL 
  - Load RDF
  - Show that you can now update data while keeping the web interface up
  - **7th tag**
