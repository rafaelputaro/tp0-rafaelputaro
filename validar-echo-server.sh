#!/bin/bash
URL=server
URL_ALT="0.0.0.0"
PORT="12345"
OK_MESSAGE="action: test_echo_server | result: success"
FAIL_MESSAGE="action: test_echo_server | result: fail"
MESSAGE_TO_SERVER="Hello Server"

RESPONSE=$(echo "$MESSAGE_TO_SERVER" | nc -w 3 $URL_ALT $PORT)
if [ "$RESPONSE" = "$MESSAGE_TO_SERVER" ]; then
  echo "$OK_MESSAGE"
else
  RESPONSE=$(echo "$MESSAGE_TO_SERVER" | nc -w 3 $URL $PORT)
  if [ "$RESPONSE" = "$MESSAGE_TO_SERVER" ]; then
    echo "$OK_MESSAGE"
  else
    echo "$FAIL_MESSAGE"
  fi
fi