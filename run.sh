#!/bin/bash

read -p "Select 'http' or 'grpc': " benchtype

case $benchtype in 
	"http")
		read -p "Select HTTP server (go-http go-fasthttp rust-actix rust-hyper rust-tinyhttp): " httpserver
		;;

	"grpc")
		read -p "Select gRPC server (go-grpc rust-tonic): " grpcserver
		;;

	*)
		echo "Unknown bench type $benchtype"
		;;
esac