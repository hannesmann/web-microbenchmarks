# web-microbenchmarks

Evaluating HTTP and gRPC frameworks for HydraGen

Tested with:

* Rust 1.70.0
* Go 1.20.5

## Results

All tests were done on 2x Intel Xeon E5-2685 at 2.1 GHz.

### HTTP simple benchmark

The simple HTTP benchmark sends 10000 requests sequentially on a single thread. The server responds to every request with a single character "r".

| Language | Framework | Startup time (first request) | Average response time |
| -------- | --------- | ---------------------------- | --------------------- |
| Go       | net/http  | 4.984 ms                     | 0.3022 ms             |
| Go       | fasthttp  | 5.505 ms                     | 0.2496 ms             |
| Python   | Gunicorn  | 485.8 ms                     | 1.401 ms              |
| Rust     | Actix Web | 68.38 ms                     | 0.2630 ms             |
| Rust     | Hyper     | 5.684 ms                     | 0.2454 ms             |
| Rust     | tiny-http | 1.722 ms                     | 0.2570 ms             |
| Rust     | Warp      | 6.357 ms                     | 0.2542 ms             |

### HTTPmon benchmark

The HTTPmon benchmark uses [cloud-control/httpmon](https://github.com/cloud-control/httpmon) to send 10000 requests, with up to 500 concurrently at the same time, to the server.

`httpmon --url $url --open --concurrency 500 --thinktime 1 --count 10000 --terminate-after-count`

| Language | Framework | Average latency | Maximum latency | 95-percentile latency | 99-percentile latency | Late requests |
| -------- | --------- | --------------- | --------------- | --------------------- | --------------------- | ------------- |
| Go       | net/http  | 1 ms            | 65 ms           | 1 ms                  | 2 ms                  | 11            |
| Go       | fasthttp  | 1 ms            | 42 ms           | 1 ms                  | 3 ms                  | 8             |
| Python   | Gunicorn  | 2 ms            | 41 ms           | 2 ms                  | 3 ms                  | 19            |
| Rust     | Actix Web | 1 ms            | 56 ms           | 1 ms                  | 1 ms                  | 3             |
| Rust     | Hyper     | 1 ms            | 56 ms           | 1 ms                  | 1 ms                  | 5             |
| Rust     | tiny-http | 5 ms            | 20077 ms        | 1 ms                  | 6 ms                  | 8             |
| Rust     | Warp      | 1 ms            | 38 ms           | 1 ms                  | 1 ms                  | 4             |

### gRPC simple benchmark

The simple gRPC benchmark works the same way as the simple HTTP benchmark. 10000 requests are sent sequentially, with a request of a single byte and a response of a single byte.

| Language | Framework | Startup time (first request) | Average response time |
| -------- | --------- | ---------------------------- | --------------------- |
| Go       | gRPC-Go   | 7.782 ms                     | 0.4264 ms             |
| Rust     | Tonic     | 6.019 ms                     | 0.3675 ms             |
