# MapReduce Implementation in Go

This project implements a distributed MapReduce system using Go and gRPC.

## Features

- Distributed processing with master and worker nodes
- Fault tolerance with task reassignment(WIP)
- Configurable through command-line flags
- Supports custom map and reduce functions
- Comprehensive test suite with 46+ unit tests

## Prerequisites

- Go 1.16+
- gRPC

## Project Structure

- `cmd/`: Command-line applications
- `master/`: Master node implementation
- `worker/`: Worker node implementation
- `rpc/`: Protocol buffer definitions and generated gRPC code
- `internal/`: Shared internal packages

## Testing

This project includes comprehensive unit tests covering core functionality.

### Run all tests
```bash
go test ./...
```

### Run tests with coverage
```bash
go test -cover ./...
```

### Run tests with verbose output
```bash
go test -v ./...
```

### Generate coverage report
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Run tests for specific packages
```bash
go test ./master/...
go test ./worker/...
```

### Current Test Coverage
- **Master module**: 12.6% coverage, 28 tests
- **Worker module**: 10.6% coverage, 18 tests
- **Total**: 46+ tests passing

## Contributing

Contributions are welcome. Please fork the repository and submit a pull request with your changes.

## TODO:

- ~~Add tests~~ ✅ **Done** (46+ tests implemented)
- Code refactoring
- Properly close worker connections
- Fault tolerance improvements
- Add integration tests with mock gRPC servers
- Increase test coverage for I/O operations