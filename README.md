[![Tegridy](https://github.com/saurabhmshr/tegridyfarms/actions/workflows/testing.yaml/badge.svg)](https://github.com/saurabhmshr/tegridyfarms/actions/workflows/testing.yaml)

# Tegridy Farms

REST-API based on South Park's Tegridy Farms

## Getting Started

+ Run HTTP server

  `go run ./cmd/main/tegridy.go`

+ To view text content, navigate to:

  1. Hello, World!: http://127.0.0.1:8080/hello
  2. Server Health: http://127.0.0.1:8080/health

+ To view rendered content, navigate to:

  1. Homepage (Image): http://127.0.0.1:8080/home
  2. Outdoor  (HTML):  http://127.0.0.1:8080/outdoor

+ Perform create, read, update, and delete operations:

  ```
  ./config/product_ops.sh
  ```
