#!/bin/bash

go build -o ./bin/main ./src
./bin/main weight 100
./bin/main distance 100
./bin/main temperature 100
