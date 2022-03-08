#!/bin/bash

set -e

psql -v ON_ERROR_STOP=true -d cli2cloud -a -U user --password 1234 -f /create_schema.sql
