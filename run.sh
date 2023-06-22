#!/bin/bash

read -p "Select 'http' or 'grpc': " benchtype

function starthttp {
	case $1 in 
		"go-nethttp")
			(cd go/http/server/nethttp && go build -o go-nethttp-server)
			mv go/http/server/nethttp/go-nethttp-server bin/
			./bin/go-http-client ./bin/go-nethttp-server
		;;

		*)
		echo "Unknown server type $1"
		;;
	esac
}

case $benchtype in 
	"http")
		echo "Compiling client..."

		mkdir -p bin
		(cd go/http/client && go build -o go-http-client)
		mv go/http/client/go-http-client bin/

		read -p "Select HTTP server (go-nethttp go-fasthttp rust-actix rust-hyper rust-tinyhttp): " httpserver
		starthttp $httpserver
		;;

	"grpc")
		read -p "Select gRPC server (go-grpc rust-tonic): " grpcserver
		;;

	*)
		echo "Unknown bench type $benchtype"
		;;
esac