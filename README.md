# Avro Serialization Comparison

This project demonstrates how to serialize and compress Go structs using Avro and Zstandard (zstd), and compares the results with JSON serialization.

## Features

- Defines a user schema in Avro (`schemas/user.avsc`).
- Generates a new user struct (`internal/user/user.go`).
- Serializes the user struct to Avro and JSON.
- Compresses both Avro and JSON outputs using zstd.
- Logs the size and compressed size of each format.

## Usage

1. Install dependencies:
   ```sh
   go mod tidy
   ```
2. Run the application:
   ```sh
   go run ./cmd/main.go
   ```

## Dependencies
- [hamba/avro](https://github.com/hamba/avro) for Avro serialization
- [klauspost/compress](https://github.com/klauspost/compress) for zstd compression
- [rs/zerolog](https://github.com/rs/zerolog) for structured logging

