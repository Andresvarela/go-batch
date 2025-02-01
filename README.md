# Go Batch Processing Application

This Go application is designed to process a large number of data records efficiently using concurrent workers. The main features of the application include:

- **Data Generation**: Generates a specified number of random data records, each of a given size.
- **Batch Processing**: Processes the generated data in batches, with a configurable batch size.
- **Concurrency**: Utilizes a worker pool to process data concurrently, with a configurable number of workers.
- **Performance Logging**: Logs the start and end time of the data processing to measure performance.

## Prerequisites

- Go 1.20 or later
- Git

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/go-batch.git
    cd go-batch
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage

1. Open the `main.go` file and configure the following parameters as needed:
    - `totalRecords`: Number of records to generate.
    - `recordSize`: Size of each record in bytes.
    - `batchSize`: Size of each batch to process.
    - `workerCount`: Number of concurrent workers.

2. Run the application:
    ```sh
    go run main.go
    ```

## Configuration

The main configuration parameters are located in the `main.go` file:
- `totalRecords`: Total number of records to generate.
- `recordSize`: Size of each record in bytes.
- `batchSize`: Number of records per batch.
- `workerCount`: Number of concurrent workers.

## Logging

The application logs the start and end time of the data processing to measure performance. Logs are printed to the console.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.