# MapReduce Implementation in Go

This project implements a distributed MapReduce system using Go and gRPC.

## Features

- Distributed processing with master and worker nodes
- Fault tolerance with task reassignment(WIP)
- Configurable through command-line flags
- Supports custom map and reduce functions

## Prerequisites

- Go 1.16+
- gRPC

## Project Structure

- `cmd/`: Command-line applications
- `master/`: Master node implementation
- `worker/`: Worker node implementation
- `rpc/`: Protocol buffer definitions and generated gRPC code

## Contributing

Contributions are welcome. Please fork the repository and submit a pull request with your changes.

## TODO:

- Add tests.
- Code refactoring
- Properly close worker connections. 
- Fault tolerance improvements.