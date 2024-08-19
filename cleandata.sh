#!/bin/bash

# clean out all data for fresh db
docker-compose down --volumes
docker volume rm tucows-grill-service_mysql-data
docker-compose up --build
