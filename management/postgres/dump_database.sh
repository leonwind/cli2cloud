#!/bin/bash

set -e

pg_dump cli2cloud --username=postgres | gzip > cli2cloud_db_`date +%F-%T`.sql.gz
