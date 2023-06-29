# Benchmark results

### HTTP simple benchmark

The simple HTTP benchmark sends 10000 requests sequentially on a single thread. The server responds to every request with a single character "r".

| Language | Framework | Startup time (first request) | Average response time |
| -------- | --------- | ---------------------------- | --------------------- |
| Go       | Echo      | 4.103 ms                     | 0.02596 ms            |
| Go       | fasthttp  | 2.906 ms                     | 0.01964 ms            |
| Go       | Gin       | 3.144 ms                     | 0.02549 ms            |
| Go       | net/http  | 3.114 ms                     | 0.02549 ms            |
| Python   | Gunicorn  | 94.79 ms                     | 0.1839 ms             |
| Rust     | Actix Web | 3.531 ms                     | 0.02061 ms            |
| Rust     | Hyper     | 3.952 ms                     | 0.01852 ms            |
| Rust     | tiny-http | 3.522 ms                     | 0.02302 ms            |
| Rust     | Warp      | 3.128 ms                     | 0.01878 ms            |

### HTTPmon benchmark

The HTTPmon benchmark uses [cloud-control/httpmon](https://github.com/cloud-control/httpmon) to send 10000 requests, with up to 1000 concurrently at the same time, to the server.

`httpmon --url $url --open --concurrency 640 --thinktime 1 --count 10000 --terminate-after-count`

| Language | Framework | Average latency | Maximum latency | 95-percentile latency | 99-percentile latency | Late requests |
| -------- | --------- | --------------- | --------------- | --------------------- | --------------------- | ------------- |
| Go       | Echo      | 0 ms            | 13 ms           | 0 ms                  | 0 ms                  | 1             |
| Go       | fasthttp  | 0 ms            | 10 ms           | 0 ms                  | 0 ms                  | 1             |
| Go       | Gin       | 0 ms            | 11 ms           | 0 ms                  | 0 ms                  | 0             |
| Go       | net/http  | 0 ms            | 21 ms           | 0 ms                  | 0 ms                  | 1             |
| Python   | Gunicorn  | 0 ms            | 11 ms           | 0 ms                  | 0 ms                  | 2             |
| Rust     | Actix Web | 0 ms            | 18 ms           | 0 ms                  | 0 ms                  | 2             |
| Rust     | Hyper     | 0 ms            | 19 ms           | 0 ms                  | 0 ms                  | 3             |
| Rust     | tiny-http | 0 ms            | 12 ms           | 0 ms                  | 0 ms                  | 3             |
| Rust     | Warp      | 0 ms            | 19 ms           | 0 ms                  | 0 ms                  | 2             |

### gRPC simple benchmark

The simple gRPC benchmark works the same way as the simple HTTP benchmark. 10000 requests are sent sequentially, with a request of a single byte and a response of a single byte.

| Language | Framework         | Startup time (first request) | Average response time |
| -------- | ----------------- | ---------------------------- | --------------------- |
| Go       | gRPC-Go           | 1.538 ms                     | 0.04473 ms            |
| Python   | google.protobuf   | 52.22 ms                     | 0.08712 ms            |
| Rust     | Tonic             | 0.8179 ms                    | 0.04009 ms            |
