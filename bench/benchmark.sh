#!/bin/sh
set -e
sleep 5
mkdir -p /results

printf "## Go Benchmark\n" > /results/report.md
wrk -t4 -c100 -d5m http://go:8080 >> /results/report.md

printf "\n## Python Benchmark\n" >> /results/report.md
wrk -t4 -c100 -d5m http://python:8080 >> /results/report.md