# Benchmark results

### HTTP simple benchmark

The simple HTTP benchmark sends 10000 requests sequentially on a single thread. The server responds to every request with a single character "r".

| Language | Framework | Startup time (first request) | Average response time |
| -------- | --------- | ---------------------------- | --------------------- |
| Go       | Echo      | 4.495 ms                     | 0.2770 ms             |
| Go       | fasthttp  | 4.398 ms                     | 0.2415 ms             |
| Go       | Gin       | 16.02 ms                     | 0.2784 ms             |
| Go       | net/http  | 4.693 ms                     | 0.2676 ms             |
| Python   | Gunicorn  | 452.5 ms                     | 1.331 ms              |
| Rust     | Actix Web | 45.65 ms                     | 0.2349 ms             |
| Rust     | Hyper     | 2.150 ms                     | 0.2071 ms             |
| Rust     | tiny-http | 1.444 ms                     | 0.2253 ms             |
| Rust     | Warp      | 2.766 ms                     | 0.2155 ms             |

### HTTPmon benchmark

The HTTPmon benchmark uses [cloud-control/httpmon](https://github.com/cloud-control/httpmon) to send 10000 requests, with up to 1000 concurrently at the same time, to the server.

`httpmon --url $url --open --concurrency 640 --thinktime 1 --count 10000 --terminate-after-count`

| Language | Framework | Average latency | Maximum latency | 95-percentile latency | 99-percentile latency | Late requests (`accOpenQueuing`) |
| -------- | --------- | --------------- | --------------- | --------------------- | --------------------- | -------------------------------- |
| Go       | Echo      | 1 ms            | 32 ms           | 1 ms                  | 1 ms                  | 4                                |
| Go       | fasthttp  | 1 ms            | 31 ms           | 1 ms                  | 1 ms                  | 3                                |
| Go       | Gin       | 1 ms            | 72 ms           | 1 ms                  | 15 ms                 | 5                                |
| Go       | net/http  | 1 ms            | 40 ms           | 1 ms                  | 1 ms                  | 4                                |
| Python   | Gunicorn  | 2 ms            | 31 ms           | 2 ms                  | 4 ms                  | 17                               |
| Rust     | Actix Web | 1 ms            | 63 ms           | 1 ms                  | 5 ms                  | 6                                |
| Rust     | Hyper     | 1 ms            | 27 ms           | 1 ms                  | 1 ms                  | 4                                |
| Rust     | Warp      | 1 ms            | 44 ms           | 1 ms                  | 1 ms                  | 5                                |

Rust tiny-http was not included because it finished with 9009 errors.

### gRPC simple benchmark

The simple gRPC benchmark works the same way as the simple HTTP benchmark. 10000 requests are sent sequentially, with a request of a single byte and a response of a single byte.

| Language | Framework         | Startup time (first request) | Average response time |
| -------- | ----------------- | ---------------------------- | --------------------- |
| Go       | gRPC-Go           | 6.223 ms                     | 0.3884 ms             |
| Python   | google.protobuf   | 684.0 ms                     | 0.7162 ms             |
| Rust     | Tonic             | 3.631 ms                     | 0.3858 ms             |
