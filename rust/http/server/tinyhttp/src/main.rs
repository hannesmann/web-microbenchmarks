use ctrlc;
use tiny_http::{Server, Response};

const HTTP_HOST: &str = "127.0.0.1:9000";

fn main() {
	println!("Rust tiny-http server started");

	let _ = ctrlc::set_handler(|| {
		println!("Rust tiny-http server stopped");
		std::process::exit(0);
	});

	let server = Server::http(HTTP_HOST).unwrap();

	// Single threaded
	for request in server.incoming_requests() {
		if request.url() != "/" {
			panic!("Invalid path {}", request.url())
		}
		
		let response = Response::from_string("r");
		let _ = request.respond(response);
	}
}