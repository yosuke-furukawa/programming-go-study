#!/bin/sh
go build -o sha256 ../../ch04/ex02/src
go build -o bzipper

./sha256 < /usr/share/dict/words
./bzipper < /usr/share/dict/words | wc -c
./bzipper < /usr/share/dict/words | bunzip2 | ./sha256
