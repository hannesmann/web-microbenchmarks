#!/bin/bash
function ctrl_c() { 
	echo "Caught SIGINT" 
	kill -9 "$gunicornpid"
	exit 0
}

trap ctrl_c SIGINT

echo "Python gunicorn server started"
gunicorn -w 4 -b 127.0.0.1:9000 app:app &
gunicornpid=$!

echo "Gunicorn process ID: $gunicornpid (may need to be killed manually)"
wait $gunicornpid
