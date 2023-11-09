# mta-hosting-optimizer

- A HTTP/REST endpoint to retrieve hostnames having less or equals X active IP
  addresses exists.

## Getting Started

These instructions will help you set up and run the project on your local machine.

### Prerequisites

- [Go](https://golang.org/doc/install) installed on your machine.

### Installing

1. Clone the repository:

    ```bash
    git clone https://github.com/sharkx018/mta-hosting-optimizer.git
    ```

2. Change into the project directory:

    ```bash
    cd mta-hosting-optimizer
    ```

3. Install dependencies:

    ```bash
    go mod tidy
    ```

### Running the Project

To run the project, execute the following command:

```bash
go run main.go