#!/bin/bash
URL="server"
PORT="12345"
OK_MESSAGE="action: test_echo_server | result: success"
FAIL_MESSAGE="action: test_echo_server | result: fail"
MESSAGE_TO_SERVER="Hello Server"

#RESPONSE=$(echo "$MESSAGE_TO_SERVER" | nc -w 3 $URL $PORT)
RESPONSE=$(echo "$MESSAGE_TO_SERVER" | nc "server" 12345)

if [ "$RESPONSE" = "$MESSAGE_TO_SERVER" ]; then
  echo "$OK_MESSAGE"
else
  echo "$FAIL_MESSAGE"
fi