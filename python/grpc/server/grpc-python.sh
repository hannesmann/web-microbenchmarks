#!/bin/bash
function ctrl_c() { 
	echo "Caught SIGINT"
	# https://stackoverflow.com/questions/2618403/how-to-kill-all-subprocesses-of-shell
	pkill -P $$
	wait $!
	exit 0
}

trap ctrl_c SIGINT

python app.py &
wait $!