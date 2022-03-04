# Music Map

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

## deps

as I go

### api

```sh
cd api
pip install -r requirements_dev.in
docker-compose up -d  # starts a postgres db
cp example.envrc .envrc  # insert environment variables
direnv allow
alembic upgrade head  # setup db schema
make run  # start the server
```
