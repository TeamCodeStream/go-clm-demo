#!/bin/bash

logfile="logs/tester.log"

sleep 5

echo "Running automated tests..." >$logfile

urls=("http://localhost:9222/" "http://localhost:9222/version" "http://localhost:9222/notice_error" "http://localhost:9222/notice_error_with_attributes" "http://localhost:9222/custom_event" "http://localhost:9222/set_name" "http://localhost:9222/add_attribute" "http://localhost:9222/add_span_attribute" "http://localhost:9222/ignore" "http://localhost:9222/segments" "http://localhost:9222/mysql" "http://localhost:9222/roundtripper" "http://localhost:9222/custommetric" "http://localhost:9222/browser" "http://localhost:9222/async" "http://localhost:9222/message" "http://localhost:9222/log")

while true; do
  for t in "${urls[@]}"; do
    curl -s -o /dev/null "$t"
    sleep 1
  done

  echo "Completed a full set of operations." >>$logfile

  # go too fast and the agent starts sampling
  sleep 4
done
