# web-microbenchmarks

Evaluating HTTP and gRPC frameworks for HydraGen

Tested with:

* Rust 1.70.0
* Go 1.20.5

## Results

All tests were done on 2x Intel Xeon E5-2685 at 2.1 GHz.

### HTTP simple benchmark

The simple benchmark sends 10000 requests sequentially on a single thread. The server responds to every request with a single character "r".

| Language | Framework | Startup time (first request) | Average response time |
| -------- | --------- | ---------------------------- | --------------------- |
| Go       | net/http  | 4.984 ms                     | 0.3022 ms             |
| Go       | fasthttp  | 5.505 ms                     | 0.2496 ms             |
| Rust     | Actix Web | 68.38 ms                     | 0.2630 ms             |
| Rust     | Hyper     | 5.684 ms                     | 0.2454 ms             |
| Rust     | tiny-http | 1.722 ms                     | 0.2570 ms             |
| Rust     | Warp      | 6.357 ms                     | 0.2542 ms             |