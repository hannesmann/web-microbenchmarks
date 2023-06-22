#!/bin/bash

read -p "Select 'http' or 'grpc': " benchtype

case $benchtype in 
	"http")
		echo "Compiling client..."

		mkdir -p bin
		(cd go/http/client && go build -o go-http-client)
		mv go/http/client/go-http-client bin/

		read -p "Select HTTP server (go-nethttp go-fasthttp rust-actix rust-hyper rust-tinyhttp): " httpserver
		;;

	"grpc")
		read -p "Select gRPC server (go-grpc rust-tonic): " grpcserver
		;;

	*)
		echo "Unknown bench type $benchtype"
		;;
esac