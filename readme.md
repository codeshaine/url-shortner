Sure, here's a detailed README for your URL shortener project:

---

# URL Shortener

This project is a URL shortener service built with Go, PostgreSQL, and Chi for routing. The service allows users to shorten URLs and redirect to the original URL using the shortened version. The project includes unit tests to ensure the functionality of the URL validation, insertion, and retrieval processes.

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Running Tests](#running-tests)
- [Project Structure](#project-structure)
- [Environment Variables](#environment-variables)
- [Build and Run](#build-and-run)
- [Makefile Commands](#makefile-commands)

## Features

- Shorten long URLs
- Redirect to the original URL using the shortened version
- Track click counts
- Expire URLs based on conditions (e.g., click count)

## Requirements

- Go 1.21.5 or higher
- PostgreSQL
- A .env file with the necessary environment variables (see below)

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/codeshaine/url-shortner.git
   cd url-shortner
   ```

2. Install dependencies:

   ```sh
   go mod tidy
   ```

3. Create a `.env` file in the root directory with the following content:

   ```env
   HOST=localhost
   DB_PORT=5432
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=url_shortner
   PORT=8080
   ```

4. Ensure PostgreSQL is running and accessible with the credentials provided in the `.env` file.

## Usage

### Starting the Server

1. Build the project:

   ```sh
   make build
   ```

2. Run the project:
   ```sh
   make run
   ```

The server will start on `http://localhost:8080`.

### Shorten a URL

Send a POST request to `/shorten` with the URL you want to shorten as a query parameter:

```sh
curl -X POST "http://localhost:8080/shorten?url=https://example.com"
```

### Redirect to Long URL

Send a GET request with the shortened URL:

```sh
curl -X GET "http://localhost:8080/{short_url}"
```

Replace `{short_url}` with the actual shortened URL string.

## Running Tests

To run the unit tests, execute the following command:

```sh
make test
```

For more verbose output, use:

```sh
make test-v
```

To force run tests, use:

```sh
make test-force
```

## Project Structure

```
url-shortner/
├── bin/                   # Compiled binary directory
├── cmd/                   # Main application entry point
│   └── main.go
├── db/                    # Database connection and query functions
│   ├── db.go
│   └── models.go
├── internal/              # Internal packages
│   ├── controller/        # HTTP handlers
│   │   └── url_controller.go
│   ├── response/          # API response structure
│   │   └── response.go
│   ├── router/            # Router setup
│   │   └── router.go
│   └── utils/             # Utility functions
│       ├── env.go
│       ├── url_generator.go
│       └── url_validator.go
├── test/                  # Test files
│   ├── db_test.go
│   ├── url_generator_test.go
│   └── url_validator_test.go
├── .env                   # Environment variables file
├── go.mod                 # Go module file
├── go.sum                 # Go dependencies file
├── Makefile               # Makefile with build and test commands
└── README.md              # Project documentation
```

## Environment Variables

The project relies on environment variables defined in a `.env` file located in the root directory. The variables are:

- `HOST`: Database host (e.g., localhost)
- `DB_PORT`: Database port (e.g., 5432)
- `DB_USER`: Database username
- `DB_PASSWORD`: Database password
- `DB_NAME`: Database name
- `PORT`: Port on which the server will run (e.g., 8080)

## Makefile Commands

The `Makefile` includes several helpful commands:

- `make build`: Build the project and generate a binary in the `bin/` directory.
- `make run`: Run the project using the generated binary.
- `make test`: Run all tests.
- `make test-v`: Run all tests with verbose output.
- `make test-force`: Force run all tests without caching.
- `make clean`: Clean the `bin/` directory by removing the compiled binary.
- `make sync`: Remove unwanted packages and update the project dependencies.

---

This README provides a comprehensive overview of your URL shortener project, detailing how to set up, use, test, and build the application. It should be sufficient for anyone looking to understand and work with the project.
