# Pretense

this be the backend. Hosting over in nubuscu.github.io

## prereq

1. python installed locally for content scripts
1. docker-compose
1. node 18 for frontend dev work
1. nubuscu.github.io is cloned in this repository alongside the other components

## how to build, etc

1. start the backend with `make run-backend`
1. if there's no data in there, run dump_from_spotify.py
1. rerun the topic generation, write_topic.py -f all
1. `make run-frontend-dev` for local dev or `make build` to compile the frontend
    1. building requires the backend to be running at the same time. Data is requested once, then the compiled frontend is shipped with it all included
1. `make copy-to-deploy` shuffles the compiled files over to the deploy repo.
1. `cd nubuscu.github.io` and commit/push. Deployment will happen automatically on push.
