#!/bin/bash
URL=server
OK_MESSAGE="action: test_echo_server | result: success"
FAIL_MESSAGE="action: test_echo_server | result: fail"
MESSAGE_TO_SERVER="Hello Server"

RESPONSE=$(echo "$MESSAGE_TO_SERVER" | nc $URL $PORT)

if [ "$RESPONSE" = "$MESSAGE_TO_SERVER" ]; then
  echo "$OK_MESSAGE"
else
  echo "$FAIL_MESSAGE"
fi