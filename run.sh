#!/bin/bash

export GODEBUG=http2client=0  # disable HTTP/2 client support
export GODEBUG=$GODEBUG,http2server=0  # disable HTTP/2 server support

read -p "Select 'http' or 'grpc': " benchtype

function starthttp {
	case $1 in 
		"go-fasthttp")
			(cd go/http/server/fasthttp && go build -o go-fasthttp-server)
			mv go/http/server/fasthttp/go-fasthttp-server bin/
			./bin/go-http-client ./bin/go-fasthttp-server
			;;

		"go-nethttp")
			(cd go/http/server/nethttp && go build -o go-nethttp-server)
			mv go/http/server/nethttp/go-nethttp-server bin/
			./bin/go-http-client ./bin/go-nethttp-server
			;;

		"rust-actix")
			(cd rust/http/server/actix && cargo build --release)
			mv rust/http/server/actix/target/release/rust-actix-server bin/
			./bin/go-http-client ./bin/rust-actix-server
			;;

		"rust-hyper")
			(cd rust/http/server/hyper && cargo build --release)
			mv rust/http/server/hyper/target/release/rust-hyper-server bin/
			./bin/go-http-client ./bin/rust-hyper-server
			;;

		"rust-tinyhttp")
			(cd rust/http/server/tinyhttp && cargo build --release)
			mv rust/http/server/tinyhttp/target/release/rust-tinyhttp-server bin/
			./bin/go-http-client ./bin/rust-tinyhttp-server
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