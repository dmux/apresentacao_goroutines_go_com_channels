## Go Benchmark (example)
```
Running 5m test @ http://go:8080
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.2ms    0.8ms    15ms   75.00%
    Req/Sec    65k      4k      80k     60.00%
  19,500,000 requests in 5.00m, 1.3GB read
```

## Python Benchmark (example)
```
Running 5m test @ http://python:8080
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     5.1ms    2.0ms    30ms   70.00%
    Req/Sec    19k      2k      24k     65.00%
  5,700,000 requests in 5.00m, 430MB read
```

These numbers are illustrative. Run `docker-compose up --build bench` to produce real results.
