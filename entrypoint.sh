#!/usr/bin/env bash

# start app
nohup ./main &

# generate load
./tester.sh
