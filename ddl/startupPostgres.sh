#!/bin/bash

RUNNING=$(podman ps -a | grep -i postgres)

if [[ "$RUNNING" != "" ]]; then
    podman stop postgres
    podman rm postgres
fi

podman run -d --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword postgres:14.19

sleep 5

go run . create