# Pretense

## What is it?

An attempt at:
1. Talking about the albums and artists I know and enjoy, how I came in contact with them
1. Recommending related bands.
1. Explaining the life-long frog boiling that got me from CCM to actually enjoying death metal.

Also an excuse to try some new technologies :shrug:

## What is it not?

- A reflection of the absolute truth on quality music
- A complete list of the things I like or actively listen to
- A hotdog stand..?

## What's in a name?
talking about music, or art in general, can feel a bit pretentious

## Parts

monorepo for convenience, docker-compose etc.

### Visualization (svelte-kit, cytoscape.js)

A read-only ui for looking, but not touching.

### Writer (tbd)

simplistic way of writing and inserting content. Likely markdown files and some sprinkles.

### API (FastAPI)

- Abstracts the db from the two apps above, or at least `vis`
- converts db to pydantic models
- is the closest to things I actually work on, so some experimentation going on

### db (gremlin/janusgraph)

A graph-based nosql thing that seemed like a good idea at the time. Definitely intended for larger scale than what I'm doing, but it looks cool :shrug:

## Setup

api and vis both need `.env` files for runtime, and the api scripts expect environement variables by the same names. If you're running direnv, make a `.envrc`.
Some of the required things are overwritten in docker-compose.

```sh
# get a python venv for the api
docker-compose up --build # to make sure there's a db instance

# open the dump_from_spotify file below and add a token (follow the url in the file)
cd api
PYTHONPATH=. python src/dump_from_spotify.py
PYTHONPATH=. python src/write_<names go here>.py

cd <root>/vis
npm install  # creates the .svelte-kit/tsconfig that jsconfig references.

# stop the compose and re-run that command
# don't `down` the database, otherwise you'll need to re-seed the data.
```