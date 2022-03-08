#!/bin/bash

set -e

psql -v ON_ERROR_STOP=true --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" /create_schema.sql
