#!/bin/bash

go build -o ./bin/lissajous ./src
./bin/lissajous > out.gif
