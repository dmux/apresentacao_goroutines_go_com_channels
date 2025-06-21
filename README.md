# Go vs Python Concurrency Benchmark

This setup compares a simple HTTP server written in Go with one written in Python using FastAPI. The benchmark uses **wrk** for load testing and can be run with Docker Compose.

## Usage

```bash
docker-compose up --build bench
```

The benchmark runs for 5 minutes against each server and saves a Markdown report inside `results/report.md`.
An example report is already included in that directory. Run the command above to generate real numbers on your machine.

## Services

- **go** – compiled static binary based on `scratch`.
- **python** – FastAPI running on `uvicorn`.
- **bench** – Alpine container running `wrk` to generate load.

Ports 8081 and 8082 expose the Go and Python servers respectively.
