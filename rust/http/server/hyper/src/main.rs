use std::convert::Infallible;
use std::net::SocketAddr;

use hyper::{Body, Request, Response, Server};
use hyper::service::{make_service_fn, service_fn};

const HTTP_ADDRESS: [u8; 4] = [127, 0, 0, 1];
const HTTP_PORT: u16 = 9000;

async fn shutdown_signal() {
	let _ = tokio::signal::ctrl_c().await;
	println!("Rust Hyper server stopped");
}

async fn respond(_req: Request<Body>) -> Result<Response<Body>, Infallible> {
	Ok(Response::new("R".into()))
}

#[tokio::main]
async fn main() {
	println!("Rust Hyper server started");

	let addr = SocketAddr::from((HTTP_ADDRESS, HTTP_PORT));

	let make_svc = make_service_fn(|_conn| async {
		Ok::<_, Infallible>(service_fn(respond))
	});

	let server = Server::bind(&addr).serve(make_svc);
	let graceful_server = server.with_graceful_shutdown(shutdown_signal());

	let _ = graceful_server.await;
}