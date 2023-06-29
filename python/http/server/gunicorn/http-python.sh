#!/bin/bash
function ctrl_c() { 
	echo "Python gunicorn server stopped"
	echo "Caught SIGINT"
	# https://stackoverflow.com/questions/2618403/how-to-kill-all-subprocesses-of-shell
	pkill -P $gunicornpid
	pkill -P $$
	wait $!
	exit 0
}

trap ctrl_c SIGINT

echo "Python gunicorn server started"
gunicorn -w 4 -b 127.0.0.1:9000 app:app &
gunicornpid=$!
wait $!