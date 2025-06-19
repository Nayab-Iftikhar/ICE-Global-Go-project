#!/bin/bash
# Migrate the database
migrate -path migrations -database "mysql://root:password@tcp(localhost:3306)/todoapp?multiStatements=true" up
