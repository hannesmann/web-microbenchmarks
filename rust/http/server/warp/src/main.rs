use warp::Filter;

const HTTP_ADDRESS: [u8; 4] = [127, 0, 0, 1];
const HTTP_PORT: u16 = 9000;

#[tokio::main(flavor = "multi_thread", worker_threads = 4)]
async fn main() {
	println!("Rust Warp server started");

    let response = warp::path!().map(|| "r");

    warp::serve(response)
        .run((HTTP_ADDRESS, HTTP_PORT))
        .await;

	println!("Rust Warp server stopped");
}