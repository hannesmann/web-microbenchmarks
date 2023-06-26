#!/bin/bash

export GODEBUG=http2client=0  # disable HTTP/2 client support
export GODEBUG=$GODEBUG,http2server=0  # disable HTTP/2 server support

all_http="go-nethttp go-fasthttp python-gunicorn rust-actix rust-hyper rust-tinyhttp rust-warp"

read -p "Select 'http' or 'grpc': " benchtype

function starthttp {
	export USEHTTPMON=$2
	http_client="$(pwd)/bin/go-http-client"

	case $1 in 
		"all")
			for i in $all_http; do echo "$i" && echo "" && starthttp $i $2 2>&1 | grep -E 'First request|Average response' && echo ""; done
			;;
		"go-fasthttp")
			(cd go/http/server/fasthttp && go mod download && go build -o go-fasthttp-server)
			mv -f go/http/server/fasthttp/go-fasthttp-server bin/
			"$http_client" ./bin/go-fasthttp-server
			;;

		"go-nethttp")
			(cd go/http/server/nethttp && go mod download && go build -o go-nethttp-server)
			mv -f go/http/server/nethttp/go-nethttp-server bin/
			"$http_client" ./bin/go-nethttp-server
			;;

		"python-gunicorn")
			(cd ./python/http/server/gunicorn && "$http_client" ./run.sh)
			;;

		"rust-actix")
			(cd rust/http/server/actix && cargo build --release)
			mv -f rust/http/server/actix/target/release/rust-actix-server bin/
			"$http_client" ./bin/rust-actix-server
			;;

		"rust-hyper")
			(cd rust/http/server/hyper && cargo build --release)
			mv -f rust/http/server/hyper/target/release/rust-hyper-server bin/
			"$http_client" ./bin/rust-hyper-server
			;;

		"rust-tinyhttp")
			(cd rust/http/server/tinyhttp && cargo build --release)
			mv -f rust/http/server/tinyhttp/target/release/rust-tinyhttp-server bin/
			"$http_client" ./bin/rust-tinyhttp-server
			;;

		"rust-warp")
			(cd rust/http/server/warp && cargo build --release)
			mv -f rust/http/server/warp/target/release/rust-warp-server bin/
			"$http_client" ./bin/rust-warp-server
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
		(cd go/http/client && go mod download && go build -o go-http-client)
		mv -f go/http/client/go-http-client bin/

		use_httpmon=0

		read -p "Use httpmon? (Y/N): " httpmon

		if [ $httpmon == "y" ] || [ $httpmon == "Y" ]; then
			use_httpmon=1
		fi

		read -p "Select HTTP server (all $all_http): " httpserver

		starthttp $httpserver $use_httpmon
		;;

	"grpc")
		echo "Compiling client..."

		mkdir -p bin
		(cd go/grpc/client && go mod download && go build -o go-grpc-client)
		mv -f go/grpc/client/go-grpc-client bin/

		read -p "Select gRPC server (go-grpc rust-tonic): " grpcserver
		;;

	*)
		echo "Unknown bench type $benchtype"
		;;
esac
