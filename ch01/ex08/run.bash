#!/bin/bash

go build -o ./bin/fetch ./src
./bin/fetch "http://recruit.co.jp"
