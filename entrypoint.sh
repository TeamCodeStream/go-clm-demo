#!/usr/bin/env bash

# start app
nohup ./main > ./logs/app.log 2>&1 &

# generate load
./tester.sh
