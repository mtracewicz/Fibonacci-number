# Fibonacci-number [![Go](https://github.com/mtracewicz/Fibonacci-number/actions/workflows/go.yml/badge.svg)](https://github.com/mtracewicz/Fibonacci-number/actions/workflows/go.yml)

To test run `docker-compose up` and send get request to `localhost:4444/n` (where n is which Fibonacci number You want).
Repository runs automatic tests based on github actions on push and pull request to master branch.

Example requests to test:
```bash
curl localhost:4444/20
curl localhost:4444/0
curl localhost:4444/-1
```