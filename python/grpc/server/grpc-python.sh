#!/bin/bash
function ctrl_c() { 
	echo "Caught SIGINT" 
	kill -9 "$servicepid"
	exit 0
}

trap ctrl_c SIGINT

python app.py &
servicepid=$!

echo "Python service process ID: $servicepid (may need to be killed manually)"
wait $servicepid
