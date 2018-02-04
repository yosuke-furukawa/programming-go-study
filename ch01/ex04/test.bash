#!/bin/bash

go build -o ./bin/dup ./src
./bin/dup "./fixtures/example1.txt" "./fixtures/example2.txt" "./fixtures/example3.txt" 
