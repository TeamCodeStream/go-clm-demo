#!/bin/bash

sleep 5

echo "Running automated tests..."

group1=("http://localhost:9222/version" "http://localhost:9222/notice_error" "http://localhost:9222/notice_error_with_attributes" "http://localhost:9222/custom_event" "http://localhost:9222/set_name" "http://localhost:9222/add_attribute" "http://localhost:9222/add_span_attribute" "http://localhost:9222/ignore")
group2=("http://localhost:9222/segments" "http://localhost:9222/mysql" "http://localhost:9222/roundtripper" "http://localhost:9222/custommetric" "http://localhost:9222/browser" "http://localhost:9222/async" "http://localhost:9222/message" "http://localhost:9222/log" "http://localhost:9222/external" "http://localhost:9222/users" "http://localhost:9222/users/state")

logit() {
  timestamp=$(date +"%F %T")
  echo "$timestamp $1"
}

count=0

SECONDS=0

urls=( ${group1[@]} )

while true; do
    for t in "${urls[@]}"; do
      statusCode=$(curl --write-out '%{http_code}' -s -o /dev/null "$t")
      logit "$t: $statusCode"
      sleep 5
    done
    logit "group $(( count % 2 + 1)) completed a full set of operations of elapsed: $SECONDS"

    # go too fast and the agent starts sampling
    sleep 5
    if (( SECONDS > 50)); then
      sleep 10
      count=$((count+1))
      if (( count %2 == 0 )); then
        urls=( ${group1[@]} )
      else
        urls=( ${group2[@]} )
      fi
      SECONDS=0
    fi
done
