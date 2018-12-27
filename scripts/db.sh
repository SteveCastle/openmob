#!/bin/bash
# Docker command to run an empty docker db with dev tern settings.
docker run --rm   --name pg-docker -e POSTGRES_PASSWORD=tern -e POSTGRES_USER=tern -d -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data  postgres
