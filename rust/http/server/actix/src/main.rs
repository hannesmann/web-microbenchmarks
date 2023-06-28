use actix_web::{get, App, HttpServer, Responder};
use tokio;

#[get("/")]
async fn index() -> impl Responder { "r" }

const HTTP_ADDRESS: &str = "127.0.0.1";
const HTTP_PORT: u16 = 9000;

#[tokio::main(flavor = "multi_thread", worker_threads = 4)]
async fn main() -> std::io::Result<()> {
	println!("Rust Actix Web server started");

	let _ = HttpServer::new(|| App::new().service(index))
		.bind((HTTP_ADDRESS, HTTP_PORT))?
		.run()
		.await?;

	println!("Rust Actix Web server stopped");
	Ok(())
}