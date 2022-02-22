# Music Map

## What is it?
An attempt at talking about the albums and artists I know and enjoy, how I came in contact with them, and who I'd recommend them to, and why... In the format of a data storage layer.

Also an excuse to try some new technologies :shrug:

## What is it not?

A hotdog stand..?

## Parts

### Visualization ()
- 'deployable' ui
- good for reading
- doesn't need to be super secure

### Writer (same as vis)
- thing for me to write content in
- semi-ui. Mostly just conveniences for me making links etc

### API (rust!?)
- the backend, keeps the other two from reading directly from the db

### db (postgres)
- where the magic is stored

## deps
as I go
```
# rust
sudo apt-get install libpq-dev
cargo install diesel_cli --no-default-features --features postgres
```