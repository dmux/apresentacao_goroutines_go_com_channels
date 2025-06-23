## Go Benchmark
Running 5m test @ http://go:8080
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.51ms    0.97ms  68.30ms   73.61%
    Req/Sec     9.96k   600.54    12.69k    73.10%
  11897335 requests in 5.00m, 1.31GB read
  Socket errors: connect 0, read 0, write 0, timeout 1
Requests/sec:  39651.56
Transfer/sec:      4.46MB

## Python Benchmark
Running 5m test @ http://python:8080
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   273.06ms  514.86ms   2.00s    83.48%
    Req/Sec     3.66k     1.11k    4.95k    88.31%
  2425561 requests in 5.00m, 326.16MB read
  Socket errors: connect 0, read 0, write 0, timeout 1186
Requests/sec:   8082.66
Transfer/sec:      1.09MB
