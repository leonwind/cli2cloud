#!/bin/bash

set -e

go build -o service
sleep 10s
./service
