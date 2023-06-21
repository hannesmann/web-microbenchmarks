#!/bin/bash

read -p "Select 'http' or 'grpc': " benchtype

case $benchtype in 
	"http")
		read -p "Select HTTP client (go-http go-fasthttp rust-hyper): " httpclient
		read -p "Select HTTP server (go-http go-fasthttp rust-actix rust-hyper rust-tinyhttp rust-warp): " httpserver
		;;

	"grpc")
		read -p "Select gRPC server (go-grpc rust-tonic): " grpcserver
		;;

	*)
		echo "Unknown bench type $benchtype"
		;;
esac