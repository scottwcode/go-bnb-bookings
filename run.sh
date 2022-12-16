#!/bin/bash

#go build -o bookings cmd/web/*.go && ./bookings
#go build -o go-bnb-bookings cmd/web/*.go && ./go-bnb-bookings
go build -o go-bnb-bookings cmd/web/*.go 
#./go-bnb-bookings
./go-bnb-bookings -dbname=go-bnb-bookings -dbuser=scott7 -cache=false -production=false
