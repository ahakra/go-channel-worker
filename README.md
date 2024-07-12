# Go Channel Worker

This project demonstrates a worker pattern using Go channels and a worker pool to read data from a CSV file, process it, and send it to a Kafka topic.

## Project Structure

```
go-channel-worker/
│
├── file-reader/
│ └── filereader.go
│
├── kafka/
│ └── kafka.go
│
├── model/
│ └── model.go
│
├── workers/
│ └── workers.go
│
└── main.go
```

## Overview

The application follows a pattern where:

- A `FileReader` reads data from a CSV file and sends it to a channel.
- A worker pool consisting of multiple goroutines consumes data from this channel and processes it.
- Each worker sends the processed data to a Kafka topic using a `KafkaProducer`.

## Key Concepts

### Channels

Channels are used to communicate between different parts of the application:

- **FileReader**: Reads data from the file and sends it through a channel.
- **Workers**: Receive data from the channel, process it, and send it to Kafka.

### Worker Pool

The worker pool pattern is used to manage a fixed number of worker goroutines that consume tasks from a shared channel. This allows efficient processing of tasks without overwhelming the system with too many concurrent goroutines.

## Packages

### filereader

Contains the `FileReader` struct and methods for initializing a CSV reader and reading data from the file into a channel.

### kafka

Contains the `KafkaProducer` struct and methods for creating a Kafka producer and sending messages to a Kafka topic.

### model

Contains the `MyData` struct which defines the structure of the data read from the CSV file.

### workers

Contains the `Workers` struct and methods for starting worker goroutines that consume messages from a channel and produce them to Kafka.

### main

The entry point of the application. It sets up the file reader, Kafka producer, and workers, and orchestrates the data flow from file reading to message production.

## Usage

1. Clone the repository:

   ```sh
   git clone https://github.com/ahakra/go-channel-worker.git
   cd go-channel-worker
   ```

2. Install dependencies:

   ```sh
   go mod tidy
   ```

3. Create a `file.txt` file with your CSV data.

4. Run the application:
   ```sh
   go run main.go
   ```

## Configuration

- `location`: The path to the CSV file.
- `urls`: The list of Kafka broker URLs.
- `topic`: The Kafka topic to which the messages will be sent.
- `workersNeeded`: The number of worker goroutines to be started.

## Dependencies

- [sarama](https://github.com/Shopify/sarama): Go library for Apache Kafka.

## Author

[Ahmad Akra](https://github.com/ahakra)
