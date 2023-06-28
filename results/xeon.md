# Benchmark results

### HTTP simple benchmark

The simple HTTP benchmark sends 10000 requests sequentially on a single thread. The server responds to every request with a single character "r".

| Language | Framework | Startup time (first request) | Average response time |
| -------- | --------- | ---------------------------- | --------------------- |
| Go       | Echo      | 4.800 ms                     | 0.2779 ms             |
| Go       | fasthttp  | 3.654 ms                     | 0.2248 ms             |
| Go       | Gin       | 15.90 ms                     | 0.2716 ms             |
| Go       | net/http  | 4.215 ms                     | 0.2719 ms             |
| Python   | Gunicorn  | 455.0 ms                     | 1.472 ms              |
| Rust     | Actix Web | - ms                         | - ms                  |
| Rust     | Hyper     | - ms                         | - ms                  |
| Rust     | tiny-http | - ms                         | - ms                  |
| Rust     | Warp      | - ms                         | - ms                  |

### HTTPmon benchmark

The HTTPmon benchmark uses [cloud-control/httpmon](https://github.com/cloud-control/httpmon) to send 10000 requests, with up to 1000 concurrently at the same time, to the server.

`httpmon --url $url --open --concurrency 1000 --thinktime 1 --count 10000 --terminate-after-count`

| Language | Framework | Average latency | Maximum latency | 95-percentile latency | 99-percentile latency | Late requests |
| -------- | --------- | --------------- | --------------- | --------------------- | --------------------- | ------------- |
| Go       | Echo      | - ms            | - ms           | - ms                   | - ms                  | -             |
| Go       | fasthttp  | - ms            | - ms           | - ms                   | - ms                  | -             |
| Go       | Gin       | - ms            | - ms           | - ms                   | - ms                  | -             |
| Go       | net/http  | - ms            | - ms           | - ms                   | - ms                  | -             |
| Python   | Gunicorn  | - ms            | - ms           | - ms                   | - ms                  | -             |
| Rust     | Actix Web | - ms            | - ms           | - ms                   | - ms                  | -             |
| Rust     | Hyper     | - ms            | - ms           | - ms                   | - ms                  | -             |
| Rust     | tiny-http | - ms            | - ms           | - ms                   | - ms                  | -             |
| Rust     | Warp      | - ms            | - ms           | - ms                   | - ms                  | -             |

### gRPC simple benchmark

The simple gRPC benchmark works the same way as the simple HTTP benchmark. 10000 requests are sent sequentially, with a request of a single byte and a response of a single byte.

| Language | Framework         | Startup time (first request) | Average response time |
| -------- | ----------------- | ---------------------------- | --------------------- |
| Go       | gRPC-Go           | - ms                         | - ms                  |
| Python   | google.protobuf   | - ms                         | - ms                  |
| Rust     | Tonic             | - ms                         | - ms                  |
