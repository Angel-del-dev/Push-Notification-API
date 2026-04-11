#!/bin/bash
docker compose down --volumes --rmi all
cd ..
docker build --no-cache -t notification-authentication-api-api .
cd scripts