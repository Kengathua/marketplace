#!/bin/bash
source env.sh

go run migrations/migrate.go

go run default_data/load_default_data.go

go run server.go