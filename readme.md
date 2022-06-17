# Music Map // Pretense

## What is it?

An attempt at talking about the albums and artists I know and enjoy, how I came in contact with them, and who I'd recommend them to, and why.

...In a client/server app so I can relate things together in a more complete manner than _just_ writing about them.

Also an excuse to try some new technologies :shrug:

## What is it not?

- A reflection of the absolute truth on quality music
- A complete list of the things I like or actively listen to
- A hotdog stand..?

## Parts

### Visualization (something js probably)

- 'deployable' ui
- good for reading
- doesn't need to be super secure

### Writer (same as vis)

- thing for me to write content in
- semi-ui. Mostly just conveniences for me making links etc

### API (Python/FastAPI)

- the backend, keeps the other two from reading directly from the db

### db (postgres)

- where the magic ~~happens~~ is stored

## Setup

api and vis both need `.env` files for runtime, and the api scripts expect environement variables by the same names. If you're running direnv, make a `.envrc`.
Some of the required things are overwritten in docker-compose.

```sh
# get a python venv for the api
docker-compose up --detach # to make sure there's a db instance
cd api
alembic upgrade head

# open the dump_from_spotify file below and add a token (follow the url in the file)
PYTHONPATH=. python src/dump_from_spotify.py
(cd src && python <write_topic_file_goes_here.py>)

cd <root>/vis
npm install  # creates the .svelte-kit/tsconfig that jsconfig references.

#stop and restart docker-compose. may need to rebuild.
# don't `down` the database, otherwise you'll need to re-seed the data.
```