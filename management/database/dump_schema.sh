#!/bin/bash

set -e

pg_dump cli2cloud --username=postgres -s > create_schema.sql
