#!/bin/bash

kill $(lsof -t -i:8001)
kill $(lsof -t -i:8002)
kill $(lsof -t -i:8003)

TZ=US/Eastern ./clock2 --port 8001 &
TZ=Asia/Tokyo ./clock2 --port 8002 &
TZ=Europe/London ./clock2 --port 8003 &

./clockwall Newyork=localhost:8001 Tokyo=localhost:8002 London=localhost:8003
