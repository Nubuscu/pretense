# Pretense

This is a project for me to over-explain my taste in music in a non-linear blog. The writing is mostly me writing for myself.
As a domain I'm interested in, I've used it to write and rewrite the same "solution" in different tech stacks - something that's very useful to broaden horizons and get familiar with other technologies.

The whole thing is:

1. built for me, my text content, my spotify likes, etc
1. built and served statically. This project only ever runs locally, then the content is compiled and served later.

> this is the code behind it all. Hosting over in nubuscu.github.io

## Running the project

### prerequisites

1. python installed locally for content scripts
1. docker-compose
1. node 18 for frontend dev work
1. go 1.19+ for api dev work
1. nubuscu.github.io is cloned in this repository alongside the other components

### how to build, etc

1. start the backend with `make run-backend`
1. if there's no data in there, run dump_from_spotify.py
1. rerun the topic generation, write_topic.py -f all
1. `make run-frontend-dev` for local dev or `make build` to compile the frontend
    1. building requires the backend to be running at the same time. Data is requested once, then the compiled frontend is shipped with it all included
1. `make copy-to-deploy` shuffles the compiled files over to the deploy repo.
1. `cd nubuscu.github.io` and commit/push. Deployment will happen automatically on push.
