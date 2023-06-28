tonic::include_proto!("benchmark");

use tonic;
use benchmark_service_server::*;

#[derive(Debug, Default)]
pub struct TestService {}

#[tonic::async_trait]
impl BenchmarkService for TestService {
	async fn benchmark(
		&self,
		_request: tonic::Request<Request>,
	) -> Result<tonic::Response<Response>, tonic::Status> {
		Ok(tonic::Response::new(Response { data: "r".into() }))
	}
}

const GRPC_ADDRESS: &str = "127.0.0.1";
const GRPC_PORT: u16 = 9500;

#[tokio::main(flavor = "multi_thread", worker_threads = 4)]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
	let addr = format!("{}:{}", GRPC_ADDRESS, GRPC_PORT).parse()?;
	let svc = TestService::default();

	println!("Rust Tonic server started");

	tonic::transport::Server::builder()
		.add_service(BenchmarkServiceServer::new(svc))
		.serve(addr)
		.await?;

	println!("Rust Tonic server stopped");

	Ok(())
}