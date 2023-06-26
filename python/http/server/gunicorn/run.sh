#!/bin/bash
function ctrl_c {
	echo "Trapped SIGINT"
	kill -INT $gunicornpid
	result=$(wait $gunicornpid)
	echo "Python gunicorn server stopped"
	exit $result
}

echo "Python gunicorn server started"
gunicorn -w 4 -b 127.0.0.1:9000 app:app &
gunicornpid=$!
trap "ctrl_c" INT
sleep infinity 
