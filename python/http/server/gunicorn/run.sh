#!/bin/bash
function ctrl_c {
	kill -INT $gunicornpid
	echo "Python gunicorn server stopped"
	exit 0
}

echo "Python gunicorn server started"
gunicorn -b 127.0.0.1:9000 app:app &
gunicornpid=$!
trap "ctrl_c" INT
sleep infinity 
