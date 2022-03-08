#!/bin/bash

set -e

go build -o service

# Wait 10s such that postgres is ready to connect clients
sleep 15s

./service
